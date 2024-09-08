package components

import "github.com/maragudk/gomponents"

// Eval is a component that accepts a function, does the evaluation
// and returns the result.
func Eval(fn func() gomponents.Node) gomponents.Node {
	return fn()
}
