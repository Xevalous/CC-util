APP = cc-util
BUILD_DIR = .
VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo dev)

build:
	GOOS=windows GOARCH=amd64 go-winres simply --out $(BUILD_DIR)/rsrc
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.version=$(VERSION)" -o $(BUILD_DIR)/$(APP).exe .

clean:
	rm -f $(BUILD_DIR)/$(APP).exe $(BUILD_DIR)/rsrc*.syso
