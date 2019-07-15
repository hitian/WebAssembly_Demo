.PHONY: clean build

clean: 
	rm -rf main.wasm
	
build:
	#cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
	GOOS=js GOARCH=wasm go build -o main.wasm

serve:
	go get github.com/shurcooL/goexec
	goexec 'http.ListenAndServe(":9000", http.FileServer(http.Dir(".")))'
