// Package goro provides a method to add 2 integers
package goro

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer
	constraints.Float
}

// Add function allows you to add 2 integers and get a resulting integer
// There is a link on adding [Math Addition]: https://www.mathsisfun.com/numbers/addition.html
func Add(num1, num2 Number) Number {
	switch num1.(type) {
	case int:
		return num1.(int) + num2.(int)
	case float64:
		return num1.(float64) + num2.(float64)
	}
	return 0
}
