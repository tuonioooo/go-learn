# Go 自定义类型（Defined Types）完全教程

## 概述

在 Go 中，你可以基于现有类型创建全新的、独立的类型。这被称为**自定义类型**（Defined Types）或**类型定义**（Type Definition）。

---

## 1. 什么是自定义类型？

### 1.1 定义

自定义类型是基于现有类型（基础类型）创建的新类型。新类型拥有独立的名字，但底层结构与基础类型相同。

### 1.2 为什么要创建自定义类型？

```go
// ❌ 不清楚含义
var weight float64 = 75.5  // 这是什么？公斤还是克？

// ✅ 类型清晰
type Kilogram float64
var weight Kilogram = 75.5  // 明确是千克
```

**创建自定义类型的好处**：

1. **类型安全** - 防止不同类型的混合使用
2. **语义清晰** - 代码意图一目了然
3. **添加方法** - 为类型添加独特的行为
4. **编译时检查** - 类型不匹配会导致编译错误

---

## 2. 基本语法

### 2.1 声明方式

```go
// 单个类型定义
type TypeName BaseType

// 示例
type Kilogram float64
type Mile int
```

### 2.2 批量声明

```go
type (
    Kilogram float64
    Pound    float64
    Ounce    float64
)
```

### 2.3 声明位置

**包级别**（推荐）：
```go
package main

// ✅ 包级别，可以在整个包中使用
type Gram float64

func main() {
    var g Gram = 100
}
```

**块级别**（不推荐，但可以）：
```go
func main() {
    // 在函数内声明
    type Millimeter float64
    var m Millimeter = 10.5
}
```

---

## 3. 底层类型（Underlying Type）

### 3.1 什么是底层类型？

底层类型是自定义类型所基于的类型。它决定了数据的实际存储方式和大小。

```go
type Gram float64   // float64 是底层类型
type Mile int       // int 是底层类型
type File string    // string 是底层类型
```

### 3.2 底层类型的作用

底层类型提供：
- ✅ 数据结构和表示
- ✅ 可用的操作（算术、比较等）
- ✅ 大小和内存布局
- ✅ 初始化行为

**示例**：
```go
type Hour int

var h Hour = 24
fmt.Println(h + 1)  // ✅ 可以加法，因为底层类型是 int
// fmt.Println(h + "text")  // ❌ 不能与字符串相加
```

---

## 4. 类型间的转换规则

### 4.1 同底层类型的转换

**如果两个类型有相同的底层类型，它们可以互相转换：**

```go
package main

import "fmt"

type (
    Gram  float64
    Ounce float64
)

func main() {
    var g Gram = 1000
    var o Ounce
    
    // ✅ 可以转换（同底层类型 float64）
    o = Ounce(g)
    
    fmt.Println(g)  // 输出: 1000
    fmt.Println(o)  // 输出: 1000
}
```

**为什么需要显式转换？**

虽然 `Gram` 和 `Ounce` 都基于 `float64`，但它们是不同的类型。Go 要求类型匹配，所以必须显式转换。

### 4.2 不同底层类型的限制

```go
type Kilogram int
type Pound float64

var k Kilogram = 5
var p Pound

// ❌ 错误：不同的底层类型
// p = Pound(k)  // int 和 float64 底层类型不同
```

---

## 5. 实战示例 1：Duration（时间间隔）

### 5.1 使用标准库的 time.Duration

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // time.Duration 是自定义类型
    h, _ := time.ParseDuration("4h30m")
    
    // 1. Duration 有特定的方法
    fmt.Println(h.Hours(), "hours")  // 输出: 4.5 hours
    
    // 2. Duration 不能直接与 int64 混用
    var m int64 = 2
    
    // ❌ 错误：不能直接乘以 int64
    // h = h * m
    
    // ✅ 正确：需要显式转换
    h *= time.Duration(m)
    
    fmt.Println(h)
    // 输出: 9h0m0s
    
    // 3. 类型信息
    fmt.Printf("Type of h: %T\n", h)
    // 输出: Type of h: time.Duration
    
    fmt.Printf("Type of h's underlying type: %T\n", int64(h))
    // 输出: Type of h's underlying type: int64
}
```

**关键点**：
- `time.Duration` 的底层类型是 `int64`
- 但 `time.Duration` 不能直接用 `int64` 相乘
- 必须转换为 `time.Duration` 再进行操作
- `Duration` 有 `Hours()` 等特定方法

---

## 6. 实战示例 2：单位转换

### 6.1 克和盎司的转换

```go
package main

import "fmt"

type (
    Gram  float64
    Ounce float64
)

func main() {
    // 定义克数
    var g Gram = 1000  // 1000 克
    var o Ounce        // 0 盎司（初始值）
    
    // ❌ 错误：不能直接使用不同类型
    // o = g * 0.035274
    
    // ✅ 正确：先转换类型
    o = Ounce(g) * 0.035274
    
    fmt.Printf("%g grams is %.2f ounce\n", g, o)
    // 输出: 1000 grams is 35.27 ounce
}
```

**转换流程**：
```
Gram(1000) → 转换为 Ounce(1000) → 乘以 0.035274 → 结果
```

---

## 7. 实战示例 3：多层类型定义

### 7.1 类型链

```go
package main

import "fmt"

type (
    // Gram 的底层类型是 int
    Gram int
    
    // Kilogram 的底层类型是 Gram（也就是最终是 int）
    Kilogram Gram
    
    // Ton 的底层类型是 Kilogram（也就是最终是 int）
    Ton Kilogram
)

func main() {
    var (
        salt   Gram     = 100
        apples Kilogram = 5
        truck  Ton      = 10
    )
    
    // ❌ 错误：不同名称的类型不能直接赋值
    // salt = apples
    
    // ✅ 正确：需要逐级转换
    salt = Gram(apples)
    apples = Kilogram(truck)
    truck = Ton(Kilogram(Gram(int(apples))))
    
    fmt.Printf("salt: %d, apples: %d, truck: %d\n",
        salt, apples, truck)
    // 输出: salt: 5, apples: 10, truck: 10
}
```

**关键点**：
- 虽然所有类型最终底层都是 `int`，但类型名称不同，所以不能直接混用
- 需要显式转换
- 长链转换可能很复杂

---

## 8. 类型与底层类型的区别

### 8.1 对比表

| 方面 | 自定义类型 | 底层类型 |
|------|-----------|----------|
| 类型名称 | `Gram` | `float64` |
| 类型标识 | 独立的 | 基础类型 |
| 类型检查 | 严格 | - |
| 互相转换 | 需要同底层类型 | - |
| 方法 | 可添加 | 标准方法 |

### 8.2 代码示例

```go
type Kilogram int

var k Kilogram = 50

// 检查类型
fmt.Printf("Type: %T\n", k)          // 输出: main.Kilogram
fmt.Printf("Underlying: %T\n", int(k)) // 输出: int

// 值相同，但类型不同
if k == Kilogram(50) {
    fmt.Println("Values are equal")
}
```

---

## 9. 跨包使用自定义类型

### 9.1 不同包的类型不相互兼容

```go
// 包 A 中
package weightsA
type Gram int

// 包 B 中
package weightsB
import "weightsA"

var g weightsA.Gram = 100
var m weightsB.Gram = 50

// ❌ 错误：不同包的相同名字的类型不兼容
// m = g
```

### 9.2 转换跨包类型

```go
package main

import "github.com/inancgumus/learngo/09-go-type-system/05-defined-types/03-underlying-types/weights"

type Gram int

func main() {
    // 从不同包的类型转换
    var myGram Gram = 100
    var theirGram weights.Gram = 50
    
    // ❌ 错误：不同包的类型
    // myGram = Gram(theirGram)  // 类型不兼容
    
    // ✅ 正确：需要先转换到基础类型
    myGram = Gram(int(theirGram))
    
    // 或者声明新类型关联
    type MyWeightsGram weights.Gram
    var mwg MyWeightsGram = MyWeightsGram(theirGram)
}
```

---

## 10. 自定义类型的实际应用

### 10.1 添加方法

```go
package main

import "fmt"

type Meter float64

// 为 Meter 添加方法
func (m Meter) ToKilometer() float64 {
    return float64(m) / 1000
}

func (m Meter) ToMile() float64 {
    return float64(m) / 1609.34
}

func main() {
    distance := Meter(5000)
    
    fmt.Printf("%.2f meters is %.2f km\n", 
        distance, distance.ToKilometer())
    // 输出: 5000.00 meters is 5.00 km
    
    fmt.Printf("%.2f meters is %.2f miles\n", 
        distance, distance.ToMile())
    // 输出: 5000.00 meters is 3.11 miles
}
```

### 10.2 类型安全

```go
package main

import "fmt"

type UserID int
type ProductID int

type Order struct {
    UserID    UserID
    ProductID ProductID
}

func main() {
    // ✅ 清晰的意图
    order := Order{
        UserID:    UserID(123),
        ProductID: ProductID(456),
    }
    
    // ❌ 如果不小心交换会编译错误
    // order := Order{
    //     UserID:    ProductID(123),      // 错误！
    //     ProductID: UserID(456),         // 错误！
    // }
    
    fmt.Println(order)
}
```

---

## 11. 常见陷阱

### 11.1 忘记类型转换

❌ **错误**：
```go
type Kilogram int
type Pound int

var k Kilogram = 10
var p Pound = k  // 编译错误！
```

✅ **正确**：
```go
var k Kilogram = 10
var p Pound = Pound(int(k))  // 需要转换到基础类型
```

### 11.2 混淆类型和底层类型

❌ **错误的认知**：
```go
type Hour int
var h Hour = 24

// 以为可以这样
fmt.Println(h == 24)  // ✅ 实际上可以！int 会自动转换
```

✅ **更好的做法**：
```go
fmt.Println(h == Hour(24))  // 明确类型
```

### 11.3 链式类型转换的复杂性

```go
type A int
type B A
type C B

var c C = 10

// ❌ 复杂的转换链
var a A = A(int(B(C(c))))

// 建议：直接转到基础类型再转回
var a A = A(int(c))
```

---

## 12. 最佳实践

### ✅ 该做的事

| 做法 | 说明 |
|------|------|
| ✅ 使用有意义的名字 | 名字应该表达类型的含义 |
| ✅ 在包级别声明 | 便于整个包访问 |
| ✅ 添加方法提供功能 | 为类型添加特定的操作 |
| ✅ 为复杂类型提供文档 | 解释为什么需要这个类型 |
| ✅ 尽量避免链式类型 | 保持类型层级简单 |

### ❌ 不该做的事

| 避免 | 原因 |
|------|------|
| ❌ 过度使用自定义类型 | 不是所有情况都需要 |
| ❌ 创建过长的类型链 | 转换会变得复杂 |
| ❌ 在函数内声明类型 | 难以维护和使用 |
| ❌ 忘记显式转换 | 会导致编译错误 |
| ❌ 使用模糊的名字 | 降低代码可读性 |

---

## 13. 快速参考

### 13.1 声明语法

```go
// 单个声明
type TypeName BaseType

// 批量声明
type (
    Type1 BaseType1
    Type2 BaseType2
    Type3 BaseType3
)
```

### 13.2 转换规则

```go
// 同底层类型可转换
type A int
type B int
var a A = 10
var b B = B(a)  // ✅ 可以（都是 int 底层）

// 不同底层类型需要中间转换
type C float64
var c C = C(int(a))  // ✅ int → float64 转换
```

### 13.3 类型检查

```go
// 查看类型
fmt.Printf("%T\n", value)

// 查看底层类型
fmt.Printf("%T\n", BaseType(value))
```

---

## 14. 完整示例

### 14.1 温度转换系统

```go
package main

import "fmt"

type (
    Celsius    float64
    Fahrenheit float64
    Kelvin     float64
)

// 摄氏度转华氏度
func (c Celsius) ToFahrenheit() Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}

// 摄氏度转开尔文
func (c Celsius) ToKelvin() Kelvin {
    return Kelvin(c + 273.15)
}

func main() {
    temp := Celsius(25)
    
    fmt.Printf("%.2f°C = %.2f°F\n", temp, temp.ToFahrenheit())
    fmt.Printf("%.2f°C = %.2fK\n", temp, temp.ToKelvin())
}
```

---

## 总结

### 关键要点

1. **自定义类型是新的类型** - 拥有独立的名字和标识
2. **底层类型提供结构** - 决定了数据的实际表示
3. **类型必须匹配** - 不同类型不能直接混用
4. **需要显式转换** - 同底层类型的转换需要显式进行
5. **为类型添加方法** - 提供类型特定的功能
6. **类型安全很重要** - 自定义类型提供编译时检查

### 快速检查清单

- [ ] 理解了底层类型的概念？
- [ ] 知道如何声明自定义类型？
- [ ] 了解类型转换的规则？
- [ ] 能够为类型添加方法？
- [ ] 理解了跨包类型的不兼容性？

掌握自定义类型让你能写出更类型安全、更易维护的 Go 代码！🎯
