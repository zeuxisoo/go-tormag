.PHONY: web

LDFLAGS += -X "github.com/zeuxisoo/go-tormag/internal/setting.BuildEnv=release"
LDFLAGS += -X "github.com/zeuxisoo/go-tormag/internal/setting.BuildTime=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')"
LDFLAGS += -X "github.com/zeuxisoo/go-tormag/internal/setting.BuildHash=$(shell git rev-parse --short=12 HEAD)"


BIN_NAME=tormag
RELEASE_DIR=release

all:
	@echo
	@echo "Commands             : Description"
	@echo "-------------------- : -----------"
	@echo "make deps            : Install the dependencies and tools"
	@echo "make run             : Run the program"
	@echo "make run-web         : Run the web application"
	@echo "make run-web-release : Run the web application with release build environment"
	@echo "make tools           : Install the tools"
	@echo "make bindata         : Generate all bindata base on go-bindata command"
	@echo "make generate        : Run the go generate command"
	@echo "make build           : Build the biniary file"
	@echo "make build-macos     : Build the macos biniary file only"
	@echo "make build-windows   : Build the windows biniary file only"
	@echo "make build-freebsd   : Build the freebsd biniary file only"
	@echo "make build-linux     : Build the linux biniary file only"
	@echo "make pack            : Pack the build files into release directory"
	@echo "make release         : Run the build and pack command"
	@echo

deps: tools
	@dep ensure

run:
	@go run *.go

run-web:
	@go run *.go web

run-web-release:
	@go generate && go run -ldflags '$(LDFLAGS)' *.go web

tools:
	@go get -u github.com/jteeuwen/go-bindata/...

web-build:
	@cd web && yarn install && yarn run build

web-copy:
	@cp -Rf web/dist/* static/

web-replace:
	@perl -p -i -e 's@/(js|css|img)@/static/$$1@g' static/index.html
	@perl -p -i -e 's@/(img)@/static/$$1@g' static/manifest.json
	@perl -p -i -e 's@/(fonts|robots|js|img|css)@/static/$$1@g' static/precache-manifest.*.js
	@perl -p -i -e 's@/(precache\-manifest)@/static/$$1@g' static/service-worker.js
	@perl -p -i -e 's@"(js)/"@"static/$$1/"@g' static/js/app.*.js
	@perl -p -i -e 's@\("".concat@\("static".concat@' static/js/app.*.js

web-integrate:
	@mv static/index.html views/home/index.html

# Add workbox config to specific line
# like => @perl -l -p -i -e 'print "workbox.setConfig({ debug: false });" if $. == 20' static/service-worker.js
	@ex -s -c '20i|workbox.setConfig({ debug: false });' -c x static/service-worker.js

web-clean:
# Remove all without .gitignore:
# like => find ./static -mindepth 1 -maxdepth 1 ! -name '.gitignore' -exec rm -rf {} \\
# rm without remove hidden file by default
	@rm -rf static/*
	@rm -rf views/home/index.html

web: web-build web-copy web-replace web-integrate

bindata:
# package: views, output: internal/view/view.go, ignore: *.go, template: views/**
	go-bindata -pkg view -o internal/view/view.go -ignore=.go views/...

# package: static, output: internal/static/static.go, ignore: *.go, template: static/**
	go-bindata -pkg static -o internal/static/static.go -ignore=.go static/...

# package: config, output: config/config.go, ignore: *.go, template: config/**
	go-bindata -pkg config -o config/config.go -ignore=.go config/...

generate:
	@go generate -x

build: web bindata build-macos build-windows build-windows build-freebsd build-linux

build-macos:
	@echo "Building for macos ..."

	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '$(LDFLAGS)' -o $(BIN_NAME)-macos

build-windows:
	@echo "Building for windows ..."

	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags '$(LDFLAGS)' -o $(BIN_NAME)-win.exe

build-freebsd:
	@echo "Building for freebsd ..."

	@CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags '$(LDFLAGS)' -o $(BIN_NAME)-freebsd

build-linux:
	@echo "Building for linux ..."

	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '$(LDFLAGS)' -o $(BIN_NAME)-linux

build-clean: web-clean
	@rm -rf $(BIN_NAME)-*

pack:
	@echo "Packing ..."

	@rm -rf $(RELEASE_DIR)
	@mkdir -p $(RELEASE_DIR)
	@mv $(BIN_NAME)-* $(RELEASE_DIR)

release: build pack

clean: build-clean
	@rm -rf storage/attachments/*
