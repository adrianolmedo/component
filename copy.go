package components

import (
	"fmt"
	"strings"

	"github.com/adrianolmedo/components/randutil"

	lucide "github.com/eduardolat/gomponents-lucide"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

// CopyButtonSimple is the same as CopyButton, but with a less
// verbose api and default props.
func CopyButtonSimple(textToCopy string) gomponents.Node {
	return CopyButton(CopyButtonProps{
		TextToCopy: textToCopy,
	})
}

// CopyButtonProps is the properties for the CopyButton component.
type CopyButtonProps struct {
	// TextToCopy is the text that should be copied to the
	// clipboard when the button is clicked.
	TextToCopy string
	// Size is the size of the button. Can be "sm", "md" (default) or "lg".
	Size Size
}

// CopyButton is a button that copies text to the clipboard when clicked.
func CopyButton(props CopyButtonProps) gomponents.Node {
	id := randutil.UUIDStrWithoutDashes()

	sc := copyButtonScript(id, props.TextToCopy)

	return html.Div(
		html.Class("inline-block"),
		sc.script,
		Button(ButtonProps{
			Color:  "neutral",
			Size:   props.Size,
			Square: true,
			Children: []gomponents.Node{
				html.ID(id),
				html.Title("Copiar"),
				sc.copyEvent,
				lucide.Copy(html.Class("size-4")),
			},
		}),
	)
}

// copyButtonScript returns a script that copies the given text to the clipboard
// and an event that calls the script when clicked.
func copyButtonScript(
	id string,
	textToCopy string,
) struct {
	script    gomponents.Node
	copyEvent gomponents.Node
} {
	escapedTextToCopy := strings.ReplaceAll(textToCopy, "`", "\\`")

	rawScript := fmt.Sprintf(
		"<script>function copy%s(){ copyToClipboard(`%s`); }</script>",
		id,
		escapedTextToCopy,
	)

	script := gomponents.Raw(rawScript)
	copyEvent := gomponents.Attr("onclick", fmt.Sprintf("copy%s()", id))

	return struct {
		script    gomponents.Node
		copyEvent gomponents.Node
	}{
		script:    script,
		copyEvent: copyEvent,
	}
}
