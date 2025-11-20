下面是你提供的 **Array Basics Quiz** 内容转换成的 **完整中文 Markdown 文档**，保持原有结构、选项、解释、格式不变，仅翻译内容。

---

# 数组基础测验（Array Basics Quiz）

## 什么是数组？

1. 星球大战里那种加速带电粒子阵列炮
2. 具有动态长度和类型的值集合
3. **具有固定长度和类型的值集合 ✔️**

---

## 下面第二个变量在内存中的地址会是多少？

```go
// 假设第一个变量存储在内存地址 20
var (
    first  int32 = 100
    second int32 = 150
)
```

1. 21
2. 22
3. 24
4. **它可以被存放在任何位置 ✔️**

> **4：**没错。它可以在任何地方。因为与数组不同，普通变量并不能保证会放在连续的内存区域中。

---

## nums 数组的第 3 个元素的内存地址是多少？

```go
// 假设：nums 数组存储在内存地址：500
var nums [5]int64
```

1. 3
2. 2
3. 502
4. 503
5. **516 ✔️**

> **2：**那只是索引而已。
> **3、4：**500+索引？接近了。
> **5：**正确！数组元素在内存中是连续的。数组从 500 开始，每个 int64 为 8 字节，因此：
>
> * 第 1 个：500
> * 第 2 个：508
> * 第 3 个：516
>
> 通用公式：
> **起始地址 + 元素大小 × (元素索引 - 1)**

---

## 这个变量本身存储多少个值？

```go
var gophers [10]string
```

1. 0
2. **1 ✔️**
3. 2
4. 10

> **2：**变量只存一个值，这里 gophers 存的是一个数组值。
> **4：**这是数组长度，不是变量存储的值数量。

---

## 这个数组的长度是多少？

```go
var gophers [5]int
```

1. **5 ✔️**
2. 1
3. 2

---

## 这个数组的长度是多少？

```go
const length = 5 * 2
var gophers [length - 1]int
```

1. 10
2. **9 ✔️**
3. 1

> **2：**5 * 2 - 1 = 9。数组长度可由常量表达式构造。

---

## 这个数组的元素类型是什么？

```go
var luminosity [100]float32
```

1. [100]float32
2. luminosity
3. **float32 ✔️**

---

## 这个数组的类型是什么？

```go
var luminosity [100]float32
```

1. **[100]float32 ✔️**
2. luminosity
3. float32

> **1：**数组类型 = 长度 + 元素类型。

---

## 下面程序会打印什么？

```go
package main
import "fmt"

func main() {
	var names [3]string

	names[len(names)-1] = "!"
	names[1] = "think" + names[2]
	names[0] = "Don't"
	names[0] += " "

	fmt.Println(names[0] + names[1] + names[2])
}
```

1. !think!Don't
2. **Don't think!! ✔️**
3. 程序不正确

---

## 下面程序会打印什么？

> 这是难题，可以用 Playground 试试：
> [https://play.golang.org/p/o0o0UM7Ktyy](https://play.golang.org/p/o0o0UM7Ktyy)

```go
package main
import "fmt"

// 这个循环会设置数组的下一个元素，
// 然后在最后一个元素前停止。
func main() {
	var sum [5]int

	for i, v := range sum {
	    if i == len(sum) - 1 {
	        break
	    }

	    sum[i+1] = 10
	    fmt.Print(v, " ")
	}
}
```

1. **0 0 0 0 ✔️**
2. 10 10 10 10
3. 0 10 10 10

> **1：**正确！`for range` 会复制数组，所以循环内修改原数组不影响迭代变量 v。
> 如果你在循环结束后打印 sum，会看到数组确实被修改了。
