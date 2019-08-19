package element

import (
	"strconv"
	"syscall/js"
)

// Element represents an HTML element
type Element struct {
	Tag        string
	InnerHTML  string
	Attributes map[string]string
	// TODO non-value attributes like 'disabled'
	Listeners map[string]js.Func
	GosxID    string
	Parent    *Element
}

type Callback func(this js.Value, args []js.Value) interface{}

var currentGosxID int

func NewElement(tag, innerHTML string, attributes map[string]string, listeners map[string]js.Func, parent *Element) *Element {
	e := &Element{
		Tag:        tag,
		InnerHTML:  innerHTML,
		Attributes: attributes,
		Listeners:  listeners,
		Parent:     parent,
	}
	e.AssignGosxID()
	return e
}

func (e *Element) AssignGosxID() string {
	e.GosxID = strconv.Itoa(currentGosxID)
	currentGosxID++
	return e.GosxID
}