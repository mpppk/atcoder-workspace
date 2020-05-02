package lib

import "strings"

// PanicIfErrorExistは、与えられたerrorがnilでなかった場合panicを発生させます.
func PanicIfErrorExist(err error) {
	if err != nil {
		panic(err)
	}
}

// TrimSpaceAndNewLineCodeAndTab は、スペース,改行,タブをTrimします.
func TrimSpaceAndNewLineCodeAndTab(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return r == ' ' || r == '\r' || r == '\n' || r == '\t'
	})
}
