package element

import (
	"reflect"
	"testing"
)

func TestParseElement(t *testing.T) {
	str := "<div className='test'>Hello</div>"
	expected := []Element{{
		Tag:        "div",
		InnerHTML:  "Hello",
		Attributes: map[string]string{"classname": "test"},
		GosxID:     "0",
	}}
	e, err := ParseElement(str)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(e, expected) {
		t.Errorf("expected %v, got %v", expected, e)
	}
}

func TestParseElementDeep(t *testing.T) {
	currentGosxID = 0
	str := `<div className='test'><img src="a"/><img src="b"/></div>`
	expected := []Element{{
		Tag:        "div",
		Attributes: map[string]string{"classname": "test"},
		GosxID:     "0",
	}, {
		Tag:        "img",
		InnerHTML:  "",
		Attributes: map[string]string{"src": "a"},
		GosxID:     "1",
	}, {
		Tag:        "img",
		InnerHTML:  "",
		Attributes: map[string]string{"src": "b"},
		GosxID:     "2",
	}}
	e, err := ParseElement(str)
	if err != nil {
		t.Error(err)
	}
	expected[1].Parent = &e[0] // TODO - better means of asserting parent
	expected[2].Parent = &e[0] // TODO - better means of asserting parent
	if len(expected) != len(e) {
		t.Errorf("Different number of elements: got %d, expected %d", len(e), len(expected))
		t.FailNow()
	}

	for i := range expected {
		if expected[i].GosxID != e[i].GosxID {
			t.Errorf("[%d] mismatched id: %s %s ", i, expected[i].GosxID, e[i].GosxID)
		}
		if expected[i].InnerHTML != e[i].InnerHTML {
			t.Errorf("[%d] mismatched InnerHTML: %s %s ", i, expected[i].InnerHTML, e[i].InnerHTML)
		}
		if expected[i].Tag != e[i].Tag {
			t.Errorf("[%d] mismatched Tag: %s %s ", i, expected[i].Tag, e[i].Tag)
		}
		if expected[i].Parent != nil && !reflect.DeepEqual(*expected[i].Parent, *e[i].Parent) {
			t.Errorf("[%d] mismatched Parent: %v %v", i, expected[i].Parent, e[i].Parent)
		}
		if expected[i].Parent == nil && e[i].Parent != nil {
			t.Errorf("[%d] expected parent to be nil, got %v", i, e[i].Parent)
		}
	}
}
