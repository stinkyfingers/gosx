package components

import (
	"context"
	"syscall/js"

	"github.com/stinkyfingers/gosx/attach"
	"github.com/stinkyfingers/gosx/element"
)

type year struct {
	value     string
	elementID string
	cb        js.Func
}

func (y *year) onChange(this js.Value, vals []js.Value) interface{} {
	y.value = this.Get("value").String()
	element := js.Global().Get("document").Call("getElementById", y.elementID)
	element.Set("innerHTML", y.value) // TOOD - not a func
	return nil
}

func newYear(value string) *year {
	y := &year{
		value: value,
	}
	y.cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		y.value = this.Get("value").String()
		element := js.Global().Get("document").Call("getElementById", y.elementID)
		element.Set("innerHTML", y.value) // TOOD - not a func
		return nil
	})
	return y
}

func YearSelect(body js.Value, ctx context.Context) {
	y := newYear("tba")
	sel := element.NewElement("select", "", nil, map[string]js.Func{"change": y.cb}, nil)
	optNil := element.NewElement("option", "", map[string]string{"value": ""}, nil, sel)
	opt1 := element.NewElement("option", "2020", map[string]string{"value": "2020"}, nil, sel)
	opt2 := element.NewElement("option", "2019", map[string]string{"value": "2019"}, nil, sel)
	opt3 := element.NewElement("option", "2018", map[string]string{"value": "2018"}, nil, sel)
	div := element.NewElement("div", "", nil, nil, nil)
	y.elementID = div.GosxID

	attach.AttachElements([]element.Element{*sel, *optNil, *opt1, *opt2, *opt3, *div}, body, nil)
	<-ctx.Done()
	y.cb.Release()
}
