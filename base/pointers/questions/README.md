## 什么是指针？
1.  一个包含十六进制值的变量
2.  一个包含内存地址的变量
3.  一个可以包含值的内存地址的值 *正确*
4.  一个指向函数的值

> **2:** 虽然指针可以被存入变量，但它不单单是一个变量。你已经很接近了！但这个区别非常重要。
>
> **3:** 指针本身就是一种值，它可以包含另一个值的内存地址。

## 对于下面的结构体，哪一个是计算机指针？
```go
type computer struct {
	brand string
}
```
1. `*computer{}`
2. `var c computer`
3. `var *c computer`
4. `var c *computer` *正确*

> **4:** 在类型前面的 `*` 表示指针类型。

## 对于下面的指针，哪一个能获取它所指向的复合值？
```go
type computer struct {
	brand string
}

c := &computer{"Apple"}
```
1. `*c` *正确*
2. `&c`
3. `c`
4. `*computer`

> **1:** 在指针值前面的 `*` 可以获取该指针所指向的值。
>
> **2:** 在值前面的 `&` 可以获取该值的内存地址。
>
> **4:** 在类型前面的 `*` 表示指针类型。

## 下面代码的输出结果是什么？
```go
type computer struct {
    brand string
}

var a, b *computer
fmt.Print(a == b)

a = &computer{"Apple"}
b = &computer{"Apple"}
fmt.Print(" ", a == b, " ", *a == *b)
```
1. false false false
2. true true true
3. true false true *正确*
4. false true true

> **3:** 开始时 a 和 b 都是 nil，所以它们相等。但是之后，它们从两个复合字面量获得了两个不同的内存地址，所以它们的地址不相等，但它们指向的值是相等的。

## 下面的代码中存在多少个变量？
```go
type computer struct {
    brand string
}

func main() {
    a := &computer{"Apple"} // 修正：添加了缺失的冒号
    b := a
    change(b)
    change(b)
}

func change(c *computer) {
    c.brand = "Indie"
    c = nil
}
```
1. 1
2. 2
3. 3
4. 4 *正确*

> **4:** 每次函数运行时，都会从其输入参数和命名的返回值（如果有）创建新的变量。这里有两个指针变量：a 和 b。然后，因为 change 函数被调用了两次，所以又产生了两个指针变量（每次调用都会创建一个参数 c）。

## 为什么不能获取 map 元素的地址？
1.  它是一个可寻址的值
2.  它是一个不可寻址的值 *正确*
3.  这样做可能会导致程序崩溃

> **2:** Map 的元素是不可寻址的，因此你不能获取它们的地址。