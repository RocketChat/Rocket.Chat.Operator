package v1alpha1

import (
	instance "k8s.libre.sh/instance"
	interfaces "k8s.libre.sh/interfaces"
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

	if &c.Workload.Service.Port.Port == nil || c.Workload.Service.Port.Port == 0 {
		c.Workload.Service.Port.Port = 3000
	}
	if len(c.Workload.Service.Port.Protocol) == 0 {
		c.Workload.Service.Port.Protocol = "TCP"
	}
	if len(c.Workload.Service.Port.Name) == 0 {
		c.Workload.Service.Port.Name = "http"
	}

	// Set Ingress Defaults
	for _, p := range c.Workload.Deployment.Parameters {
		if p.Key == "ROOT_URL" && len(p.Value) > 0 {
			c.Workload.Network.Ingress.Host = p.Value
		}
	}

}

func (app *Rocketchat) SetDefaults() {
	app.Spec.Settings.SetDefaults(app)
	app.Spec.App.SetDefaults(app)
}
