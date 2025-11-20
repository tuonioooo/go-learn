package main

import "fmt"

// 我在讲座中没有讨论过这个
// 作为附注，我想把它放在这里
// 请审查一下

var declareMeAgain = 10

func nested() { // 块级作用域开始

	// 声明同名变量
	// 它们可以同时存在
	// 这个变量只属于这个作用域
	// 包级别的变量仍然保持不变
	var declareMeAgain = 5
	fmt.Println("inside nested:", declareMeAgain)

} // 块级作用域结束
