package components

import (
	"strconv"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// TinyCounterProps are the props for the tiny counter component.
type TinyCounterProps struct {
	Value               int
	MinusButtonChildren []gomponents.Node
	PlusButtonChildren  []gomponents.Node
}

// TinyCounter is a tiny counter component.
func TinyCounter(props TinyCounterProps) gomponents.Node {
	return html.Div(
		html.Class("flex justify-center items-center"),

		html.Div(
			html.Class("flex justify-center items-center rounded border border-base-300 bg-base-100"),

			Button(ButtonProps{
				Square: true,
				Children: []gomponents.Node{
					lucide.Minus(html.Class("w-4 h-4")),
					gomponents.Group(props.MinusButtonChildren),
				},
			}),
			html.Div(
				html.Class("w-[30px] h-full flex justify-center items-center"),
				html.Span(
					html.Class("font-bold text-lg"),
					gomponents.Text(
						strconv.Itoa(props.Value),
					),
				),
			),
			Button(ButtonProps{
				Square: true,
				Children: []gomponents.Node{
					lucide.Plus(html.Class("w-4 h-4")),
					gomponents.Group(props.PlusButtonChildren),
				},
			}),
		),
	)
}
