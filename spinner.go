package components

import (
	"fmt"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func Spinner(size string) gomponents.Node {
	return lucide.LoaderCircle(Classes{
		"animate-spin inline-block": true,
		"w-5 h-5":                   size == "sm",
		"w-8 h-8":                   size == "md",
		"w-12 h-12":                 size == "lg",
	})
}

func SpinnerSm() gomponents.Node {
	return Spinner("sm")
}

func SpinnerMd() gomponents.Node {
	return Spinner("md")
}

func SpinnerLg() gomponents.Node {
	return Spinner("lg")
}

func SpinnerContainer(size string, height string) gomponents.Node {
	return html.Div(
		Classes{
			"flex justify-center": true,
			"items-center w-full": true,
		},
		html.Style(fmt.Sprintf("height: %s;", height)),
		Spinner(size),
	)
}

func SpinnerContainerSm() gomponents.Node {
	return SpinnerContainer("sm", "300px")
}

func SpinnerContainerMd() gomponents.Node {
	return SpinnerContainer("md", "300px")
}

func SpinnerContainerLg() gomponents.Node {
	return SpinnerContainer("lg", "300px")
}
