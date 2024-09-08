package components

import (
	"github.com/adrianolmedo/components/gomputils"
	"github.com/adrianolmedo/components/randutil"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// PushbarProps is the config props for the Pushbar component.
type PushbarProps struct {
	// ID is the ID of the modal dialog. If empty, a random ID will be generated.
	ID string
	// Content is the content of the pushbar.
	Content gomponents.Node
	// DisableCloseButton hides the close button.
	DisableCloseButton bool
	// Direction is the direction of the pushbar.
	// Can be 'left', 'right', 'top' or 'bottom'.
	// Defaults to 'right'.
	Direction string
}

// PushbarResult is the result of creating a pushbar.
type PushbarResult struct {
	// HTML is the pushbar HTML.
	HTML gomponents.Node
	// OpenerAttr is the attribute to add to the element that opens the pushbar.
	OpenerAttr gomponents.Node
}

// Pushbar creates a pushbar.
func Pushbar(props PushbarProps) PushbarResult {
	id := props.ID
	if id == "" {
		id = randutil.UUIDStrWithoutDashes()
	}

	direction := "right"
	if props.Direction != "" {
		direction = props.Direction
	}

	openerAttr := gomponents.Attr("data-pushbar-target", id)

	htmlContent := html.Div(
		html.ID(id),
		gomponents.Attr("data-pushbar-id", id),
		gomponents.Attr("data-pushbar-direction", direction),

		// Close button
		gomponents.If(
			!props.DisableCloseButton,
			html.Div(
				gomputils.ManyClasses(
					"sticky left-0 top-0 z-50 flex justify-end border-b-2",
					"border-base-content border-opacity-10 bg-base-100 p-2",
				),
				Button(ButtonProps{
					Ghost:  true,
					Square: true,
					Children: []gomponents.Node{
						gomponents.Attr("data-pushbar-close"),
						lucide.X(
							html.Class("w-5 h-5"),
						),
					},
				}),
			),
		),

		// Content
		html.Div(
			html.Class("p-2"),
			props.Content,
		),
	)

	return PushbarResult{
		HTML:       htmlContent,
		OpenerAttr: openerAttr,
	}
}
