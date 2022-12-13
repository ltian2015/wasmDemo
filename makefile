#该makefile支持 make,make clean,make bin,make image 等4个可执行命令。
#注意：makefile 中只有Shello执行命令” 文本行的开头用tab符，而且必须用tab符。
PROJECT="webassemly.helloworld"
GOROOT=$(shell go env GOROOT)
CUR_DIR=$(shell pwd)
WASM_OUT_DIR=./assets
WASM_SRC_DIR=./cmd/wasm
WASM_OUT_FILE_NAME=golangDemo.wasm
WEB_SERVER_SRC="./cmd/server"
WEB_SERVER_FILE=webserver
setup:
	go get -u github.com/boombuler/barcode
defualt: 
#显示项目的名字
	@echo $(PROJECT)
clean:
#清理已生成的目标文件
	rm -rf $(WASM_OUT_DIR)/$(WASM_OUT_FILE_NAME)
	rm -rf $(WASM_OUT_DIR)/wasm_exec.js
wasmTiny: clean
	cp "$(GOROOT)/misc/wasm/wasm_exec.js"  $(WASM_OUT_DIR)
	tinygo build -o $(WASM_OUT_DIR)/$(WASM_OUT_FILE_NAME) -target wasm $(WASM_SRC_DIR)
wasm: clean
#构建浏览器中运行的wasm格式的文件
	cp "$(GOROOT)/misc/wasm/wasm_exec.js"  $(WASM_OUT_DIR)
	GOOS=js GOARCH=wasm go build  -o $(WASM_OUT_DIR)/$(WASM_OUT_FILE_NAME) $(WASM_SRC_DIR)
server:
	rm -rf $(WEB_SERVER_FILE)
	go build  -o $(WEB_SERVER_SRC)/$(WEB_SERVER_FILE) $(WEB_SERVER_SRC)
run: server
	$(WEB_SERVER_SRC)/$(WEB_SERVER_FILE) 
test:
	go test -v -cover ./...
test-watch:
	gomon -R -t
build-watch:
	gomon -R -- make wasm
