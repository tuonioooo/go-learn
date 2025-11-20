package main

import (
	// 文件作用域
	"fmt"
)

func main() {
	fmt.Println("inside main:", declareMeAgain)
	fmt.Println("Hello!")
	// go run . 会正常工作（运行当前包的所有文件）
	bye()
	hey()
	nested()

	rename_func()

	// package-level declareMeAgain isn't effected
	// from the change inside the nested func
	fmt.Println("inside main:", declareMeAgain)

	fmt.Println("包作用域:", ok)

}

// 练习：取消下面函数的注释
//       并分析错误信息
// func bye() {
// 	fmt.Println("Bye!")
// }

// ***** 解释 *****
//
// 错误：bye 函数被重新声明
//       在"这个代码块"中
//
// "这个代码块"的意思 = "main 包"
//
// "重新声明"的意思 = 你在同一个作用域中
//   再次使用了相同的名称
//
// main 包的作用域范围是：
// 所有在同一个 main 包中的源代码文件
