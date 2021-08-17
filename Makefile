.PHONY: clean build

clean: 
	rm -rf main.wasm
	
build:
	cp "`go env GOROOT`/misc/wasm/wasm_exec.js" .
	GOOS=js GOARCH=wasm go build -ldflags="-s -w -X main.version=`date +'%Y-%m-%d_%H_%M_%S'`" -o main.wasm

serve:
	go get github.com/shurcooL/goexec
	goexec 'http.ListenAndServe("127.0.0.1:9000", http.FileServer(http.Dir(".")))'
