package components

import (
	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// CollapseProps are the properties of the collapse component.
type CollapseProps struct {
	// Title is the title string of the collapse.
	// If you need to add more than just a string, use
	// TitleNode instead.
	Title string
	// TitleNode is the title node of the collapse.
	// Use this if you need to add more than just a string.
	TitleNode gomponents.Node
	// Color is the color of the collapse. Can be "light" or "dark". Default is "light".
	Color string
	// Content is the content of the collapse.
	Content []gomponents.Node
}

// Collapse renders a collapse component.
func Collapse(props CollapseProps) gomponents.Node {
	color := "light"
	if props.Color != "" {
		color = props.Color
	}

	return html.Div(
		gomponents.Attr("x-data", "{ open: false }"),
		components.Classes{
			"collapse-component w-full rounded p-4":    true,
			"border border-base-300 text-base-content": true,
			"bg-base-100": color == "light",
			"bg-base-200": color == "dark",
		},

		html.Div(
			gomponents.Attr("x-on:click", "open = !open"),
			components.Classes{
				"collapse-component-title":          true,
				"cursor-pointer w-full select-none": true,
				"flex justify-start items-center":   true,
			},

			html.Div(
				html.Span(
					gomponents.Attr("x-show", "!open"),
					lucide.CirclePlus(html.Class("w-6 h-6 mr-2")),
				),
				html.Span(
					gomponents.Attr("x-show", "open"),
					gomponents.Attr("x-cloak"),
					lucide.CircleMinus(html.Class("w-6 h-6 mr-2")),
				),
			),

			gomponents.If(
				props.Title != "",
				html.Span(
					html.Class("font-bold text-lg"),
					gomponents.Text(props.Title),
				),
			),
			gomponents.If(
				props.TitleNode != nil,
				props.TitleNode,
			),
		),

		html.Div(
			gomponents.Attr("x-show", "open"),
			gomponents.Attr("x-transition"),
			gomponents.Attr("x-cloak"),

			components.Classes{
				"collapse-component-content":         true,
				"border-t border-base-300 mt-4 pt-4": true,
			},

			gomponents.Group(props.Content),
		),
	)
}
