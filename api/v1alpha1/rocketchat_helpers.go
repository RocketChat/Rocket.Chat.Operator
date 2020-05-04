package v1alpha1

import (
	instance "k8s.libre.sh/instance"
	interfaces "k8s.libre.sh/interfaces"
)

func (o *Rocketchat) GetOwner() interfaces.Object { return o }
func (o *Rocketchat) GetName() string             { return o.Name }
func (o *Rocketchat) GetNamespace() string        { return o.Namespace }
func (o *Rocketchat) GetInstance() string         { return o.Name }
func (o *Rocketchat) GetVersion() string          { return o.Spec.Version }
func (o *Rocketchat) GetComponent() string        { return "instance" }
func (o *Rocketchat) GetPartOf() string           { return "rocketchat" }
func (o *Rocketchat) GetManagedBy() string        { return "rocketchat-operator" }
func (o *Rocketchat) GetApplication() string      { return "rocketchat" }

func (o *Rocketchat) GetSettings() instance.Settings {
	return o.Spec.Settings
}

func (app *Rocketchat) GetComponents() map[int]instance.Component {
	components := map[int]instance.Component{
		0: app.Spec.App,
	}

	return components
}

func (o *App) GetName() string { return "app" }

func (c *App) SetDefaults(i instance.Instance) {
	c.Workload.SetDefaults(i)
}

func (app *Rocketchat) SetDefaults() {
	app.Spec.Settings.SetDefaults(app)
	app.Spec.App.SetDefaults(app)
}
