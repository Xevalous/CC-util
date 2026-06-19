APP = CC-util
BUILD_DIR = .

build:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP).exe .

clean:
	rm -f $(BUILD_DIR)/$(APP).exe
