module github.com/chremoas/auth-esi-poller

go 1.14

require (
	github.com/chremoas/auth-srv v1.3.0
	github.com/chremoas/esi-srv v1.3.0
	github.com/chremoas/services-common v1.3.2
	github.com/micro/go-micro v1.9.1
	github.com/petergtz/pegomock v2.8.0+incompatible // indirect
	github.com/smartystreets/goconvey v0.0.0-20190710185942-9d28bd7c0945 // indirect
	go.uber.org/zap v1.10.0
	golang.org/x/net v0.0.0-20190724013045-ca1201d0de80
)

replace github.com/chremoas/auth-esi-poller => ../auth-esi-poller

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
