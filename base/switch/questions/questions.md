# Go 语言 Switch 语句测验题

## if 语句和 switch 语句之间有什么区别？
1. if 语句控制执行流程，但 switch 语句不控制
2. if 语句是比 switch 语句更具可读性的替代方案
3. switch 语句是比 if 语句更具可读性的替代方案 **✓ 正确答案**

> **解析 1：** 两者都控制执行流程。
> 
> **解析 2：** 有时是对的，但对于复杂的 if 语句，switch 语句可以使它们更具可读性。

---

## 你可以使用什么类型的值作为 switch 条件？
```go
switch condition {
    // ....
}
```
1. int
2. bool
3. string
4. 任何类型的值 **✓ 正确答案**

> **解析 4：** 与大多数其他基于 C 的语言不同，在 Go 中，switch 语句实际上是 if 语句的语法糖。这意味着 Go 在幕后将 switch 语句转换为 if 语句。因此，任何类型的值都可以用作条件。

---

## 对于以下 switch 语句，你可以使用什么类型的值作为 case 条件？
```go
switch false {
case condition:
    // ...
}
```
1. int
2. bool **✓ 正确答案**
3. string
4. 任何类型的值

> **解析 2：** 是的，你只能使用 bool 值，因为在示例中，switch 的条件是 bool 类型。

---

## 对于以下 switch 语句，你可以使用什么类型的值作为 case 条件？
```go
switch "go" {
case condition:
    // ...
}
```
1. int
2. bool
3. string **✓ 正确答案**
4. 任何类型的值

> **解析 3：** 是的，你只能使用 string 值，因为在示例中，switch 的条件是字符串。

---

## 以下代码有什么问题？
```go
switch 5 {
case 5:
case 6:
case 5:
}
```
1. case 子句没有任何语句
2. 相同的常量在 case 条件中被多次使用 **✓ 正确答案**
3. switch 条件不能是无类型的 int

> **解析 2：** 没错。5 被多个 case 子句用作 case 条件。它应该只使用一次。

---

## 当以下代码运行时，哪个 case 将被执行？
```go
weather := "hot"
switch weather {
case "very cold", "cold":
    // ...
case "very hot", "hot":
    // ...
default:
    // ...
}
```
1. 都不执行
2. 第一个 case 子句
3. 第二个 case 子句 **✓ 正确答案**
4. default 子句

> **解析 3：** 没错。当 switch 的条件是 "hot" 或 "very hot" 时，第二个 case 将被执行。

---

## 当以下代码运行时，哪个子句将被执行？
```go
switch weather := "too hot"; weather {
default:
    // ...
case "very cold", "cold":
    // ...
case "very hot", "hot":
    // ...
}
```
1. 都不执行
2. 第一个 case 子句
3. 第二个 case 子句
4. default 子句 **✓ 正确答案**

> **解析 4：** 没错。switch 的条件与任何 case 条件都不匹配，因此 default 子句将作为最后的选择被执行。而且 default 子句可以在 switch 语句内的任何位置。

---

## 当以下代码运行时，哪个 case 将被执行？
```go
switch weather := "too hot"; weather {
case "very cold", "cold":
    // ...
case "very hot", "hot":
    // ...
}
```
1. 都不执行 **✓ 正确答案**
2. 第一个 case 子句
3. 第二个 case 子句
4. default 子句

> **解析 1：** 没错。switch 的条件与任何 case 条件都不匹配，并且没有 default 子句。因此，所有子句都不会被执行。

---

## 以下程序会打印什么？
```go
richter := 2.5

switch r := richter; {
case r < 2:
    fallthrough
case r >= 2 && r < 3:
    fallthrough
case r >= 5 && r < 6:
    fmt.Println("Not important")
case r >= 6 && r < 7:
    fallthrough
case r >= 7:
    fmt.Println("Destructive")
}
```
1. 什么都不打印
2. "Not important" **✓ 正确答案**
3. "Destructive"

> **解析 2：** 没错！当使用 fallthrough 时，后续子句将被执行而不检查其条件。

---

## 以下代码有什么问题？
```go
richter := 14.5

switch r := richter; {
case r >= 5 && r < 6:
    fallthrough
    fmt.Println("Moderate")
default:
    fmt.Println("Unknown")
case r >= 7:
    fmt.Println("Destructive")
}
```
1. default 子句应该是最后一个子句
2. default 子句应该有一个 fallthrough 语句
3. 第一个 case 应该使用 fallthrough 语句作为其最后一条语句 **✓ 正确答案**

> **解析 3：** 没错！在 case 代码块中，fallthrough 只能用作最后一条语句。

---

## 以下程序会打印什么？
```go
n := 8

switch {
case n > 5, n < 1:
    fmt.Println("n is big")
case n == 8:
    fmt.Println("n is 8")
}
```
1. 什么都不打印
2. "n is 8"
3. "n is big" **✓ 正确答案**

> **解析 3：** 没错！switch 从上到下运行，case 条件从左到右运行。这里，第一个 case 的第一个条件表达式（即 n > 5）将产生 true，因此第一个 case 将被执行。