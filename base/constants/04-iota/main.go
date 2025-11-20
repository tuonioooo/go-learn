// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	fmt.Println("======with_iota=======")
	with_iota()
	fmt.Println("======with_iota_expressions=======")
	with_iota_expressions()
	fmt.Println("======blank_identifier=======")
	blank_identifier()
	fmt.Println("======blank_identifier_iota=======")
	blank_identifier_iota()
	fmt.Println("======blank_identifier_zhanwei_iota=======")
	blank_identifier_zhanwei_iota()
}

func with_iota() {
	const (
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday,
		saturday, sunday)
}

func with_iota_expressions() {
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday,
		saturday, sunday)
}

func blank_identifier() {
	const (
		EST = -5
		MST = -7
		PST = -8
	)

	fmt.Println(EST, MST, PST)

}

func blank_identifier_iota() {
	const (
		EST = -(5 + iota) // CORRECT: -5
		MST               // CORRECT: -6
		PST               // CORRECT: -7
	)

	fmt.Println(EST, MST, PST)
}

func blank_identifier_zhanwei_iota() {
	const (
		EST = -(5 + iota) // CORRECT: -5
		_
		MST // CORRECT: -7
		PST // CORRECT: -8
	)

	fmt.Println(EST, MST, PST)
}
