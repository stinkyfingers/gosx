package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"syscall/js"

	"github.com/stinkyfingers/gosx/example/components"
)

//https://github.com/siongui/godom/tree/master/wasm
func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM|syscall.SIGQUIT|syscall.SIGINT)

	body := js.Global().Get("document").Get("body")

	// div := &element.Element{
	// 	Tag: "div",
	// }
	// attach.AppendToValue(div, body, nil)
	//
	// a := &element.Element{
	// 	Tag:        "a",
	// 	InnerHTML:  "cnn",
	// 	Attributes: map[string]string{"href": "https://cnn.com"},
	// }
	// attach.Append(a, div, nil)
	//
	// button := js.Global().Get("document").Call("createElement", "div")
	// text := js.Global().Get("document").Call("createTextNode", "BUTTON")
	// button.Call("setAttribute", "id", "myButton") // set if not popuilated
	// button.Call("appendChild", text)
	// body.Call("appendChild", button)
	// cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	fmt.Println("button clicked")
	// 	return nil
	// })
	// js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)
	// // funcs := map[string]js.Func{"test": test}
	//
	// tags := `<div className="foobar">
	// 	<a href="https://www.ebay.com" onClick="test">EBAY</a>
	// 	<div onClick="test">CLICK</div>
	// 	<img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTiTnm8jW-nmGGT08sRrxsPFiZQ7r5tI2IYIABNcS56JIbpJTl-" />
	// </div>`
	// elements, err := element.ParseElement(tags)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// attach.AttachElements(elements, body, funcs)
	ctx, cf := context.WithCancel(context.Background())
	components.YearSelect(body, ctx)
	<-ch
	cf()
}
