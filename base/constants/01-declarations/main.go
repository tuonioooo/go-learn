// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// This program uses magic values

func main() {

	// 常量定义
	const meters int = 100

	cm := 100
	m := cm / meters // using a named constant
	fmt.Printf("%dcm is %dm\n", cm, m)

	cm = 200
	m = cm / meters // using a named constant
	fmt.Printf("%dcm is %dm\n", cm, m)

	// When argument to len is a constant, len can be used
	// while initializing a constant
	//
	// Here, "Hello" is a constant.

	const max int = len("Hello")

	fmt.Println(max)


	fmt.Println("==========test1=============")
	test1()
	fmt.Println("==========test2=============")
	test2()
	fmt.Println("==========test3=============")
	test3()
	fmt.Println("======multiple_declarations======")
	multiple_declarations()
	fmt.Println("======const_inherit======")
	const_inherit()
}



func test1() {
	const min int = 1
	const pi float64 = 3.14
	const version string = "2.0.1"
	const debug bool = true
	const A rune = 'A'

	fmt.Println(min, pi, version, debug, A)

}

func test2() {
	const min = 1
	const pi = 3.14
	const version = "2.0.1"
	const debug = true
	const A = 'A'

	fmt.Println(min, pi, version, debug, A)

}

func test3() {
	const min = 1 + 1
	const pi = 3.14 * min
	const version = "2.0.1" + "-beta"
	const debug = !true
	const A rune = 'A' + 1

	fmt.Println(min, pi, version, debug, A)
}

func multiple_declarations() {
	// below declaration is the same as this one:

	// const (
	// 	min int = 1
	// 	max int = 1000
	// )

	const min, max int = 1, 1000

	fmt.Println(min, max)

	// print the types of min and max
	fmt.Printf("%T %T\n", min, max)

}

func const_inherit() {
	// constants repeat the previous type and expression
	const (
		min int = 1 + 1
		max     // int = 1 + 1
	)

	fmt.Println(min, max)

	// print the types of min and max
	fmt.Printf("%T %T\n", min, max)

}