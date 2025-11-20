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
	fmt.Println(
		2+2*4/2,
		2+((2*4)/2), // same as above
	)

	fmt.Println(
		1+4-2,
		(1+4)-2, // same as above
	)

	fmt.Println(
		(2+2)*4/2,
		(2+2)*(4/2), // same as above
	)

	n, m := 1, 5

	fmt.Println(2 + 1*m/n)
	fmt.Println(2 + ((1 * m) / n)) // same as above

	// let's change the precedence using parentheses
	fmt.Println(((2 + 1) * m) / n)

	celsius := 35.

	// Wrong formula  :  9*celsius + 160  / 5
	// Correct formula: (9*celsius + 160) / 5
	fahrenheit := (9*celsius + 160) / 5

	fmt.Printf("%g ºC is %g ºF\n", celsius, fahrenheit)
}