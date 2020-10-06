version = 99.0.0
provider_path = registry.terraform.io/pmorillon/grid5000/$(version)/darwin_amd64/

install_macos:
	go build -o bin/terraform-provider-grid5000_$(version)

	mkdir -p ~/Library/Application\ Support/io.terraform/plugins/$(provider_path)
	cp bin/terraform-provider-grid5000_$(version)  ~/Library/Application\ Support/io.terraform/plugins/$(provider_path)