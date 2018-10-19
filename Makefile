BIN_NAME=tormag
RELEASE_DIR=release

all:
	@echo
	@echo "Commands          : Description"
	@echo "----------------- : -----------"
	@echo "make deps         : Install the dependencies and tools"
	@echo "make run          : Run the program"
	@echo "make tools        : Install the tools"
	@echo "make bindata		 : Generate all bindata base on go-bindata command"
	@echo "make generate     : Run the go generate command"
	@echo "make build        : Build the biniary file"
	@echo "make build-macos  : Build the macos biniary file only"
	@echo "make build-windows: Build the windows biniary file only"
	@echo "make build-freebsd: Build the freebsd biniary file only"
	@echo "make build-linux  : Build the linux biniary file only"
	@echo "make pack         : Pack the build files into release directory"
	@echo "make release      : Run the build and pack command"
	@echo

deps: tools
	@dep ensure

run:
	@go run *.go

tools:
	@go get -u github.com/jteeuwen/go-bindata/...

bindata:
	# package: views, output: pkg/view/view.go, ignore: *.go, template: views/**
	@go-bindata -pkg view -o pkg/view/view.go -ignore=.go views/...

	# package: static, output: pkg/static/static.go, ignore: *.go, template: static/**
	@go-bindata -pkg static -o pkg/static/static.go -ignore=.go static/...

	# package: config, output: config/config.go, ignore: *.go, template: config/**
	@go-bindata -pkg config -o config/config.go -ignore=.go config/...

generate:
	@go generate -x

build: bindata build-macos build-windows build-windows build-freebsd build-linux

build-macos:
	@echo "Building for macos ..."

	@CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -o $(BIN_NAME)-macos

build-windows:
	@echo "Building for windows ..."

	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o $(BIN_NAME)-win.exe

build-freebsd:
	@echo "Building for freebsd ..."

	@CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o $(BIN_NAME)-freebsd

build-linux:
	@echo "Building for linux ..."

	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN_NAME)-linux

pack:
	@echo "Packing ..."

	@rm -rf $(RELEASE_DIR)
	@mkdir -p $(RELEASE_DIR)
	@mv $(BIN_NAME)-* $(RELEASE_DIR)

release: build pack
