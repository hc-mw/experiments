package fooer

import "strconv"

func Fooer(n int) string {
	if n%3 == 0 {
		return "Foo"
	}
	return strconv.Itoa(n)
}
