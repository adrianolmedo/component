package components

import (
	"github.com/maragudk/gomponents"
	gcomponents "github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// HxLoading returns a loading indicator.
func HxLoading(centered bool, id ...string) gomponents.Node {
	return hxLoading(centered, "md", id...)
}

// HxLoadingSm returns a small loading indicator.
func HxLoadingSm(centered bool, id ...string) gomponents.Node {
	return hxLoading(centered, "sm", id...)
}

// HxLoadingLg returns a large loading indicator.
func HxLoadingLg(centered bool, id ...string) gomponents.Node {
	return hxLoading(centered, "lg", id...)
}

func hxLoading(centered bool, size string, id ...string) gomponents.Node {
	style := "width: 25px; height: 25px;"
	if size == "sm" {
		style = "width: 15px; height: 15px;"
	}
	if size == "lg" {
		style = "width: 40px; height: 40px;"
	}

	pickedID := ""
	if len(id) > 0 {
		pickedID = id[0]
	}

	return html.Div(
		gomponents.If(
			pickedID != "",
			html.ID(pickedID),
		),
		html.Class("htmx-indicator"),
		html.Div(
			gcomponents.Classes{
				"flex justify-center items-center": centered,
				"w-full h-full":                    true,
			},
			html.Div(
				html.Style(style),
				gcomponents.Classes{
					"border-primary border-t-transparent animate-spin rounded-full": true,
					"border-[2px]": size == "sm",
					"border-[3px]": size == "md",
					"border-[5px]": size == "lg",
				},
			),
		),
	)
}

// HxResWrapper is a wrapper for htmx responses.
// It is hidden by default and only shown when there is content inside
// of it. This avoids problems with margins and padding when the
// response is empty.
//
// https://developer.mozilla.org/en-US/docs/Web/CSS/:empty
func HxResWrapper(id string) gomponents.Node {
	return html.Div(
		html.Class("empty:hidden"),
		html.ID(id),
	)
}
