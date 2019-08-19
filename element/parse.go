package element

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func ParseElement(str string) ([]Element, error) {
	var elements []Element
	var parentSlice []*Element

	reader := strings.NewReader(str)
	tt := html.NewTokenizer(reader)

	for {
		z := tt.Next()
		if z == html.ErrorToken {
			if tt.Err() != io.EOF {
				return nil, tt.Err()
			}
			break
		}
		token := tt.Token()

		switch token.Type {
		case html.SelfClosingTagToken:
			// parse/append element
			e := newElementFromToken(token, parentSlice)
			elements = append(elements, *e)

		case html.StartTagToken:
			// parse/append element and advance parent slice
			e := newElementFromToken(token, parentSlice)
			elements = append(elements, *e)
			parentSlice = append(parentSlice, e)

		case html.TextToken:
			// strip control chars
			data := strings.Trim(token.Data, " \r\n\t")
			if data == "" {
				continue
			}
			fmt.Println("D", data)
			// look for inline code
			if data[0] == '{' && data[len(data)-1] == '}' {
				fmt.Println("code")
			}
			// set data of most recent element
			elements[len(elements)-1].InnerHTML = data

		case html.EndTagToken:
			// reverse parent slice
			if len(parentSlice) > 0 {
				parentSlice = parentSlice[:len(parentSlice)-1]
			}
		}
	}
	return elements, nil
}

func newElementFromToken(token html.Token, parentSlice []*Element) *Element {
	e := &Element{Attributes: make(map[string]string)}
	e.Tag = token.Data
	for _, attr := range token.Attr {
		if attr.Key == "id" {
			continue
		}
		e.Attributes[attr.Key] = attr.Val
	}
	e.AssignGosxID()
	if len(parentSlice) > 0 {
		e.Parent = parentSlice[len(parentSlice)-1]
	}
	return e
}
