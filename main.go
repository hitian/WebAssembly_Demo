package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
