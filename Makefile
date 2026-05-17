APP := pswg
VERSION ?= $(shell git describe --tags --dirty --always)
COMMIT ?= $(shell git rev-parse --short HEAD)
DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
GOFLAGS := -trimpath
LDFLAGS := -s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)
BUILD_DIR := build
DIST_DIR := dist

.PHONY: build check clean fmt release test vet

build:
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP) .

check: fmt test vet

fmt:
	gofmt -w pswg.go pswg_test.go genutil/genutil.go genutil/genutil_test.go

test:
	go test ./...

vet:
	go vet ./...

release: clean check
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
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(DIST_DIR)/$(APP).exe .
	tar -C $(DIST_DIR) -czf $(DIST_DIR)/$(APP)_$(VERSION)_windows_amd64.tar.gz $(APP).exe
	rm $(DIST_DIR)/$(APP).exe
	cd $(DIST_DIR) && shasum -a 256 *.tar.gz > SHA256SUMS

clean:
	rm -rf $(BUILD_DIR) $(DIST_DIR)
