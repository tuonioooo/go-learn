package main

import (
	"testing"
)

func ProcessArray(arr [5]int) {
	// 这个函数只接收长度为5的数组
}

func TestArrayLength(t *testing.T) {
	// ✅ 正确
	arr1 := [5]int{1, 2, 3, 4, 5}
	ProcessArray(arr1)

	// ❌ 错误 - 编译器提示类型不匹配
	// arr2 := [6]int{1, 2, 3, 4, 5, 6}
	// ProcessArray(arr2) // 编译错误：cannot use arr2 (variable of type [6]int) as [5]int
}
