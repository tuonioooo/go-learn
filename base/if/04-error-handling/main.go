package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("==========demo1=============")
	demo1()
	fmt.Println("==========demo2=============")
	demo2([]string{"cmd", "123"})
	fmt.Println("==========demo3=============")
	demo3([]string{"cmd", "abc"})
	fmt.Println("==========demo4=============")
	fmt.Println("----------")
	demo4([]string{"cmd", "abc"})
	fmt.Println("----------")
	demo4([]string{"cmd", "10.5"})
}

func demo1() {
	// Itoa doesn't return any errors
	// So, you don't have to handle the errors for it

	s := strconv.Itoa(42)

	fmt.Println(s)
}

func demo2(osArgs []string) {
	// Atoi returns an error value
	// So, you should always check it

	n, err := strconv.Atoi(osArgs[1])
	fmt.Println("Converted number    :", n)
	fmt.Println("Returned error value:", err)
}

func demo3(osArgs []string) {

	age := osArgs[1]

	// Atoi returns an int and an error value
	// So, you need to handle the errors

	n, err := strconv.Atoi(age)

	// handle the error immediately and quit
	// don't do anything else here

	if err != nil {
		fmt.Println("ERROR:", err)

		// quits/returns from the main function
		// so, the program ends
		return
	}

	// DO NOT DO THIS:
	// else {
	//   happy path
	// }

	// DO THIS:

	// happy path (err is nil)
	fmt.Printf("SUCCESS: Converted %q to %d.\n", age, n)
}

const usage = `
Feet to Meters
--------------
This program converts feet into meters.

Usage:
feet [feetsToConvert]`

func demo4(osArgs []string) {

	if len(osArgs) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}

	arg := osArgs[1]

	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("error: '%s' is not a number.\n", arg)
		return
	}

	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
