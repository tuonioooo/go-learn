package main

import (
	"fmt"
	f "fmt" // 重命名导入
)

func rename_func() {
	fmt.Println("Hello rename_func!")
	f.Println("Hello f!")
}
