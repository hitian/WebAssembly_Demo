package main

import (
	"errors"
	"fmt"
	"sync"
	"syscall/js"
)

var (
	version = "debug"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	fmt.Println("version ", version)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	registerJsFunctions()

	wg.Wait()
}

func registerJsFunctions() {
	//add callback to element.
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("button clicked")
		cb.Release()
		return nil
	})
	js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)

	//add global function
	var testFunctionParams js.Func
	testFunctionParams = js.FuncOf(printFunctionArgs)
	js.Global().Set("print_args", testFunctionParams)

	js.Global().Set("add", add())
}

func printFunctionArgs(this js.Value, args []js.Value) interface{} {
	fmt.Println(this)
	fmt.Println(args)
	return nil
}

func add() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) (result interface{}) {
		defer func() {
			if err := recover(); err != nil {
				result = createErrorResult(fmt.Errorf("%s", err))
				return
			}
		}()
		if len(args) != 2 {
			return createErrorResult(errors.New("wrong args length, need 2"))
		}

		a, b := args[0].Int(), args[1].Int()
		val := createResult()
		val["data"] = a + b
		result = val
		return
	})
}

func createErrorResult(err error) map[string]interface{} {
	result := make(map[string]interface{})
	result["ok"] = false
	result["error"] = err.Error()
	return result
}

func createResult() map[string]interface{} {
	result := make(map[string]interface{})
	result["ok"] = true
	return result
}
