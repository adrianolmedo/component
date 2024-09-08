package components

import (
	"fmt"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type DropdownButtonProps struct {
	// Direction is the direction of the dropdown. If empty, it will be "left".
	Direction string
	// Content is the content of the dropdown.
	Content []gomponents.Node
	// Width is the width of the dropdown. If empty, it will be "300px".
	Width string
	// ButtonProps are the props for the button that opens the dropdown.
	ButtonProps ButtonProps
}

func DropdownButton(props DropdownButtonProps) gomponents.Node {
	direction := "left"
	if props.Direction != "" {
		direction = props.Direction
	}

	width := "300px"
	if props.Width != "" {
		width = props.Width
	}

	return html.Div(
		gomponents.Attr("x-data", "{ open: false }"),
		html.Class("relative inline-block text-left"),

		html.Div(
			html.Class("flex items-center justify-center"),
			Button(ButtonProps{
				Classes: props.ButtonProps.Classes,
				Color:   props.ButtonProps.Color,
				Size:    props.ButtonProps.Size,
				Block:   props.ButtonProps.Block,
				Outline: props.ButtonProps.Outline,
				Ghost:   props.ButtonProps.Ghost,
				Square:  props.ButtonProps.Square,
				Circle:  props.ButtonProps.Circle,
				Children: []gomponents.Node{
					gomponents.Group(props.ButtonProps.Children),
					gomponents.Attr("x-on:click", "open = !open"),
				},
			}),
		),

		html.Div(
			gomponents.Attr("x-show", "open"),
			gomponents.Attr("x-on:click.outside", "open = false"),
			gomponents.Attr("x-transition"),
			gomponents.Attr("x-cloak"),

			html.Style(fmt.Sprintf("width: %s;", width)),
			components.Classes{
				"absolute z-20 mt-2 p-4":                     true,
				"rounded bg-base-100 border border-base-300": true,
				"right-0": direction == "left",
				"left-0":  direction == "right",
			},

			html.Div(
				html.Class("py-1"),
				html.Role("none"),
				gomponents.Group(props.Content),
			),
		),
	)
}
