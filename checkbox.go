package components

import (
	"github.com/adrianolmedo/components/randutil"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type CheckboxProps struct {
	// Classes is the additional class name(s) for the input.
	Classes string
	// Color can be primary, secondary, neutral, info, success, warning, error.
	// Default is neutral.
	Color string
	// Size can be sm, md, lg. Default is md.
	Size string
	// Children are the contents of the input.
	Children []gomponents.Node
}

func Checkbox(props CheckboxProps) gomponents.Node {
	color := "neutral"
	if props.Color != "" {
		color = props.Color
	}

	size := "md"
	if props.Size != "" {
		size = props.Size
	}

	return html.Input(
		html.Type("checkbox"),
		components.Classes{
			"rounded bg-base-100 cursor-pointer":                   true,
			"focus:ring focus:ring-opacity-50 focus:ring-offset-0": true,

			"w-4 h-4": size == "sm",
			"w-6 h-6": size == "md",
			"w-9 h-9": size == "lg",

			"focus:ring-primary text-primary":     color == "primary",
			"focus:ring-secondary text-secondary": color == "secondary",
			"focus:ring-neutral text-neutral":     color == "neutral",
			"focus:ring-info text-info":           color == "info",
			"focus:ring-success text-success":     color == "success",
			"focus:ring-warning text-warning":     color == "warning",
			"focus:ring-error text-error":         color == "error",

			props.Classes: props.Classes != "",
		},
		gomponents.If(
			props.Children != nil,
			gomponents.Group(props.Children),
		),
	)
}

// CheckboxControlProps is the properties for the CheckboxControl component.
type CheckboxControlProps struct {
	// ID is the HTML id attribute value for the checkbox control.
	ID string
	// Name is the HTML name attribute value for the checkbox control.
	Name string
	// Title is the text label associated with the checkbox control.
	Title string
	// TitleNodes is the text label associated with the checkbox control.
	// Use this instead of Title if you need to include custom HTML.
	TitleNodes []gomponents.Node
	// HelpText is additional text displayed below the checkbox, usually providing guidance.
	HelpText string
	// Color is the color class applied to the checkbox control for styling.
	Color string
	// Align is the alignment of the checkbox control. Can be left (default) or right.
	Align string
	// Required, when true, indicates that the checkbox is mandatory for form submission.
	Required bool
	// Error, when true, applies error styling to the checkbox control.
	Error bool
	// CheckboxChildren are additional HTML nodes to be included inside the checkbox element.
	CheckboxChildren []gomponents.Node
	// CheckboxClasses are additional CSS classes to be applied to the checkbox element.
	CheckboxClasses string
	// ContainerClasses are additional CSS classes to be applied to the checkbox control container.
	ContainerClasses string
}

// CheckboxControl is a component that renders a group of nodes
// that are used to present a form checkbox control.
func CheckboxControl(props CheckboxControlProps) gomponents.Node {
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

	align := "left"
	if props.Align != "" {
		align = props.Align
	}

	checkbox := Checkbox(CheckboxProps{
		Color:   color,
		Classes: props.CheckboxClasses,
		Children: []gomponents.Node{
			html.ID(id),
			html.Name(props.Name),
			gomponents.If(props.Required, html.Required()),
			gomponents.If(hasHelperText, html.Aria("describedby", describedById)),
			gomponents.Group(props.CheckboxChildren),
		},
	})

	return html.Div(
		components.Classes{
			"checkbox-control-wrapper": true,
			"flex items-center":        true,
			"justify-start":            align == "left",
			"justify-end":              align == "right",
		},

		html.Label(
			components.Classes{
				"checkbox-control-container": true,
				"space-y-1":                  true,
				props.ContainerClasses:       props.ContainerClasses != "",
			},
			html.For(id),

			html.Div(
				html.Class("flex justify-start items-center space-x-2"),

				gomponents.If(align == "left", checkbox),

				html.Span(
					components.Classes{
						"checkbox-control-label block": true,
						"text-error":                   props.Error,
					},
					gomponents.If(
						len(props.TitleNodes) > 0,
						gomponents.Group(props.TitleNodes),
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

				gomponents.If(align == "right", checkbox),
			),

			gomponents.If(hasHelperText,
				html.Span(
					components.Classes{
						"checkbox-control-help-text block text-sm": true,
						"text-error": props.Error,
					},
					html.ID(describedById),
					gomponents.Text(props.HelpText),
				),
			),
		),
	)
}
