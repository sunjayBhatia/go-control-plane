module github.com/envoyproxy/go-control-plane/xdsmatcher

go 1.20

replace github.com/envoyproxy/go-control-plane => ../

replace github.com/envoyproxy/go-control-plane/envoy => ../envoy

replace github.com/envoyproxy/go-control-plane/ratelimit => ../ratelimit

require (
	github.com/cncf/xds/go v0.0.0-20230428030218-4003588d1b74
	github.com/envoyproxy/go-control-plane/envoy v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.4
	google.golang.org/protobuf v1.30.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
