/*
User interface contains NumericInput control, which accepts only digits.

Extend NumericInput structure so that:

It implements UserInput interface.
Add(rune) should add only decimal digits to the input. Other runes should be ignored.
GetValue() should return the current input.
For example, the following code should output "10":

var input UserInput = &NumericInput{}
input.Add('1')
input.Add('a')
input.Add('0')
fmt.Println(input.GetValue())
*/
package main

import "fmt"

type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

func (n *NumericInput) Add(inp rune) {
	for i := 48; i <= 57; i++ {
		if i == int(inp) {
			n.input += string(inp)
		}
	}
}

func (n *NumericInput) GetValue() string {
	return n.input
}
func main() {
	var input UserInput = &NumericInput{}
	input.Add('1')
	input.Add('a')
	input.Add('0')
	fmt.Println(input.GetValue())
}
