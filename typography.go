package components

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func H1(children ...gomponents.Node) gomponents.Node {
	return html.H1(
		html.Class("h1"),
		gomponents.Group(children),
	)
}

func H2(children ...gomponents.Node) gomponents.Node {
	return html.H2(
		html.Class("h2"),
		gomponents.Group(children),
	)
}

func H3(children ...gomponents.Node) gomponents.Node {
	return html.H3(
		html.Class("h3"),
		gomponents.Group(children),
	)
}

func H4(children ...gomponents.Node) gomponents.Node {
	return html.H4(
		html.Class("h4"),
		gomponents.Group(children),
	)
}

func H5(children ...gomponents.Node) gomponents.Node {
	return html.H5(
		html.Class("h5"),
		gomponents.Group(children),
	)
}

func H6(children ...gomponents.Node) gomponents.Node {
	return html.H6(
		html.Class("h6"),
		gomponents.Group(children),
	)
}

// H1Text is a convenience function to create an H1 element with a
// simple text node as its child.
func H1Text(text string) gomponents.Node {
	return H1(gomponents.Text(text))
}

// H2Text is a convenience function to create an H2 element with a
// simple text node as its child.
func H2Text(text string) gomponents.Node {
	return H2(gomponents.Text(text))
}

// H3Text is a convenience function to create an H3 element with a
// simple text node as its child.
func H3Text(text string) gomponents.Node {
	return H3(gomponents.Text(text))
}

// H4Text is a convenience function to create an H4 element with a
// simple text node as its child.
func H4Text(text string) gomponents.Node {
	return H4(gomponents.Text(text))
}

// H5Text is a convenience function to create an H5 element with a
// simple text node as its child.
func H5Text(text string) gomponents.Node {
	return H5(gomponents.Text(text))
}

// H6Text is a convenience function to create an H6 element with a
// simple text node as its child.
func H6Text(text string) gomponents.Node {
	return H6(gomponents.Text(text))
}

// PText is a convenience function to create a P element with a
// simple text node as its child.
func PText(text string) gomponents.Node {
	return html.P(gomponents.Text(text))
}

// SpanText is a convenience function to create a Span element with a
// simple text node as its child.
func SpanText(text string) gomponents.Node {
	return html.Span(gomponents.Text(text))
}

// DivText is a convenience function to create a Div element with a
// simple text node as its child.
func DivText(text string) gomponents.Node {
	return html.Div(gomponents.Text(text))
}

// TdText is a convenience function to create a Td element with a
// simple text node as its child.
func TdText(text string) gomponents.Node {
	return html.Td(gomponents.Text(text))
}

// ThText is a convenience function to create a Th element with a
// simple text node as its child.
func ThText(text string) gomponents.Node {
	return html.Th(gomponents.Text(text))
}
