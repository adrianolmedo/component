package components

import (
	"regexp"
	"strings"
	"text/template"

	"github.com/maragudk/gomponents"
)

// TextWithBoldStars returns a text node with <strong> tags around
// the text with bold stars.
//
// Example:
//
//	TextWithBoldStars("Hello *world*")
//	returns Hello <strong>world</strong>
func TextWithBoldStars(text string) gomponents.Node {
	text = template.HTMLEscapeString(text)
	return gomponents.Raw(textWithBoldStarsStr(text))
}

// textWithBoldStarsStr is the core functionality of TextWithBoldStars.
func textWithBoldStarsStr(input string) string {
	re := regexp.MustCompile(`\*(.*?)\*`)
	return re.ReplaceAllStringFunc(input, func(s string) string {
		return "<strong>" + strings.Trim(s, "*") + "</strong>"
	})
}
