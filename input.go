package components

import (
	"github.com/adrianolmedo/components/randutil"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type InputProps struct {
	// Classes is the additional class name(s) for the input.
	Classes string
	// Color can be primary, secondary, neutral, info, success, warning, error.
	// Default is neutral.
	Color string
	// Block makes the input 100% width. Default is false.
	Block bool
	// Children are the contents of the input.
	Children []gomponents.Node
}

func Input(props InputProps) gomponents.Node {
	color := "neutral"
	if props.Color != "" {
		color = props.Color
	}

	colorClasses := inputColorClasses(color)

	return html.Input(
		components.Classes{
			"disabled:cursor-not-allowed":                               true,
			"rounded bg-base-100":                                       true,
			"placeholder:text-base-content placeholder:text-opacity-70": true,
			"focus:ring focus:ring-opacity-50":                          true,
			"w-full":                                                    props.Block,
			colorClasses:                                                true,
			props.Classes:                                               props.Classes != "",
		},
		gomponents.Group(props.Children),
	)
}

func inputColorClasses(color string) string {
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

// InputControlProps is the properties for the InputControl component.
type InputControlProps struct {
	// ID is the HTML id attribute value for the input control.
	ID string
	// Name is the HTML name attribute value for the input control.
	Name string
	// Title is the text label associated with the input control.
	Title string
	// TitleNode is the content inside the label associated with the input control.
	// If you need only a string, use Title instead.
	TitleNode gomponents.Node
	// Placeholder is the HTML placeholder attribute value, displayed when the input is empty.
	Placeholder string
	// HelpText is additional text displayed below the input, usually providing guidance.
	HelpText string
	// Color is the color class applied to the input control for styling.
	Color string
	// Required, when true, indicates that the input is mandatory for form submission.
	Required bool
	// Error, when true, applies error styling to the input control.
	Error bool
	// InputChildren are additional HTML nodes to be included inside the input element.
	InputChildren []gomponents.Node
	// InputClasses are additional CSS classes to be applied to the input element.
	InputClasses string
	// ContainerClasses are additional CSS classes to be applied to the input control container.
	ContainerClasses string
}

// InputControl is a component that renders a group of nodes
// that are used to present a form input control.
func InputControl(props InputControlProps) gomponents.Node {
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
			"input-control-container block space-y-1": true,
			props.ContainerClasses:                    props.ContainerClasses != "",
		},
		html.For(id),

		html.Span(
			components.Classes{
				"input-control-label block": true,
				"text-error":                props.Error,
			},

			gomponents.If(
				props.TitleNode != nil,
				props.TitleNode,
			),

			gomponents.If(
				props.Title != "",
				gomponents.Text(props.Title),
			),

			gomponents.If(props.Required, html.Span(
				html.Class("text-error"),
				gomponents.Text(" *"),
			)),
		),
		Input(InputProps{
			Color:   color,
			Block:   true,
			Classes: props.InputClasses,
			Children: []gomponents.Node{
				html.ID(id),
				html.Name(props.Name),
				html.Placeholder(props.Placeholder),
				gomponents.If(props.Required, html.Required()),
				gomponents.If(hasHelperText, html.Aria("describedby", describedById)),
				gomponents.Group(props.InputChildren),
			},
		}),
		gomponents.If(hasHelperText,
			html.Span(
				components.Classes{
					"input-control-help-text block text-sm": true,
					"text-error":                            props.Error,
				},
				html.ID(describedById),
				gomponents.Text(props.HelpText),
			),
		),
	)
}

// FileInputControlProps is the properties for the FileInputControl component.
type FileInputControlProps struct {
	// ID is the HTML id attribute value for the input control.
	ID string
	// Name is the HTML name attribute value for the input control.
	Name string
	// Title is the text label associated with the input control.
	Title string
	// HelpText is additional text displayed below the input, usually providing guidance.
	HelpText string
	// Required, when true, indicates that the input is mandatory for form submission.
	Required bool
	// Error, when true, applies error styling to the input control.
	Error bool
	// Accept is a comma-separated list of file types that the input control can accept.
	// If not set, the input control will accept all file types.
	Accept string
	// InputChildren are additional HTML nodes to be included inside the input element.
	InputChildren []gomponents.Node
	// InputClasses are additional CSS classes to be applied to the input element.
	InputClasses string
	// ContainerClasses are additional CSS classes to be applied to the input control container.
	ContainerClasses string
}

// FileInputControl is a component that renders a group of nodes
// that are used to present a form input control.
func FileInputControl(props FileInputControlProps) gomponents.Node {
	id := props.ID
	if id == "" {
		id = randutil.UUIDStrWithoutDashes()
	}

	hasAccept := props.Accept != ""
	hasHelperText := props.HelpText != ""
	describedById := id + "-help-text"

	return html.Label(
		components.Classes{
			"input-control-container block space-y-1": true,
			props.ContainerClasses:                    props.ContainerClasses != "",
		},
		html.For(id),

		html.Span(
			components.Classes{
				"input-control-label block": true,
				"text-error":                props.Error,
			},
			gomponents.Text(props.Title),
			gomponents.If(props.Required, html.Span(
				html.Class("text-error"),
				gomponents.Text(" *"),
			)),
		),
		Input(InputProps{
			Block:   true,
			Classes: props.InputClasses,
			Children: []gomponents.Node{
				html.ID(id),
				html.Name(props.Name),
				html.Type("file"),
				gomponents.If(hasAccept, html.Accept(props.Accept)),
				gomponents.If(props.Required, html.Required()),
				gomponents.If(hasHelperText, html.Aria("describedby", describedById)),
				gomponents.Group(props.InputChildren),
			},
		}),
		gomponents.If(hasHelperText,
			html.Span(
				components.Classes{
					"input-control-help-text block text-sm": true,
					"text-error":                            props.Error,
				},
				html.ID(describedById),
				gomponents.Text(props.HelpText),
			),
		),
	)
}
