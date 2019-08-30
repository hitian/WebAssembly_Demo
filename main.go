package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"syscall/js"
)

type base64OperationType int

const (
	base64Encode base64OperationType = iota
	base64Decode
)

var (
	version = "debug"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	fmt.Println("version ", version)

	registerJsFunctions()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func registerJsFunctions() {
	//add callback to element.
	var cb js.Func
	//find element by id.
	button := js.Global().Get("document").Call("getElementById", "myButton")
	//define a callback function
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Call("alert", "button click called from webasswmbly.")
		button.Call("removeEventListener", "click", cb)
		button.Call("setAttribute", "disabled", "disabled")
		cb.Release()
		return nil
	})
	// add event listener on element.
	button.Call("addEventListener", "click", cb)

	goVersionDiv := js.Global().Get("document").Call("getElementById", "go_version")
	if goVersionDiv.Type() != js.TypeNull {
		goVersionDiv.Set("innerHTML", "go runtime version: "+runtime.Version())
	}

	//add global function
	var testFunctionParams js.Func
	testFunctionParams = js.FuncOf(printFunctionArgs)
	js.Global().Set("print_args", testFunctionParams)

	js.Global().Set("add", add())
	js.Global().Set("base64_encode", base64Coder(base64Encode))
	js.Global().Set("base64_decode", base64Coder(base64Decode))
	js.Global().Set("go_version", js.FuncOf(func(this js.Value, args []js.Value) (result interface{}) {
		return runtime.Version()
	}))
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

func base64Coder(t base64OperationType) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) (result interface{}) {
		defer func() {
			if err := recover(); err != nil {
				result = createErrorResult(fmt.Errorf("%s", err))
				return
			}
		}()
		if len(args) < 1 {
			return createErrorResult(errors.New("wrong args length"))
		}
		str := args[0].String()
		var res string
		switch t {
		case base64Encode:
			res = base64.StdEncoding.EncodeToString([]byte(str))
		case base64Decode:
			r, err := base64.StdEncoding.DecodeString(str)
			if err != nil {
				return createErrorResult(err)
			}
			res = string(r)
		}
		val := createResult()
		val["data"] = res
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
