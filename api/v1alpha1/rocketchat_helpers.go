package v1alpha1

import (
	application "k8s.libre.sh/application"
	components "k8s.libre.sh/application/components"
	interfaces "k8s.libre.sh/interfaces"
	meta "k8s.libre.sh/meta"
)

func (o *Rocketchat) GetOwner() interfaces.Object { return o }
func (o *Rocketchat) GetName() string             { return o.Name }
func (o *Rocketchat) GetNamespace() string        { return o.Namespace }
func (o *Rocketchat) GetInstance() string         { return o.Name }
func (o *Rocketchat) SetInstance(s string)        {}
func (o *Rocketchat) GetVersion() string          { return o.Spec.Version }
func (o *Rocketchat) SetVersion(s string)         {}
func (o *Rocketchat) GetComponent() string        { return "instance" }
func (o *Rocketchat) SetComponent(s string)       {}
func (o *Rocketchat) GetPartOf() string           { return "rocketchat" }
func (o *Rocketchat) SetPartOf(s string)          {}
func (o *Rocketchat) GetManagedBy() string        { return "rocketchat-operator" }
func (o *Rocketchat) SetManagedBy(s string)       {}
func (o *Rocketchat) GetApplication() string      { return "rocketchat" }
func (o *Rocketchat) SetApplication(s string)     {}

func (o *Rocketchat) GetSettings() application.Settings {
	s := application.NewSettings(o.Spec.Settings)
	return s
}

func (app *Rocketchat) GetComponents() map[int]application.Component {
	components := map[int]application.Component{
		0: app.Spec.App,
	}

	return components
}

// func (o *App) GetName() string { return "app" }

func (c *App) SetDefaults() {

	c.Workload.SetDefaults()

	if c.Workload.Backend == nil {
		c.Workload.Backend = &components.Backend{}
	}
	if &c.Workload.Backend.Port == nil || c.Workload.Backend.Port.Port == 0 {
		c.Workload.Backend.Port.Port = 3000
	}
	if len(c.Workload.Backend.Port.Protocol) == 0 {
		c.Workload.Backend.Port.Protocol = "TCP"
	}
	if len(c.Workload.Backend.Port.Name) == 0 {
		c.Workload.Backend.Port.Name = "http"
	}

	if len(c.Workload.Backend.Paths) == 0 {
		c.Workload.Backend.Paths = []string{"/"}
	}

	// Set Ingress Defaults
	/* 	for _, p := range *c.Workload.Settings.Parameters {
		if p.Key == "ROOT_URL" && len(p.Value) > 0 {
			c.Workload.Network.Ingress.Host = p.Value
		}
	} */

	// TODO TO FIX !!
	// 	c.Workload.Backend.Hostname = "chat.libre.sh"

}

func (app *Rocketchat) SetDefaults() {

	if app.Spec.Settings.ObjectMeta == nil {
		app.Spec.Settings.ObjectMeta = new(meta.ObjectMeta)
	}

	if app.Spec.Settings.ObjectMeta.Labels == nil {
		app.Spec.Settings.ObjectMeta.Labels = make(map[string]string)
	}

	app.Spec.Settings.ObjectMeta.SetComponent("settings")

	meta.SetObjectMeta(app, app.Spec.Settings.ObjectMeta)
	app.Spec.Settings.SetDefaults()
	//	app.Spec.App.SetDefaults()
}
