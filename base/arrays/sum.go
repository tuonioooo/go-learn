package main

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAllTails 计算多个切片中除第一个元素外其他元素的总和
// 参数 numbersToSum ...[]int 表示可以传入任意数量的整数切片
// 返回一个切片，包含每个输入切片的"尾部和"（除第一个元素外的所有元素之和）
func SumAllTails(numbersToSum ...[]int) []int {
	// 创建一个空切片用于存储每个输入切片的尾部和
	var sums []int

	// 遍历每个输入的切片
	for _, numbers := range numbersToSum {
		// 如果切片为空，将0添加到结果中
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// numbers[1:] 创建一个新的切片，包含从索引1到末尾的所有元素（即去掉第一个元素）
			tail := numbers[1:]
			// 计算尾部切片的总和并添加到结果中
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
