package main

import "fmt"
import "strings"

func main() {
	fmt.Println("==========demo1=============")
	demo1()
	fmt.Println("==========demo2=============")
	demo2()
	fmt.Println("==========demo3=============")
	demo3()
	fmt.Println("==========demo4=============")
	demo4()
}

func demo1() {
	speed := 100 // #1
	// speed := 10 // #2

	fast := speed >= 80
	slow := speed < 20

	fmt.Printf("fast's type is %T\n", fast)

	fmt.Printf("going fast? %t\n", fast)
	fmt.Printf("going slow? %t\n", slow)

	fmt.Printf("is it 100 mph? %t\n", speed == 100)
	fmt.Printf("is it not 100 mph? %t\n", speed != 100)
}

func demo2() {
	var speed int
	// speed = "100" // NOT OK

	var running bool
	// running = 50 // NOT OK

	var force float64
	// speed = force // NOT OK

	_, _, _ = speed, running, force
}

func demo3() {
	var speed int
	speed = 50 // OK

	var running bool
	running = true // OK

	var force float64
	speed = int(force)

	_, _, _ = speed, running, force
}

func demo4() {
	speed := 100 // #1
	// speed := 10 // #2

	fast := speed >= 80
	slow := speed < 20

	fmt.Printf("going fast? %t\n", fast)
	fmt.Printf("going slow? %t\n", slow)

	fmt.Printf("is it 100 mph? %t\n", speed == 100)
	fmt.Printf("is it not 100 mph? %t\n", speed != 100)

	fmt.Println(strings.Repeat("-", 25))

	// -------------------------
	// #3
	speedB := 150.5
	faster := speedB > 100 // ok: untyped

	fmt.Println("is the other car going faster?", faster)

	// ERROR: Type Mismatch
	// faster = speedB > speed
	faster = speedB > float64(speed)
	fmt.Println("is the other car going faster?", faster)

	// #4
	// ERROR:
	// only numeric values are comparable with
	// ordering operators: <, <=, >, >=

	// fmt.Println(true > faster)
}
