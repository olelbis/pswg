APP := pswg
VERSION ?= $(shell git describe --tags --dirty --always)
PACKAGE_VERSION := $(patsubst v%,%,$(VERSION))
COMMIT ?= $(shell git rev-parse --short HEAD)
DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GOFLAGS := -trimpath
LDFLAGS := -s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)
BUILD_DIR := build
DIST_DIR := dist
NFPM ?= nfpm

.PHONY: build check checksums clean fmt link-check package-manifests packages release release-archives test vet

build:
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP) .

check: fmt test vet link-check

fmt:
	gofmt -w pswg.go pswg_test.go genutil/genutil.go genutil/genutil_test.go genutil/example_test.go

test:
	go test ./...

vet:
	go vet ./...

link-check:
	python3 tools/check-links.py

release: clean check release-archives packages checksums package-manifests

checksums:
	cd $(DIST_DIR) && find . -maxdepth 1 -type f ! -name SHA256SUMS -print | sed 's#^\./##' | sort | xargs shasum -a 256 > SHA256SUMS

package-manifests:
	python3 tools/generate-package-manifests.py --version $(VERSION) --dist $(DIST_DIR)
	cd $(DIST_DIR) && find . -maxdepth 1 -type f ! -name SHA256SUMS -print | sed 's#^\./##' | sort | xargs shasum -a 256 > SHA256SUMS

release-archives:
	mkdir -p $(DIST_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(DIST_DIR)/$(APP) .
	tar -C $(DIST_DIR) -czf $(DIST_DIR)/$(APP)_$(VERSION)_darwin_arm64.tar.gz $(APP)
	rm $(DIST_DIR)/$(APP)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(DIST_DIR)/$(APP) .
	tar -C $(DIST_DIR) -czf $(DIST_DIR)/$(APP)_$(VERSION)_darwin_amd64.tar.gz $(APP)
	rm $(DIST_DIR)/$(APP)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(DIST_DIR)/$(APP) .
	tar -C $(DIST_DIR) -czf $(DIST_DIR)/$(APP)_$(VERSION)_linux_amd64.tar.gz $(APP)
	rm $(DIST_DIR)/$(APP)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(DIST_DIR)/$(APP) .
	tar -C $(DIST_DIR) -czf $(DIST_DIR)/$(APP)_$(VERSION)_linux_arm64.tar.gz $(APP)
	rm $(DIST_DIR)/$(APP)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(DIST_DIR)/$(APP).exe .
	tar -C $(DIST_DIR) -czf $(DIST_DIR)/$(APP)_$(VERSION)_windows_amd64.tar.gz $(APP).exe
	zip -j $(DIST_DIR)/$(APP)_$(VERSION)_windows_amd64.zip $(DIST_DIR)/$(APP).exe
	rm $(DIST_DIR)/$(APP).exe
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(DIST_DIR)/$(APP).exe .
	tar -C $(DIST_DIR) -czf $(DIST_DIR)/$(APP)_$(VERSION)_windows_arm64.tar.gz $(APP).exe
	zip -j $(DIST_DIR)/$(APP)_$(VERSION)_windows_arm64.zip $(DIST_DIR)/$(APP).exe
	rm $(DIST_DIR)/$(APP).exe

packages:
	mkdir -p $(DIST_DIR) $(BUILD_DIR)/package
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/package/$(APP) .
	VERSION=$(PACKAGE_VERSION) $(NFPM) package --config nfpm.amd64.yaml --packager deb --target $(DIST_DIR)/$(APP)_$(VERSION)_linux_amd64.deb
	VERSION=$(PACKAGE_VERSION) $(NFPM) package --config nfpm.amd64.yaml --packager rpm --target $(DIST_DIR)/$(APP)_$(VERSION)_linux_amd64.rpm
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/package/$(APP) .
	VERSION=$(PACKAGE_VERSION) $(NFPM) package --config nfpm.arm64.yaml --packager deb --target $(DIST_DIR)/$(APP)_$(VERSION)_linux_arm64.deb
	VERSION=$(PACKAGE_VERSION) $(NFPM) package --config nfpm.arm64.yaml --packager rpm --target $(DIST_DIR)/$(APP)_$(VERSION)_linux_arm64.rpm

clean:
	rm -rf $(BUILD_DIR) $(DIST_DIR)
