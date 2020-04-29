module git.indie.host/operators/rocketchat-operator

go 1.13

require (
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.8.1
	github.com/presslabs/controller-util v0.2.2
	github.com/yuin/goldmark v1.1.30 // indirect
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/tools v0.0.0-20200428204708-317da45f2f19 // indirect
	k8s.io/api v0.18.1
	k8s.io/apimachinery v0.18.1
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.libre.sh v0.0.0-00010101000000-000000000000
	sigs.k8s.io/controller-runtime v0.5.0
)

replace k8s.libre.sh => ./../new

replace k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible => k8s.io/client-go v0.17.2

replace k8s.io/apimachinery v0.18.1 => k8s.io/apimachinery v0.17.2
