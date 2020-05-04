/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	appsv1alpha1 "git.indie.host/operators/rocketchat-operator/api/v1alpha1"
	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	instance "k8s.libre.sh/instance"

	"github.com/presslabs/controller-util/syncer"
)

// RocketchatReconciler reconciles a Rocketchat object
type RocketchatReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

func ignoreNotFound(err error) error {
	if apierrs.IsNotFound(err) {
		return nil
	}
	return err
}

// NextcloudReconciler implements Reconcile interface
func (r *RocketchatReconciler) GetClient() client.Client          { return r.Client }
func (r *RocketchatReconciler) GetScheme() *runtime.Scheme        { return r.Scheme }
func (r *RocketchatReconciler) GetRecorder() record.EventRecorder { return r.Recorder }
func (r *RocketchatReconciler) GetLogger() logr.Logger            { return r.Log }

// +kubebuilder:rbac:groups=apps.libre.sh,resources=rocketchats,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.libre.sh,resources=rocketchats/status,verbs=get;update;patch

func (r *RocketchatReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("rocketchat", req.NamespacedName)
	log.Info("reconciling")

	app := &appsv1alpha1.Rocketchat{}
	if err := r.Get(ctx, req.NamespacedName, app); err != nil {
		log.Error(err, "unable to fetch Nextcloud")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, ignoreNotFound(err)
	}

	app.SetDefaults()

	objects := app.Spec.Settings.GetObjects()

	for _, obj := range objects {
		s := instance.NewObjectSyncer(obj, app, r)

		if err := syncer.Sync(context.TODO(), s, r.Recorder); err != nil {
			return ctrl.Result{}, err
		}
	}

	syncers := instance.NewSyncers(app, r, app)

	err := r.sync(syncers)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *RocketchatReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1alpha1.Rocketchat{}).
		Complete(r)
}

func (r *RocketchatReconciler) sync(syncers []syncer.Interface) error {
	for _, s := range syncers {
		if err := syncer.Sync(context.TODO(), s, r.Recorder); err != nil {
			return err
		}
	}
	return nil
}
