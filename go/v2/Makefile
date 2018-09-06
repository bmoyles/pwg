VERSION=	$(shell git describe --tags --always --dirty)
GOVERSION=	$(shell go version)
BUILDTIME=	$(shell date -u +"%FT%T%z")
GO15VENDOREXPERIMENT?= 1
VERSIONFLAGS= \
	-X 'main.appVersion=$(VERSION)' \
	-X 'main.goVersion=$(GOVERSION)' \
	-X 'main.buildTime=$(BUILDTIME)'

# set GO_LD_FLAGS to provide additional flags to the linker
# set GO_BUILDFLAGS to provide additional build flags
# set GO_TESTFLAGS to provide additional test flags

GOLDFLAGS=$(VERSIONFLAGS) $(GO_LD_FLAGS)
GOBUILDFLAGS=$(GO_BUILDFLAGS)
GOTESTFLAGS=$(GO_TESTFLAGS)

.PHONY: test install

all: test install

test:
	go test $(GOTESTFLAGS) . ./cmd/...

install:
	GO15VENDOREXPERIMENT=$(GO15VENDOREXPERIMENT)
	touch cmd/*/version.go
	go install -ldflags "$(GOLDFLAGS)" $(GOBUILDFLAGS) . ./cmd/...
