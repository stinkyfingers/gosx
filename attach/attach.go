package attach

import (
	"strings"
	"syscall/js"

	"github.com/stinkyfingers/gosx/element"
)

// AttachElements attaches a slice of elements to a js.Value
func AttachElements(elements []element.Element, body js.Value) {
	for _, e := range elements {
		if e.Parent == nil {
			AppendToValue(&e, body)
		} else {
			Append(&e, e.Parent)
		}
	}
}

// AppendToValue ...
func AppendToValue(e *element.Element, v js.Value) {
	element := js.Global().Get("document").Call("createElement", e.Tag)
	text := js.Global().Get("document").Call("createTextNode", e.InnerHTML)
	if e.GosxID == "" {
		e.AssignGosxID()
	}
	element.Call("setAttribute", "id", e.GosxID)
	element.Call("appendChild", text)
	for key, val := range e.Attributes {
		element.Call("setAttribute", key, val)
	}
	for key, val := range e.Listeners {
		element.Call("addEventListener", strings.TrimRight(key, "on"), val)
	}
	v.Call("appendChild", element)
}

// Append ...
func Append(e *element.Element, elem *element.Element) {
	parent := js.Global().Get("document").Call("getElementById", elem.GosxID)
	AppendToValue(e, parent)
}
