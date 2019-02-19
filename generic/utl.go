package generic

import "strings"

func PanicIfErrorExist(err error) {
	if err != nil {
		panic(err)
	}
}

func TrimSpaceAndNewLineCodeAndTab(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return r == ' ' || r == '\r' || r == '\n' || r == '\t'
	})
}
