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
	"strings"
)

// NOTE: You should always pass it at least one argument
// go run main.go hello! 携带一个参数运行
func main() {
	// os.Args[0]是程序本身的名字，所以如果长度小于2，说明没有提供命令行参数。
	if len(os.Args) < 2 {
		fmt.Println("Please provide a message eg: hello!")
		// 在提示错误后，使用 return 终止程序的正常执行。
		return
	}
	msg := os.Args[1]

	l := len(msg)

	s := msg + strings.Repeat("!", l)
	s = strings.ToUpper(s)

	fmt.Println(s)
}
