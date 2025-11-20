package main

type Hello func() string

func (h Hello) Greet() string {
	return h()
}

func main() {
	hello := Hello(func() string {
		return "Hello, World!"
	})
	println(hello.Greet())

}