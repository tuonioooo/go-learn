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
	"os"
	"strconv"
)

func main() {
	if n, err := strconv.Atoi("42"); err == nil {
		// n and err are available here
		fmt.Println("There was no error, n is", n)
	}

	// n and err are not available here
	// fmt.Println(n, err)

	// UNCOMMENT THIS TO SEE IT IN ACTION:
	// var n int

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		// else if here shadows the main func's n
		// variable and it declares it own
		// with the same name: n

		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		fmt.Printf("%s * 2 is %d\n", a[1], n*2)
	}

	// n 的作用域在这里不起作用
	// fmt.Printf("n is %d. 👻 👻 👻 - you've been shadowed ;-)\n", n)

}
