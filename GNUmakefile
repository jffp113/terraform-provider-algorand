host_name = terraform.local
namespace = local
type = algorand
target = darwin_amd64
version = 1.0.0
bin = terraform-provider-algorand

build:
	go build
	mkdir -p ~/.terraform.d/plugins/$(host_name)/$(namespace)/$(type)/$(version)/$(target)/
	mv $(bin) ~/.terraform.d/plugins/$(host_name)/$(namespace)/$(type)/$(version)/$(target)/$(bin)