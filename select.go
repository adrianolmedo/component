package components

import (
	"github.com/adrianolmedo/components/randutil"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type SelectProps struct {
	// Classes is the additional class name(s) for the select.
	Classes string
	// Color can be primary, secondary, neutral, info, success, warning, error.
	// Default is neutral.
	Color string
	// Block makes the select 100% width. Default is false.
	Block bool
	// Placeholder is the text displayed when the select is empty.
	Placeholder string
	// Children are the contents of the select.
	Children []gomponents.Node
}

func Select(props SelectProps) gomponents.Node {
	color := "neutral"
	if props.Color != "" {
		color = props.Color
	}

	colorClasses := selectColorClasses(color)

	return html.Select(
		components.Classes{
			"rounded bg-base-100": true,
			"placeholder:text-base-content placeholder:text-opacity-70": true,
			"focus:ring focus:ring-opacity-50":                          true,
			"w-full":                                                    props.Block,
			colorClasses:                                                true,
			props.Classes:                                               props.Classes != "",
		},

		gomponents.If(
			props.Placeholder != "",
			html.Option(
				html.Value(""),
				html.Disabled(),
				html.Selected(),
				gomponents.Text(props.Placeholder),
			),
		),

		gomponents.Group(props.Children),
	)
}

func selectColorClasses(color string) string {
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

// SelectControlProps is the properties for the SelectControl component.
type SelectControlProps struct {
	// ID is the HTML id attribute value for the select control.
	ID string
	// Name is the HTML name attribute value for the select control.
	Name string
	// Title is the text label associated with the select control.
	Title string
	// Placeholder is the HTML placeholder attribute value, displayed when the select is empty.
	Placeholder string
	// HelpText is additional text displayed below the select, usually providing guidance.
	HelpText string
	// Color is the color class applied to the select control for styling.
	Color string
	// Required, when true, indicates that the select is mandatory for form submission.
	Required bool
	// Error, when true, applies error styling to the select control.
	Error bool
	// Multiple, when true, allows multiple options to be selected.
	Multiple bool
	// SelectChildren are additional HTML nodes to be included inside the select element.
	SelectChildren []gomponents.Node
	// SelectClasses are additional CSS classes to be applied to the select element.
	SelectClasses string
	// ContainerClasses are additional CSS classes to be applied to the select control container.
	ContainerClasses string
}

// SelectControl is a component that renders a group of nodes
// that are used to present a form select control.
func SelectControl(props SelectControlProps) gomponents.Node {
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
			"select-control-container block space-y-1": true,
			props.ContainerClasses:                     props.ContainerClasses != "",
		},
		html.For(id),

		html.Span(
			components.Classes{
				"select-control-label block": true,
				"text-error":                 props.Error,
			},
			gomponents.Text(props.Title),
			gomponents.If(props.Required, html.Span(
				html.Class("text-error"),
				gomponents.Text(" *"),
			)),
		),
		Select(SelectProps{
			Color:       color,
			Block:       true,
			Classes:     props.SelectClasses,
			Placeholder: props.Placeholder,
			Children: []gomponents.Node{
				html.ID(id),
				html.Name(props.Name),
				gomponents.If(props.Multiple, html.Multiple()),
				gomponents.If(props.Required, html.Required()),
				gomponents.If(hasHelperText, html.Aria("describedby", describedById)),
				gomponents.Group(props.SelectChildren),
			},
		}),
		gomponents.If(hasHelperText,
			html.Span(
				components.Classes{
					"select-control-help-text block text-sm": true,
					"text-error":                             props.Error,
				},
				html.ID(describedById),
				gomponents.Text(props.HelpText),
			),
		),
	)
}
