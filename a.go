package components

import (
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// AProps are props for the A component.
type AProps struct {
	// Href is the link.
	Href string
	// Classes are additional classes.
	Classes string
	// Disabled is whether the link is disabled.
	Disabled bool
	// Children are the child nodes.
	Children []gomponents.Node
}

// A is a link with some default classes.
func A(props AProps) gomponents.Node {
	href := props.Href
	if props.Disabled {
		href = "javascript:void(0)"
	}

	return html.A(
		components.Classes{
			"text-neutral underline": true,
			"hover:text-secondary":   !props.Disabled,
			"cursor-not-allowed":     props.Disabled,
			props.Classes:            props.Classes != "",
		},
		html.Href(href),
		gomponents.Group(props.Children),
	)
}
