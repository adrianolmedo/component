package components

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// Table is a table with a wrapper div.
//
// The styles are defined in a css file located where the
// tailwindcss css file is located.
func Table(children ...gomponents.Node) gomponents.Node {
	return html.Div(
		html.Class("table-wrapper"),
		html.Table(
			html.Class("table"),
			gomponents.Group(children),
		),
	)
}

// TableSm is the same as Table, but with smaller padding.
func TableSm(children ...gomponents.Node) gomponents.Node {
	return html.Div(
		html.Class("table-wrapper"),
		html.Table(
			html.Class("table table-sm"),
			gomponents.Group(children),
		),
	)
}

// TableLg is the same as Table, but with larger padding.
func TableLg(children ...gomponents.Node) gomponents.Node {
	return html.Div(
		html.Class("table-wrapper"),
		html.Table(
			html.Class("table table-lg"),
			gomponents.Group(children),
		),
	)
}
