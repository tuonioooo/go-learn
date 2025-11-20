// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "carl"

	// strings are made up of bytes
	// len function counts the bytes in a string value
	fmt.Println(len(name))

	// strings are made up of bytes

	// len function counts the bytes in a string value.
	//
	// This string literal contains unicode characters.
	//
	// And, unicode characters can be 1-4 bytes.
	// So, "İnanç" is 7 bytes long, not 5.
	//
	// İ = 2 bytes
	// n = 1 byte
	// a = 1 byte
	// n = 1 byte
	// ç = 2 bytes
	// TOTAL = 7 bytes
	names2 := "İnanç"
	fmt.Printf("%q is %d bytes\n", names2, len(names2))

	// To get the actual characters (or runes) inside
	// a utf-8 encoded string value, you should do this:
	fmt.Printf("%q is %d characters\n",
		names2, utf8.RuneCountInString(names2))
}