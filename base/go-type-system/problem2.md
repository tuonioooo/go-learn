# Go 语言自定义类型常见问题

## 1. 为什么要定义新类型?

**答案：以上所有选项都正确**

定义新类型的原因包括：
1. 声明新方法
2. 类型安全
3. 提高可读性和表达意义

**说明：**
- 选项 1-3 单独来看都只是其中一个原因
- 实际上，定义新类型可以同时获得以上所有好处

---

## 2. 自定义类型不会从其底层类型继承哪个属性?

假设你声明了以下自定义类型：

```go
type Millennium time.Duration
```

**答案：方法（Methods）**

**说明：**
- ✅ **方法**：自定义类型不会继承源类型的方法
- ❌ **表示方式**：会从底层类型继承
- ❌ **大小**：会从底层类型继承
- ❌ **取值范围**：会从底层类型继承

---

## 3. 如何使用 float32 定义新类型?

**答案：`type radius float32`**

**说明：**
- ❌ `var radius float32` - 这是声明变量,不是定义类型
- ❌ `radius = type float32` - 语法错误
- ✅ `type radius float32` - 正确!`radius` 是基于 `float32` 定义的新类型
- ❌ `type radius = float32` - 这声明了 `radius` 作为 `float32` 的别名,它们是相同的类型

---

## 4. 如何修复以下代码?

```go
type Distance int
var (
    village Distance = 50
    city = 100
)
fmt.Print(village + city)
```

**答案：`village + Distance(city)`**

**说明：**
- ❌ `int(village + city)` - 无法修复类型不匹配
- ❌ `village + int(city)` - 无法修复类型不匹配
- ❌ `village(int) + city` - 语法错误
- ✅ `village + Distance(city)` - 正确!现在 `city` 在表达式中的类型是 `Distance`

**问题分析：** `village` 是 `Distance` 类型,而 `city` 是 `int` 类型,不能直接相加

---

## 5. 对于以下程序,你应该声明哪些类型?

```go
package main
import "fmt"
func main() {
	celsius := 35.
	fahrenheit := (9*celsius + 160) / 5
	fmt.Printf("%g ºC is %g ºF\n", celsius, fahrenheit)
}
```

**答案：使用 float64 定义 Celsius 和 Fahrenheit**

**说明：**
- ❌ **使用 int64 定义 Celsius 和 Fahrenheit** - 温度值有小数部分,使用整数不是最佳选择
- ✅ **使用 float64 定义 Celsius 和 Fahrenheit** - float64 可以表示温度值
- ❌ **使用 int64 定义 Temperature** - 温度值有小数部分,且有两种不同的温度单位(摄氏度和华氏度),最好创建两个不同的类型
- ❌ **使用 uint32 定义 Temperature** - 同上

---

## 6. Millennium 类型的底层类型是什么?

```go
type (
    Duration int64
    Century Duration
    Millennium Century
)
```

**答案：`int64`**

**说明：**
- ✅ **int64** - 正确!Go 的类型系统是扁平的。自定义类型的底层类型是具有真实结构的类型。int64 不仅仅是一个名称,它有自己的结构,是预声明类型
- ❌ **Duration** - Duration 只是一个新类型名称,没有自己的结构
- ❌ **Century** - Century 只是一个新类型名称,没有自己的结构

---

## 7. 哪些类型之间不需要相互转换?

**提示：** 别名类型不需要类型转换

**答案：byte 和 uint8**

**说明：**
- ✅ **byte 和 uint8** - byte 是 uint8 的别名,它们是相同的类型,不需要转换
- ❌ **byte 和 rune** - 不同类型,需要转换
- ❌ **rune 和 uint8** - 不同类型,需要转换
- ❌ **byte 和 uint32** - 不同类型,需要转换

---

## 8. 哪些类型之间不需要相互转换?

**提示：** 别名类型不需要类型转换

**答案：rune 和 int32**

**说明：**
- ❌ **byte 和 uint32** - 不同类型,需要转换
- ❌ **byte 和 rune** - 不同类型,需要转换
- ✅ **rune 和 int32** - rune 是 int32 的别名,它们是相同的类型,不需要转换
- ❌ **byte 和 int8** - 不同类型,需要转换

---

## 关键概念总结

### 自定义类型 vs 类型别名

```go
// 自定义类型 - 创建新类型
type Celsius float64

// 类型别名 - 创建别名
type Fahrenheit = float64
```

### Go 的预声明类型别名

- `byte` 是 `uint8` 的别名
- `rune` 是 `int32` 的别名

这些别名类型之间可以直接使用,无需类型转换!