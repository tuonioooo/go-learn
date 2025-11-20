package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = `
Feet to Meters
--------------
This program converts feet into meters.

Usage:
feet [feetsToConvert]`

func main() {
	fmt.Println("==========demo1=============")
	demo1()
	fmt.Println("==========demo2=============")
	demo2()
	fmt.Println("==========demo3=============")
	demo3()
	fmt.Println("==========demo4=============")
	demo4(os.Args)
}

func demo1() {
	score, valid := 5, true

	if score > 3 && valid {
		fmt.Println("good")
	}
}

func demo2() {
	score, valid := 3, true

	if score > 3 && valid {
		fmt.Println("good")
	} else {
		fmt.Println("low")
	}
}

func demo3() {
	score := 3

	if score > 3 {
		fmt.Println("good")
	} else if score == 3 {
		fmt.Println("on the edge")
	} else {
		fmt.Println("low")
	}
}

func demo4(osArgs []string) {
	if len(osArgs) < 2 {
		fmt.Println(strings.TrimSpace(usage))

		// ALTERNATIVE:
		// fmt.Println("Please tell me a value in feet")

		return
	}

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * 0.3048

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
}

// func demo5() {
// }
