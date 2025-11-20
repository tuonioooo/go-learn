## 什么时候应该使用结构体类型？
1. 用于存储相同类型的值
2. 用于在运行时添加额外类型的值
3. 用于将不同类型组合成一个单一类型以表示一个概念 *正确*

> **1:** 数组、切片和映射更适合这个用途。
>
> **2:** 结构体的字段在编译时是固定的，你不能在运行时添加额外字段，也不能删除它们。
>
> **3:** 正确。结构体类型将一个类型的不同字段组合成一个单一类型。你可以使用结构体类型来表示一个概念。
>

## 结构体字段有哪些属性？
1. 它们都应该是相同的类型
2. 每个字段都应该有一个名称，并且可能具有不同的类型 *正确*
3. 你可以在运行时添加额外的字段
4. 你可以在运行时删除现有的字段

> **2:** 是的，每个字段都应该有一个唯一的名称。同时，每个字段都应该有一个类型，但每个字段可以有不同的类型。
>

## 下面的代码有什么问题？
```go
type weather struct {
    temperature, humidity float64
    windSpeed             float64
    temperature           float64
}
```
1. 没有问题
2. `temperature, humidity float64` 字段存在语法错误
3. `temperature` 字段不唯一 *正确*

> **2:** 这是一种并行定义。它定义了两个 float64 字段：temperature 和 humidity。这是正确的。
>
> **3:** 对！结构体字段名称应该是唯一的。
>

## 以下结构体值的零值是什么？
```go
var movie struct {
    title, genre string
    rating       float64
    released     bool
}
```
1. `{}`
2. `{title: "", genre: "", rating: 0, released: false}` *正确*
3. `{title: "", genre: "", rating: 0, released: true}`
4. `{"title, genre": "", rating: 0, released: false}`

> **1:** 那是一个没有字段的空结构体值。
>
> **2:** 对！Go 会根据字段的类型将其初始化为零值。
>

## 以下结构体的类型是什么？
```go
avengers := struct {
    title, genre string
    rating       float64
    released     bool
}{
    "avengers: end game", "sci-fi", 8.9, true,
}

fmt.Printf("%T\n", avengers)
```
1. `struct{}`
2. `struct{ string; string; float64; bool }`
3. `struct{ title string; genre string; rating float64; released bool }` *正确*
4. `{title: "avengers: end game"; genre: "sci-fi"; rating: 8.9; released: true}`

> **1:** 那是一个没有字段的空结构体类型。
>
> **2:** 字段名称也是结构体类型的一部分。
>
> **3:** 对！字段名称和类型是结构体类型的一部分。
>
> **4:** 不对，那是结构体值。
>

## 以下结构体值相等吗？
```go
type movie struct {
    title, genre string
    rating       float64
    released     bool
}

avengers := movie{"avengers: end game", "sci-fi", 8.9, true}
clone    := movie{
                title: "avengers: end game", genre: "sci-fi",
                rating: 8.9, released: true,
            }
```
1. 有语法错误
2. 相等 *正确*
3. 不相等

> **2:** 创建结构体值时，是否使用字段名称并不重要。所以，它们是相等的。
>

## 以下结构体值相等吗？如果不相等，为什么？
```go
type movie struct {
    title, genre string
    rating       float64
    released     bool
}

avengers := movie{
    title: "avengers: end game", genre: "sci-fi",
    rating: 8.9, released: true,
}

clone := movie{title: "avengers: end game", genre: "sci-fi"}

fmt.Println(avengers == clone)
```
1. 相等：它们有相同的字段集合
2. 不相等：它们不可比较
3. 不相等：字段值不同 *正确*

> **1:** 对，这意味着它们可以比较，但这还不够。
>
> **2:** 不，它们可以比较。它们使用相同的结构体类型。
>
> **3:** 是的，当你省略某些字段时，Go 会为它们分配零值。这里，"clone" 结构体值的 "rating" 和 "released" 字段分别是：0 和 false。
>

## movie 和 performance 结构体类型具有相同的类型吗？
```go
type item        struct { title string }
type movie       struct { item }
type performance struct { item }
```
1. 是：它们有相同的字段集合
2. 否：它们有不同的类型名称 *正确*
3. 否：嵌入字段无法比较

> **2:** 对！具有不同名称的类型不能直接比较。但是，因为它们有相同的字段集合，你可以将其中一个转换为另一个。`movie{} == movie(performance{})` 是可以的，反之亦然。
>

## 这个程序会打印什么？
```go
type item struct{ title string }

type movie struct {
    item
    title string
}

m := movie{
    title: "avengers: end game",
    item:  item{"midnight in paris"},
}

fmt.Println(m.title, "&", m.item.title)
```
1. midnight in paris & midnight in paris
2. avengers: end game & avengers: end game
3. midnight in paris & avengers: end game
4. avengers: end game & midnight in paris *正确*

> **4:** 对！`m.title` 返回 "avengers: end game"，因为外层类型总是优先。然而，`m.item.title` 返回 "midnight in paris"，因为你明确地从内层类型 item 中获取它。
>

## 什么是字段标签？
1. 它允许 Go 更有效地索引结构体字段
2. 你可以用它来为代码写文档
3. 它就像一个注释
4. 关联关于字段的元数据 *正确*

> **4:** 正确。例如，json 包可以根据关联的元数据进行读取和编码/解码。

## 关于字段标签，哪一项是正确的？
1. 它需要根据某些规则进行类型化
2. 你可以在运行时将其更改为不同的值
3. 它只是一个字符串值，本身没有意义 *正确*

> **1:** 这在一定程度上是对的，但它可以有任何值。
>
> **2:** 字段标签是结构体类型定义的一部分，因此你无法在运行时更改它们的值。
>
> **3:** 对！它只是一个字符串值。只有当其他代码读取它时，它才有意义。例如，json 包可以读取它，并根据字段标签的值进行编码/解码。
>

## 下面的程序有什么问题？
```go
type movie struct {
    title string `json:"title"`
}

m := movie{"black panthers"}
encoded, _ := json.Marshal(m)

fmt.Println(string(encoded))
```
1. `movie` 是未导出的，因此无法编码
2. `title` 是未导出的，因此无法编码 *正确*
3. 缺少错误处理，因此无法编码

> **1:** json 包可以编码一个结构体，即使它的类型是未导出的。
>
> **2:** 对！json 包只能编码已导出的字段。
>
> **3:** 最好处理错误，但这不是这里的主要问题。
>

## 为什么需要将指针传递给 Unmarshal 函数？
1. 为了使其工作更快、更高效
2. 以便它可以在内存中更新值 *正确*
3. 为了防止错误

> **2:** 否则，它将无法更新给定的值。这是因为 Go 中的每个值都是按值传递的。所以函数只能更改副本，而不能更改原始值。但是，通过指针，函数可以更改原始值。