package attach

import (
	"syscall/js"

	"github.com/stinkyfingers/gosx/element"
)

func Remove(e element.Element) {
	elem := js.Global().Get("document").Call("getElementById", e.GosxID)
	elem.Call("remove")
}
