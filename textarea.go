package components

import (
	"github.com/adrianolmedo/components/randutil"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type TextareaProps struct {
	// Classes is the additional class name(s) for the textarea.
	Classes string
	// Color can be primary, secondary, neutral, info, success, warning, error.
	// Default is neutral.
	Color string
	// Block makes the textarea 100% width. Default is false.
	Block bool
	// Children are the contents of the textarea.
	Children []gomponents.Node
}

func Textarea(props TextareaProps) gomponents.Node {
	color := "neutral"
	if props.Color != "" {
		color = props.Color
	}

	colorClasses := textareaColorClasses(color)

	return html.Textarea(
		components.Classes{
			"rounded bg-base-100": true,
			"placeholder:text-base-content placeholder:text-opacity-70": true,
			"focus:ring focus:ring-opacity-50":                          true,
			"w-full":                                                    props.Block,
			colorClasses:                                                true,
			props.Classes:                                               props.Classes != "",
		},
		gomponents.Group(props.Children),
	)
}

func textareaColorClasses(color string) string {
	if color == "primary" {
		return "border-primary focus:border-primary focus:ring-primary"
	}

	if color == "secondary" {
		return "border-secondary focus:border-secondary focus:ring-secondary"
	}

	if color == "neutral" {
		return "border-neutral focus:border-neutral focus:ring-neutral"
	}

	if color == "info" {
		return "border-info focus:border-info focus:ring-info"
	}

	if color == "success" {
		return "border-success focus:border-success focus:ring-success"
	}

	if color == "warning" {
		return "border-warning focus:border-warning focus:ring-warning"
	}

	if color == "error" {
		return "border-error focus:border-error focus:ring-error"
	}

	return ""
}

// TextareaControlProps is the properties for the TextareaControl component.
type TextareaControlProps struct {
	// ID is the HTML id attribute value for the textarea control.
	ID string
	// Name is the HTML name attribute value for the textarea control.
	Name string
	// Title is the text label associated with the textarea control.
	Title string
	// Placeholder is the HTML placeholder attribute value, displayed when the textarea is empty.
	Placeholder string
	// HelpText is additional text displayed below the textarea, usually providing guidance.
	HelpText string
	// Color is the color class applied to the textarea control for styling.
	Color string
	// Required, when true, indicates that the textarea is mandatory for form submission.
	Required bool
	// Error, when true, applies error styling to the textarea control.
	Error bool
	// TextareaChildren are additional HTML nodes to be included inside the textarea element.
	TextareaChildren []gomponents.Node
	// TextareaClasses are additional CSS classes to be applied to the textarea element.
	TextareaClasses string
	// ContainerClasses are additional CSS classes to be applied to the textarea control container.
	ContainerClasses string
}

// TextareaControl is a component that renders a group of nodes
// that are used to present a form textarea control.
func TextareaControl(props TextareaControlProps) gomponents.Node {
	id := props.ID
	if id == "" {
		id = randutil.UUIDStrWithoutDashes()
	}

	hasHelperText := props.HelpText != ""
	describedById := id + "-help-text"

	color := props.Color
	if props.Error {
		color = "error"
	}

	return html.Label(
		components.Classes{
			"textarea-control-container block space-y-1": true,
			props.ContainerClasses:                       props.ContainerClasses != "",
		},
		html.For(id),

		html.Span(
			components.Classes{
				"textarea-control-label block": true,
				"text-error":                   props.Error,
			},
			gomponents.Text(props.Title),
			gomponents.If(props.Required, html.Span(
				html.Class("text-error"),
				gomponents.Text(" *"),
			)),
		),
		Textarea(TextareaProps{
			Color:   color,
			Block:   true,
			Classes: props.TextareaClasses,
			Children: []gomponents.Node{
				html.ID(id),
				html.Name(props.Name),
				html.Placeholder(props.Placeholder),
				gomponents.If(props.Required, html.Required()),
				gomponents.If(hasHelperText, html.Aria("describedby", describedById)),
				gomponents.Group(props.TextareaChildren),
			},
		}),
		gomponents.If(hasHelperText,
			html.Span(
				components.Classes{
					"textarea-control-help-text block text-sm": true,
					"text-error": props.Error,
				},
				html.ID(describedById),
				gomponents.Text(props.HelpText),
			),
		),
	)
}
