package parse

import (
	"strings"

	"github.com/stinkyfingers/gosx/element"
)

func ParseElement(str string) (element.Element, error) {
	angleOpenCount := strings.Count(str, "<")
	angleCloseCount := strings.Count(str, ">")
	return element.Element{}
}
