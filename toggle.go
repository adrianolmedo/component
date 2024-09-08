package components

import (
	"github.com/adrianolmedo/components/alpine"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/google/uuid"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type ToggleParams struct {
	// Color can be primary, secondary, neutral, info, success, warning, error.
	// Default is neutral.
	Color string
	// Size can be sm, md, lg. Default is md.
	Size string
	// Name is the name of the input.
	Name string
	// UnCheckedValue is the value when the toggle is off. Default is "off".
	UnCheckedValue string
	// CheckedValue is the value when the toggle is on. Default is "on".
	CheckedValue string
	// Value is the initial value of the toggle. Default is UnCheckedValue.
	Value string
	// ID is the id of the input. Default is random.
	ID string
	// InputChildren are the children of the input.
	InputChildren []gomponents.Node
}

func Toggle(params ToggleParams) gomponents.Node {
	color := "neutral"
	size := "md"
	name := params.Name
	uncheckedValue := params.UnCheckedValue
	checkedValue := params.CheckedValue

	if params.Color != "" {
		color = params.Color
	}

	if params.Size != "" {
		size = params.Size
	}

	if uncheckedValue == "" {
		uncheckedValue = "off"
	}

	if checkedValue == "" {
		checkedValue = "on"
	}

	value := uncheckedValue
	if params.Value == checkedValue {
		value = checkedValue
	}

	id := params.ID
	if params.ID == "" {
		id = "toggle-" + uuid.NewString()
	}

	iconClasses := Classes{
		"w-8 h-8":   size == "sm",
		"w-10 h-10": size == "md",
		"w-14 h-14": size == "lg",

		"text-primary":   color == "primary",
		"text-secondary": color == "secondary",
		"text-neutral":   color == "neutral",
		"text-info":      color == "info",
		"text-success":   color == "success",
		"text-warning":   color == "warning",
		"text-error":     color == "error",
	}

	return html.Div(
		html.Class("inline-block select-none"),

		html.Input(
			html.Class("hidden"),
			html.Type("hidden"),
			html.Name(name),
			html.ID(id),
			gomponents.Group(params.InputChildren),
		),

		html.Div(
			alpine.XData(`{
				value: "`+value+`",
				uncheckedValue: "`+uncheckedValue+`",
				checkedValue: "`+checkedValue+`",
				isChecked: false,
				setChecked() {
					this.isChecked = this.value === this.checkedValue
				},
				changeValue() {
					if (this.isChecked) {
						this.value = this.uncheckedValue
					} else {
						this.value = this.checkedValue
					}
				},
				setInputValue() {
					document.getElementById("`+id+`").value = this.value
					document.getElementById("`+id+`").dispatchEvent(new Event("change"))
				},
				toggle() {
					this.changeValue()
					this.setChecked()
					this.setInputValue()
				},
				init() {
					this.setChecked()
					this.setInputValue()
				}
			}`),

			html.Div(
				html.Class("cursor-pointer"),
				alpine.XOn("click", "toggle()"),
				alpine.TemplateEl(
					alpine.XIf("!isChecked"),
					lucide.ToggleLeft(iconClasses),
				),
				alpine.TemplateEl(
					alpine.XIf("isChecked"),
					lucide.ToggleRight(iconClasses),
				),
			),
		),
	)
}
