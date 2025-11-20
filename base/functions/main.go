package main

type Person struct {
	Name string
	Age  int
}

func Greet(p Person) string {
	return "Hello, " + p.Name
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func main() {
	p := Person{Name: "Alice", Age: 20}

	// Greet 是普通函数（没有接收器），所以直接调用函数名
	greeting := Greet(p)

	// IsAdult 是方法（有接收器），所以通过实例调用
	isAdult := p.IsAdult()

	println(greeting) // 输出: Hello, Alice
	println(isAdult)  // 输出: true
}
