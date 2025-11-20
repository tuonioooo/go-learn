package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 字符串拼接表达式
	fmt.Println("Hello!" + "!")

	// runtime.NumCPU() 是一个函数调用表达式
	fmt.Println(runtime.NumCPU() + 1)

	comments_func()

	// 算术表达式示例
	fmt.Println(10 + 20)  // 加法
	fmt.Println(30 - 15)  // 减法
	fmt.Println(5 * 6)    // 乘法
	fmt.Println(100 / 10) // 除法
	fmt.Println(17 % 5)   // 取模

	// 布尔表达式
	fmt.Println(5 > 3)         // 比较表达式
	fmt.Println(true && false) // 逻辑与
	fmt.Println(true || false) // 逻辑或

	// 变量赋值表达式
	x := 100
	fmt.Println(x + 50)

	// 函数调用表达式
	fmt.Println(len("Go")) // len() 返回字符串长度
}
