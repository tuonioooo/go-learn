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
	var n int

	// ALTERNATIVES:
	// n = n + 1
	// n += 1

	// BETTER:
	n++

	fmt.Println(n)

	m := 10

	// ALTERNATIVES:
	// n = n - 1
	// n -= 1

	// BETTER:
	m--

	fmt.Println(m)


	var counter int

	// following "statements" are correct:

	counter++ // 1
	counter++ // 2
	counter++ // 3
	counter-- // 2
	fmt.Printf("There are %d line(s) in the file\n",
		counter)

	// the following "expressions" are incorrect:

	// counter = 5+counter--
	// counter = ++counter + counter--


	width, height := 5., 12.

	// calculates the area of a rectangle
	area := width * height
	fmt.Printf("%gx%g=%g\n", width, height, area)

	area = area - 10 // decreases area by 10
	area = area + 10 // increases area by 10
	area = area * 2  // doubles the area
	area = area / 2  // divides the area by 2
	fmt.Printf("area=%g\n", area)

	// // ASSIGNMENT OPERATIONS
	area -= 10 // decreases area by 10
	area += 10 // increases area by 10
	area *= 2  // doubles the area
	area /= 2  // divides the area by 2
	fmt.Printf("area=%g\n", area)

	// finds the remainder of area variable
	// since: area is float, this won't work:
	// area %= 7

	// this works
	area = float64(int(area) % 7)
	fmt.Printf("area=%g\n", area)
}