package components

import (
	"fmt"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// ButtonProps are the props for a button.
type ButtonProps struct {
	// Classes is the additional class name(s) for the button.
	Classes string
	// Color can be primary, secondary, neutral, info, success, warning, error.
	// Default is neutral.
	Color string
	// Size can be sm, md, lg. Default is md.
	Size Size
	// Block makes the button 100% width. Default is false.
	Block bool
	// Outline makes the button outlined. Default is false.
	Outline bool
	// Ghost makes the button ghost. Default is false.
	Ghost bool
	// Square makes the button square. Default is false.
	Square bool
	// Circle makes the button circular. Default is false.
	Circle bool
	// Children are the contents of the button.
	Children []gomponents.Node
}

// Button renders a button.
func Button(props ButtonProps) gomponents.Node {
	color := "neutral"
	if props.Color != "" {
		color = props.Color
	}

	size := SizeMd
	if props.Size != "" {
		size = props.Size
	}

	block := props.Block
	outline := props.Outline
	ghost := props.Ghost

	colorClasses := buttonColorRegularClasses(color)
	if outline {
		colorClasses = buttonColorOutlineClasses(color)
	}
	if ghost {
		colorClasses = buttonColorGhostClasses(color)
	}

	sizeClasses := buttonSizeRegularClasses(size)
	if props.Square {
		sizeClasses = buttonSizeSquareClasses(size)
	}
	if props.Circle {
		sizeClasses = buttonSizeCircleClasses(size)
	}

	return html.Button(
		components.Classes{
			"disabled:cursor-not-allowed disabled:opacity-50": true,
			"rounded transition duration-100":                 true,
			"inline-flex justify-center items-center":         true,
			"active:scale-95":                                 true,
			"w-full":                                          block,
			colorClasses:                                      true,
			sizeClasses:                                       true,
			props.Classes:                                     props.Classes != "",
		},
		gomponents.Group(props.Children),
	)
}

func buttonColorRegularClasses(color string) string {
	if color == "primary" {
		return "bg-primary text-primary-content"
	}

	if color == "secondary" {
		return "bg-secondary text-secondary-content"
	}

	if color == "neutral" {
		return "bg-neutral text-neutral-content"
	}

	if color == "info" {
		return "bg-info text-info-content"
	}

	if color == "success" {
		return "bg-success text-success-content"
	}

	if color == "warning" {
		return "bg-warning text-warning-content"
	}

	if color == "error" {
		return "bg-error text-error-content"
	}

	return ""
}

func buttonColorOutlineClasses(color string) string {
	if color == "primary" {
		return "border border-primary text-primary"
	}

	if color == "secondary" {
		return "border border-secondary text-secondary"
	}

	if color == "neutral" {
		return "border border-neutral text-neutral"
	}

	if color == "info" {
		return "border border-info text-info"
	}

	if color == "success" {
		return "border border-success text-success"
	}

	if color == "warning" {
		return "border border-warning text-warning"
	}

	if color == "error" {
		return "border border-error text-error"
	}

	return ""
}

func buttonColorGhostClasses(color string) string {
	if color == "primary" {
		return "text-primary hover:bg-primary hover:bg-opacity-10"
	}

	if color == "secondary" {
		return "text-secondary hover:bg-secondary hover:bg-opacity-10"
	}

	if color == "neutral" {
		return "text-neutral hover:bg-neutral hover:bg-opacity-10"
	}

	if color == "info" {
		return "text-info hover:bg-info hover:bg-opacity-10"
	}

	if color == "success" {
		return "text-success hover:bg-success hover:bg-opacity-10"
	}

	if color == "warning" {
		return "text-warning hover:bg-warning hover:bg-opacity-10"
	}

	if color == "error" {
		return "text-error hover:bg-error hover:bg-opacity-10"
	}

	return ""
}

func buttonSizeRegularClasses(size Size) string {
	if size == "sm" {
		return "px-2 py-1 text-sm"
	}

	if size == "md" {
		return "px-4 py-2 text-md"
	}

	if size == "lg" {
		return "px-6 py-3 text-lg"
	}

	return ""
}

func buttonSizeSquareClasses(size Size) string {
	if size == "sm" {
		return "w-6 h-6"
	}

	if size == "md" {
		return "w-8 h-8"
	}

	if size == "lg" {
		return "w-10 h-10"
	}

	return ""
}

func buttonSizeCircleClasses(size Size) string {
	return fmt.Sprintf("rounded-full %s", buttonSizeSquareClasses(size))
}
