package main

import "fmt"

func numbers_demo1() {
	// 当整数和浮点数一起用在表达式中时
	// 结果总是浮点数
	fmt.Println(8 * -4.0) // -32.0 而不是 -32

	// 两个整数相除结果是整数
	fmt.Println(-4 / 2)

	// 取余运算符
	// 只能用于整数
	fmt.Println(5 % 2)
	// fmt.Println(5.0 % 2) // 错误

	// 加法运算符
	fmt.Println(1 + 2.5)
	fmt.Println(2 - 3)

	// 取反运算符
	fmt.Println(-(-2))
	fmt.Println(- -2) // 这也可以
}

func numbers_demo2() {
	var (
		myAge   = 30
		yourAge = 35
		average float64
	)

	average = float64(myAge+yourAge) / 2

	fmt.Println(average)
}

func numbers_demo3() {
	ratio := 1.0 / 10.0

	// 经过 10 次操作后
	// 精度误差就很明显了
	//
	// 顺便说一下，暂时不用关心这个循环语法
	// 我稍后会解释
	for range [...]int{10: 0} {
		ratio += 1.0 / 10.0
	}

	fmt.Printf("%.60f", ratio)
}

func numbers_demo4() {
	// Go 编译器将这些数字视为整数，
	//   因为整数值中没有小数部分，
	// 所以结果是 1 而不是 1.5

	// 因此，这里的 ratio 变量是一个 int 变量，
	//   这是因为 3 除以 2 的结果
	//   是一个整数。

	ratio := 3 / 2

	fmt.Printf("%d", ratio)
}

func numbers_demo5() {
	// 当你在计算中使用浮点数和整数时，
	// 结果总是浮点数。

	ratio := 3.0 / 2

	// 或者:
	// ratio = 3 / 2.0

	// 或者 - 如果 3 在一个 int 变量中:
	// n := 3
	// ratio = float64(n) / 2

	fmt.Printf("%f", ratio)
}

func numbers_demo6() {

	fmt.Println("sum :", 3+2)   // sum - int
	fmt.Println("sum :", 2+3.5) // sum - float64
	fmt.Println("dif :", 3-1)   // difference - int
	fmt.Println("dif :", 3-0.5) // difference - float64
	fmt.Println("prod:", 4*5)   // product - int
	fmt.Println("prod:", 5*2.5) // product - float64
	fmt.Println("quot:", 8/2)   // quotient - int
	fmt.Println("quot:", 8/1.5) // quotient - float64

	// remainder is only for integers
	fmt.Println("rem :", 8%3)
	// fmt.Println("rem:", 8.0%3) // error

	// you can do this
	// since the fractional part of a float is zero
	f := 8.0
	fmt.Println("rem :", int(f)%3)
}

func numbers_demo7() {
	// what's the value of the ratio?
	// 3 / 2 = 1.5?
	var ratio float64 = 3 / 2
	fmt.Println(ratio)

	// explain
	// above expression equals to this:
	ratio = float64(int(3) / int(2))
	fmt.Println(ratio)

	// how to fix it?
	//
	// remember, when one of the values is a float value
	// the result becomes a float
	ratio = float64(3) / 2
	fmt.Println(ratio)

	// or
	ratio = 3.0 / 2
	fmt.Println(ratio)
}

func main() {
	fmt.Println("demo1===")
	numbers_demo1()
	fmt.Println("demo2===")
	numbers_demo2()
	fmt.Println("demo3===")
	numbers_demo3()
	fmt.Println("demo4===")
	numbers_demo4()
	fmt.Println("demo5===")
	numbers_demo5()
	fmt.Println("demo6===")
	numbers_demo6()
	fmt.Println("demo7===")
	numbers_demo7()
}