package methods

import "strings"

type MyInt int
type MyString string

func (i MyInt) CheckInt() bool {
	return i%2 == 0
}

func (s MyString) CheckString() string {
	return strings.ToUpper(string(s))
}
