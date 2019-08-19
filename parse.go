package gosx

import (
	"strconv"
	// "syscall/js"
)

// TODO - handle nil errorss

// Element represents an HTML element
type Element struct {
	Tag        string
	InnerHTML  string
	Attributes map[string]string
	// TODO non-value attributes like 'disabled'
	// TODO skip ID attribute and set on own
	GosxID string
}

type Globaler interface {
	Global() Valuer
}

type Valuer interface {
	Call(m string, args ...interface{}) Valuer
	Get(p string) Valuer
}

var currentGosxID int

func (e *Element) asssignGosxID() string {
	e.GosxID = strconv.Itoa(currentGosxID)
	currentGosxID++
	return e.GosxID
}

// AppendToValue ...
func (e *Element) AppendToValue(v Valuer, g Globaler) {
	element := g.Global().Get("document").Call("createElement", e.Tag)
	text := g.Global().Get("document").Call("createTextNode", e.InnerHTML)
	element.Call("appendChild", text)
	for k, val := range e.Attributes {
		element.Call("setAttribute", k, val)
	}
	element.Call("setAttribute", "id", e.asssignGosxID())
	v.Call("appendChild", element)
}

// Append ...
func (e *Element) Append(elem Element, g Globaler) {
	parent := g.Global().Get("document").Call("getElementById", elem.GosxID)
	e.AppendToValue(parent, g)
}

func ParseElement(str string) Element {

	return Element{}
}
