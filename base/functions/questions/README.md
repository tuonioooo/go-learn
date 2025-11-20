## 哪个是正确的函数声明？
1. add func(a, b int) {}
2. function run(a int, b int) {}
3. func run(int a, b) {}
4. func run(a, b int) {} *正确*

## 以下函数的输入参数和结果名称分别是什么？
```go
func run(p Process, id1, id2 int) (pid int, err error) {}
```
1. 输入参数: p, id1, 和 id2。结果: pid, err。 *正确*
2. 输入参数: Process, int, int。结果: int, error。
3. 输入参数: run, p, id1, id2。结果: pid, err。
4. 声明语法错误。

## 什么是 return 语句？
1. 终止一个程序。
2. 终止一个函数，并向调用函数返回零个或多个值。 *正确*
3. 跳过下一条语句并执行再下一条。

## 下面的代码有什么问题？
```go
func add(a, b int) {
    return a + b
    return
}
```
1. return 语句应该只调用一次
2. 最后一条 return 语句是错误的
3. 它应该声明一个 int 类型的结果值，并移除最后一条 return 语句 *正确*
4. 它应该声明任意数值类型的结果值

> **2:** 实际上，这是正确的，因为该函数没有声明结果值。
>
> **3:** 正确。应该是：`func add(a, b int) int { return a + b }`

## 如何修复下面的代码？
```go
func incr(a int) {
    a++
    return
}

num := 10

// 你希望它打印 11，但它却打印 10。
fmt.Println( incr(num) )
```
1. 修改函数：`func incr(a int) int { a++; return a }` *正确*
2. 修改函数：`func incr(a int, newA int) { a++; newA = a }`
3. 修改函数：`func incr(a int) int { return a++ }`

> **1:** Go 是 100% 的按值传递语言。因此，函数的输入参数是该函数的局部变量：修改在函数外部不可见。

## 为什么不应该使用包级变量？
1. 它们没什么问题
2. 函数不能使用包级变量
3. 它们可能会增加代码耦合性并导致代码脆弱 *正确*

> **3:** 原因是：任何人都可以访问和修改它们。

## 为什么应该从以下函数返回一个错误？
```go
// 为什么这样？
func incr(n string) (int, error) {
	m, err := strconv.Atoi(n)
	return m + 1, err
}

// 而不是这样？
func incr(n string) int {
	m, _ := strconv.Atoi(n)
	return m + 1
}
```
1. 你希望让调用者知道何时出了问题 *正确*
2. 当发生错误时，`Atoi` 会返回 0，所以你不需要返回错误

> **2:** 有时，这在一定程度上是对的，但最好还是让调用者知道何时出了问题。

## 下面的 return 语句是如何工作的？为什么能工作？
```go
func spread(samples int, P int) (estimated float64) {
	for i := 0; i < P; i++ {
		estimated += estimate(i, P)
	}
	return
}
```
1. `estimated` 是一个命名的结果值。所以裸返回（naked return）会自动返回 `estimated`。 *正确*
2. return 语句在那里不是必须的。Go 会自动返回 `estimated`。
3. 结果值不能有名称。这段代码是错误的。

## 下面的代码能工作吗？如果能，为什么？
它应该打印：map[1:11 10:3]
```go
func main() {
    stats := map[int]int{1: 10, 10: 2}
    incrAll(stats)
    fmt.Print(stats)
}

func incrAll(stats map[int]int) {
    for k := range stats {
        stats[k]++
    }
}
```
1. 不，它不能工作：Go 是按值传递的语言。`incrAll` 不能更新 map 的值。
2. 是的，它能工作：`incrAll` 可以更新 map 的值。 *正确*

> **2:** Map 值是指针。所以，`incrAll` 可以更新 map 的值。

## 下面的代码能工作吗？如果能，为什么？
它应该打印：[10 5 2]
```go
func main() {
    stats := []int{10, 5}
    add(stats, 2)
    fmt.Print(stats)
}

func add(stats []int, n int) {
    stats = append(stats, n)
}
```
1. 不，它不能工作：`add()` 不能更新原始的切片头（slice header）。 *正确*
2. 是的，它能工作：`add()` 可以向原始切片头添加新元素。

> **1:** Go 是按值传递的编程语言。add() 创建了原始切片头的一个副本，并将新元素添加到新的切片头中，但它从未返回更新后的切片头。因此，它无法更新原始的切片头。它本应该返回原始的切片头。