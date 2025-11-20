下面是你提供的 **Arrays Quiz** 完整中文版（保持原有 Markdown 结构、格式、选项、解释，全部为中文）。

---

# 数组测验（Arrays Quiz）

## 这个数组字面量的长度是多少？

```go
gadgets := [...]string{"Mighty Mouse", "Amazing Keyboard", "Shiny Monitor"}
```

1. 0
2. 1
3. 2
4. **3 ✔️**

> **4：**没错！元素列表里有 3 个元素，因此数组长度就是 3。

---

## 这个数组字面量的类型与长度是什么？

```go
gadgets := [...]string{}
```

1. **[0]string 和 0 ✔️**
2. [0]string{} 和 0
3. [1]string 和 1
4. [1]string{} 和 1

> **1：**正确！元素列表为空，因此 Go 将数组长度设为 0。

---

## 下面这个程序会打印什么？

```go
package main
import "fmt"

func main() {
	gadgets := [3]string{"Confused Drone"}
	fmt.Printf("%q\n", gadgets)
}
```

1. [3]string{"Confused Drone", "", ""}
2. [1]string{"Confused Drone"}
3. **["Confused Drone" "" ""] ✔️**
4. ["Confused Drone"]

> **1：**%q 不会打印数组类型。
> **2, 4：**数组长度不能由元素数量自动变化。
> **3：**正确！未初始化的元素会使用零值（空字符串）。

---

## 这两个数组可比较吗？

```go
gadgets := [3]string{"Confused Drone"}
gears   := [...]string{"Confused Drone"}

fmt.Println(gadgets == gears)
```

1. 是的，因为它们的类型和元素都相同
2. **不，因为它们的类型不同 ✔️**
3. 不，因为它们的元素不同

> **2：**gadgets 的类型是 [3]string，而 gears 的类型是 [1]string，它们不是同一种类型。

---

## 这个程序会打印什么？

```go
gadgets := [3]string{"Confused Drone", "Broken Phone"}
gears   := gadgets

gears[2] = "Shiny Mouse"

fmt.Printf("%q\n", gadgets)
```

1. ["Confused Drone" "Broken Phone" "Shiny Mouse"]
2. **["Confused Drone" "Broken Phone" ""] ✔️**
3. ["" "" "Shiny Mouse"]
4. ["" "" ""]

> **2：**对数组赋值会创建副本。gadgets 和 gears 是两个独立数组，互不影响。

---

## digits 数组的类型是什么？

```go
digits := [...][5]string{
	{
		"## ",
		" # ",
		" # ",
		" # ",
		"###",
	},
	[5]string{
		"###",
		"  #",
		"###",
		"  #",
		"###",
	},
}
```

1. [...][5]string
2. [2][2]string
3. **[2][5]string ✔️**
4. [5][5]string

> **3：**外层数组有两个元素，因此长度是 2；元素类型是 `[5]string`。第二个元素前面的 `[5]string` 是可省略的。

---

## 这个程序会打印什么？

```go
rates := [...]float64{
    5: 1.5,
    2.5,
    0: 0.5,
}

fmt.Printf("%#v\n", rates)
```

1. **[7]float64{0.5, 0, 0, 0, 0, 1.5, 2.5} ✔️**
2. [7]float64{1.5, 2.5, 0.5, 0, 0, 0, 0}
3. [3]float64{1.5, 2.5, 0.5}
4. [3]float64{0.5, 2.5, 1.5}

> **1：**正确！结构请参考课程仓库中的示例：
> [https://github.com/inancgumus/learngo/tree/master/14-arrays/11-keyed-elements/06-keyed-and-unkeyed](https://github.com/inancgumus/learngo/tree/master/14-arrays/11-keyed-elements/06-keyed-and-unkeyed)

---

## 这些数组相等吗？

```go
type three [3]int

nums  := [3]int{1, 2, 3}
nums2 := three{1, 2, 3}

fmt.Println(nums == nums2)
```

**注意：要回答这个，需要了解比较规则与未命名类型。**

1. **是的，因为它们有相同的底层类型和元素 ✔️**
2. 不，因为它们的类型不同
3. 不，因为它们长度不同

> **1：**正确！不同命名类型但底层类型相同的数组可以比较。

---

## 这些数组变量相等吗？

```go
type (
    threeA [3]int
    threeB [3]int
)

nums  := threeA{1, 2, 3}
nums2 := threeA(threeB{1, 2, 3})

fmt.Println(nums == nums2)
```

**注意：需要理解比较规则与类型转换。**

1. **是的，因为它们有相同的底层类型和元素 ✔️**
2. 不，因为它们类型不同
3. 不，因为它们长度不同

> **1：**正确！虽然 threeA 与 threeB 是不同的类型，但通过显式转换成同一类型后，它们可以比较。
