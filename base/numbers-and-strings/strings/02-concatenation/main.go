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
	// 字符串连接使用加号运算符
	//
	// 注意：Go 不支持字符串插值
	age := 30
	name, last := "carl", "sagan"

	// assignment operation using string concat
	name += " edward"

	// equals to this:
	// name = name + " edward"

	fmt.Println(name + " " + last)

	fmt.Println(name+" "+last, age)
}
