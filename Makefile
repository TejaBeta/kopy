
PROJECTNAME=$(shell basename "$(PWD)")
BUILD_DIR="bin"
GOARCH_AMD64="amd64"

build:
	@echo "  >  Building "$(PROJECTNAME)": "
	@echo "================================"
	mkdir $(BUILD_DIR)
	@echo "---------Linux Build------------"
	@echo "                                "
	mkdir $(BUILD_DIR)/linux
	GOOS=linux GOARCH=$(GOARCH_AMD64) go build -o $(BUILD_DIR)/linux/$(PROJECTNAME) main.go
	@echo "                                "
	@echo "-----Completed linux build------"
	@echo "                                "
	mkdir $(BUILD_DIR)/mac
	@echo "----------Mac Build-------------"
	@echo "                                "
	GOOS=darwin GOARCH=$(GOARCH_AMD64) go build -o $(BUILD_DIR)/darwin/$(PROJECTNAME) main.go
	@echo "                                "
	@echo "-----Completed Mac build--------"
	@echo "                                "
	@echo "---------Win64 Build------------"
	@echo "                                "
	mkdir $(BUILD_DIR)/win64
	GOOS=windows GOARCH=$(GOARCH_AMD64) go build -o $(BUILD_DIR)/win64/$(PROJECTNAME) main.go
	@echo "                                "
	@echo "-----Completed Win64 build------"
	@echo "                                "
	@echo "---------FreeBSD Build----------"
	@echo "                                "
	mkdir $(BUILD_DIR)/freebsd
	GOOS=freebsd GOARCH=$(GOARCH_AMD64) go build -o $(BUILD_DIR)/freebsd/$(PROJECTNAME) main.go
	@echo "                                "
	@echo "----Completed FreeBSD build-----"

clean:
	@echo "  >  Cleaning build cache"
	if [ -d "$(BUILD_DIR)" ]; then rm -Rf $(BUILD_DIR); fi
	go clean

test:
	@echo "  > Running unit test cases"
	go test ./... -v

cover:
	@echo "  > Calculating unit test cases coverage"
	go test ./... -cover

.PHONY: clean