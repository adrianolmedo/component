package gomputils

import (
	"strings"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// ManyClasses returns a string with all the classes joined by a space.
func ManyClasses(classes ...string) gomponents.Node {
	str := strings.Join(classes, " ")
	return html.Class(str)
}
