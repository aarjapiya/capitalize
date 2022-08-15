/*
	1 . Write a function with this signature:

	func CapitalizeEveryThirdAlphanumericChar(s string) string {
		filtered := make([]string, 0)

		j := 0
		for _, ch := range strings.ToLower(s) {
			if (j+1)%3 == 0 && j != 0 && unicode.IsLetter(ch) {
				filtered = append(filtered, strings.ToUpper(string(ch)))
				j++
			} else if unicode.IsLetter(ch) {
				filtered = append(filtered, string(ch))
				j++
			} else {
				filtered = append(filtered, string(ch))
			}
		}

		return strings.Join(filtered, "")
	}
	that capitalizes *only* every third alphanumeric character of a string, so that if I hand you
	Aspiration.com
	You hand me back
	asPirAtiOn.cOm
	Please note:
	- Characters other than each third should be downcased
	- For the purposes of counting and capitalizing every three characters, consider only alphanumeric
	  characters, ie [a-z][A-Z][0-9].
	2. Do the same problem, but this time, implement the below code:
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

type customString []rune

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

func (s customString) String() string {
	res := ""
	for _, ch := range s {
		res += string(ch)
	}
	return fmt.Sprintf("%v", res)
}

// And change your code to look like this:
func (s customString) TransformRune(pos int) {
	s[pos] = []rune(string(s[pos]))[0]
}

func (s customString) GetValueAsRuneSlice() []rune {
	res := make([]rune, 0)
	for _, ch := range s {
		res = append(res, ch)
	}
	return res
}

func NewSkipString(i int, s string) customString {
	var filtered customString

	j := 0
	for _, ch := range strings.ToLower(s) {
		if (j+1)%i == 0 && j != 0 && unicode.IsLetter(ch) {
			filtered = append(filtered, []rune(strings.ToUpper(string(ch)))[0])
			j++
		} else if unicode.IsLetter(ch) {
			filtered = append(filtered, ch)
			j++
		} else {
			filtered = append(filtered, ch)
		}
	}
	return filtered
}

func main() {
	s := NewSkipString(3, "Aspiration.com")
	MapString(&s)
	fmt.Println(s)
}

/*
   Make sure your fmt.Println(s) output looks nice and clean, ie implement the Stringer interface.
*/
