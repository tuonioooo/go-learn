package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello!")

	// 语句改变执行流
	// 特别是控制流语句，比如 if
	if 5 > 1 {
		fmt.Println("bigger")
	}
}
