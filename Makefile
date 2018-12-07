GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=hkuvrd

BUILD_DIR=build
BUILD_SRC=daemon/hkuvrd.go
BUILD_DEST=

# use arm64 once Rasbian supports 64-bits
PACKAGE_RPI=$(BINARY_NAME)_linux_armhf
PACKAGE_MAC=$(BINARY_NAME)_darwin_amd64

all: test build
build:
	$(GOBUILD) -o $(BUILD_DEST) -i $(BUILD_SRC)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME)

package-rpi: build-rpi
	tar -cvzf $(PACKAGE_RPI).tar.gz -C $(BUILD_DIR) $(BINARY_NAME)

package-mac: build-mac
	tar -cvzf $(PACKAGE_MAC).tar.gz -C $(BUILD_DIR) $(BINARY_NAME)

build-rpi:
	GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -i $(BUILD_SRC)

build-mac:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -i $(BUILD_SRC)