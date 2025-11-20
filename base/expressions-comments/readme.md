# Go 表达式和注释完全指南

## 目录
1. [基本概念](#1-基本概念)
2. [表达式详解](#2-表达式详解)
3. [语句详解](#3-语句详解)
4. [运算符](#4-运算符)
5. [注释详解](#5-注释详解)
6. [文档生成](#6-文档生成)
7. [常见问题解答](#7-常见问题解答)

---

## 1. 基本概念

### 1.1 表达式 vs 语句

#### 表达式（Expression）

**定义**：表达式是**产生值的代码片段**。

```go
package main

import "fmt"

func main() {
    fmt.Println(5 + 3)           // 5 + 3 是表达式，产生值 8
    fmt.Println("Hello" + " Go") // 字符串拼接表达式，产生新字符串
    fmt.Println(true && false)   // 布尔表达式，产生 false
    
    x := 10              // 10 是表达式，产生值
    fmt.Println(x * 2)   // x * 2 是表达式，产生值 20
}
```

**特点**：
- ✅ 产生值
- ✅ 可以被赋值给变量
- ✅ 可以作为函数参数
- ✅ 可以组合成更大的表达式

#### 语句（Statement）

**定义**：语句是**指示 Go 做某事的代码片段**，它改变执行流程但不产生值。

```go
package main

import "fmt"

func main() {
    // 这些都是语句
    var x int              // 声明语句
    x = 10                 // 赋值语句
    fmt.Println("Hi")      // 函数调用语句
    
    if x > 5 {             // if 语句（控制流语句）
        fmt.Println("x is large")
    }
    
    for i := 0; i < 3; i++ {  // for 语句（控制流语句）
        fmt.Println(i)
    }
}
```

**特点**：
- ✅ 指示 Go 做某事
- ✅ 改变执行流程
- ❌ 不产生值（不能单独存在）
- ❌ 不能赋值给变量

### 1.2 表达式 vs 语句的关键区别

| 特性 | 表达式 | 语句 |
|------|--------|------|
| **产生值** | ✅ 是 | ❌ 否 |
| **可单独使用** | ❌ 不能 | ✅ 可以 |
| **可作为参数** | ✅ 可以 | ❌ 不能 |
| **可赋值** | ✅ 可以 | ❌ 不能 |
| **改变执行流** | ❌ 否 | ✅ 是 |
| **例子** | `5 + 3`、`x * 2` | `if`、`for`、`var` |

### 1.3 执行流程

**Go 代码执行流程：从上到下，一次执行一条语句**

```go
package main

import "fmt"

func main() {
    fmt.Println("第一行")      // 第一步：执行
    fmt.Println("第二行")      // 第二步：执行
    fmt.Println("第三行")      // 第三步：执行
    
    // 输出：
    // 第一行
    // 第二行
    // 第三行
}
```

**控制语句改变执行流**：

```go
if condition {
    // 条件为 true 时执行
} else {
    // 条件为 false 时执行
}

for i := 0; i < 5; i++ {
    // 循环执行多次
}
```

---

## 2. 表达式详解

### 2.1 字面量表达式（Literal Expressions）

直接写出值的表达式。

```go
5                    // 整数字面量
3.14                 // 浮点数字面量
"Hello"              // 字符串字面量
true                 // 布尔字面量
```

### 2.2 算术表达式（Arithmetic Expressions）

```go
package main

import "fmt"

func main() {
    // 加法
    fmt.Println(10 + 20)       // 输出：30
    fmt.Println("Hello" + " Go") // 输出：Hello Go
    
    // 减法
    fmt.Println(30 - 15)       // 输出：15
    
    // 乘法
    fmt.Println(5 * 6)         // 输出：30
    
    // 除法
    fmt.Println(100 / 10)      // 输出：10
    
    // 取模（求余数）
    fmt.Println(17 % 5)        // 输出：2
    
    // 组合表达式
    fmt.Println((10 + 20) * 2)  // 输出：60
}
```

### 2.3 比较表达式（Comparison Expressions）

产生布尔值（true 或 false）。

```go
package main

import "fmt"

func main() {
    fmt.Println(5 > 3)          // true
    fmt.Println(5 < 3)          // false
    fmt.Println(5 >= 5)         // true
    fmt.Println(5 <= 3)         // false
    fmt.Println(5 == 5)         // true
    fmt.Println(5 != 3)         // true
}
```

### 2.4 逻辑表达式（Logical Expressions）

```go
package main

import "fmt"

func main() {
    // 逻辑与（AND）- 两个都为 true 才是 true
    fmt.Println(true && true)      // true
    fmt.Println(true && false)     // false
    fmt.Println(false && false)    // false
    
    // 逻辑或（OR）- 至少一个为 true 就是 true
    fmt.Println(true || false)     // true
    fmt.Println(false || false)    // false
    
    // 逻辑非（NOT）- 取反
    fmt.Println(!true)             // false
    fmt.Println(!false)            // true
    
    // 组合逻辑表达式
    fmt.Println((5 > 3) && (10 < 20))  // true
    fmt.Println((5 > 3) || (10 > 20))  // true
}
```

### 2.5 变量引用表达式（Variable Reference Expressions）

```go
package main

import "fmt"

func main() {
    x := 10
    y := 20
    
    fmt.Println(x)              // 10
    fmt.Println(x + y)          // 30
    fmt.Println(x * 2 + y)      // 40
}
```

### 2.6 函数调用表达式（Function Call Expressions）

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    // 函数调用是表达式，产生值
    fmt.Println(runtime.NumCPU())        // 返回 CPU 核心数
    fmt.Println(len("Go"))               // 返回字符串长度：2
    fmt.Println(runtime.NumCPU() + 1)    // 组合表达式
}

// 自定义函数
func add(a, b int) int {
    return a + b
}

func example() {
    result := add(5, 3)  // 函数调用表达式，产生 8
    fmt.Println(result)  // 输出：8
}
```

### 2.7 类型转换表达式（Type Conversion Expressions）

```go
package main

import "fmt"

func main() {
    x := 10
    y := float64(x)      // int 转 float64
    
    fmt.Println(y)       // 10
    fmt.Println(y + 3.5) // 13.5
}
```

---

## 3. 语句详解

### 3.1 声明语句（Declaration Statement）

```go
var x int              // 声明变量
var y int = 10         // 声明并初始化
z := 20                // 短声明（只能在函数内）

const PI = 3.14159    // 声明常量

func hello() {}        // 声明函数

type User struct {     // 声明类型
    Name string
}
```

### 3.2 赋值语句（Assignment Statement）

```go
x := 10        // 初始赋值
x = 20         // 重新赋值
x += 5         // 加等于：x = x + 5
x -= 3         // 减等于：x = x - 3
x *= 2         // 乘等于：x = x * 2
x /= 4         // 除等于：x = x / 4
```

### 3.3 控制流语句（Control Flow Statement）

#### if 语句

```go
if condition {
    // 条件为 true 时执行
} else if otherCondition {
    // 其他条件为 true 时执行
} else {
    // 都不满足时执行
}
```

#### for 循环

```go
// 基础 for 循环
for i := 0; i < 5; i++ {
    fmt.Println(i)  // 0, 1, 2, 3, 4
}

// while 风格
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}

// 无限循环
for {
    // 需要 break 才能退出
}
```

#### switch 语句

```go
switch value {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
```

### 3.4 函数调用作为语句

```go
fmt.Println("Hello")    // 函数调用语句（忽略返回值）
x := add(5, 3)          // 函数调用表达式（使用返回值）
```

---

## 4. 运算符

### 4.1 运算符的定义

**运算符** 是**组合表达式的符号**。

```go
// + 是加法运算符，组合两个数字表达式
5 + 3

// && 是逻辑与运算符，组合两个布尔表达式
true && false

// + 是字符串拼接运算符，组合两个字符串表达式
"Hello" + " Go"
```

### 4.2 运算符的重要概念

```go
package main

import "fmt"

func main() {
    // ✅ 运算符可以组合表达式
    fmt.Println(5 + 3)              // 加法表达式
    fmt.Println(true && false)      // 逻辑与表达式
    fmt.Println("Hello" + " Go")    // 字符串拼接表达式
    
    // ✅ 函数调用也是表达式，可以用运算符组合
    fmt.Println(len("Go") + 5)      // len("Go") 返回 2，与 5 相加
    
    // ❌ 运算符不能组合语句
    // if x > 5 + import "fmt"     // 编译错误
}
```

### 4.3 常见运算符

| 类型 | 运算符 | 例子 | 说明 |
|------|--------|------|------|
| **算术** | `+` | `5 + 3` | 加法 |
| | `-` | `10 - 3` | 减法 |
| | `*` | `5 * 6` | 乘法 |
| | `/` | `20 / 4` | 除法 |
| | `%` | `17 % 5` | 取模 |
| **比较** | `>` | `5 > 3` | 大于 |
| | `<` | `5 < 10` | 小于 |
| | `>=` | `5 >= 5` | 大于等于 |
| | `<=` | `5 <= 10` | 小于等于 |
| | `==` | `5 == 5` | 等于 |
| | `!=` | `5 != 3` | 不等于 |
| **逻辑** | `&&` | `true && false` | 逻辑与 |
| | `\|\|` | `true \|\| false` | 逻辑或 |
| | `!` | `!true` | 逻辑非 |
| **赋值** | `=` | `x = 10` | 赋值 |
| | `+=` | `x += 5` | 加等于 |
| | `-=` | `x -= 3` | 减等于 |

---

## 5. 注释详解

### 5.1 注释的定义和作用

**注释** 是**给代码添加解释或生成自动文档的文本**。

**作用**：
- 📝 解释代码的目的和逻辑
- 📚 生成自动化文档
- 👥 帮助其他开发者理解代码

### 5.2 单行注释（Single-line Comments）

使用 `//` 开头，注释该行到行尾的所有内容。

```go
package main

import "fmt"

func main() {
    // 这是一个注释
    fmt.Println("Hello")  // 这也是注释
    
    // x 存储用户年龄
    x := 25
    
    // 检查年龄是否大于 18
    if x > 18 {
        fmt.Println("成人")
    }
}
```

**特点**：
- 从 `//` 开始到行尾
- 不能嵌套
- Go 会忽略注释的所有内容

### 5.3 多行注释（Multi-line Comments）

使用 `/* ... */` 包围的内容。

```go
/*
这是一个多行注释
可以跨越多行
Go 会忽略这里面的所有内容
*/

func main() {
    /*
    这个函数的作用是...
    参数说明...
    返回值说明...
    */
    fmt.Println("Hello")
}
```

**特点**：
- 从 `/*` 开始到 `*/` 结束
- 可以跨越多行
- 不能嵌套（`/* /* */ */` 会出错）

### 5.4 注释规则和最佳实践

#### ✅ 正确用法

```go
// ✅ 正确的单行注释
func Hello() {
    fmt.Println("Hi")
}

// ✅ 正确的多行注释
/*
这是一个文档注释
说明函数的功能
*/
func Goodbye() {
    fmt.Println("Bye")
}

// ✅ 函数内的注释
func main() {
    // 初始化变量
    x := 10
    
    // 输出结果
    fmt.Println(x)
}
```

#### ❌ 错误用法

```go
// ❌ 错误：/ 不是注释开始
/ this is not a comment
func Hello() {}

// ❌ 错误：// 注释会忽略该行剩余代码
func main() {
    // fmt.Println("Hello") fmt.Println("World")
    // 第二个 Println 也被忽略了
}

// ❌ 错误：// 注释必须是完整的，不能中断
func test() {
    fmt.Println("Hi" // 这会导致编译错误
    // 因为字符串没有闭合
}
```

### 5.5 导出声明的文档注释

**为导出的声明添加注释，可以自动生成文档**。

规则：
- 以导出名称开头
- 立即跟在声明前面

```go
package math

// Add 返回两个整数的和
func Add(a, b int) int {
    return a + b
}

// User 表示应用中的用户
type User struct {
    // Name 是用户的全名
    Name string
    // Age 是用户的年龄
    Age int
}

// Greet 向指定的用户问好
func (u User) Greet() string {
    return "Hello, " + u.Name
}
```

运行 `go doc` 会生成：
```
func Add(a, b int) int
    Add 返回两个整数的和

type User struct
    User 表示应用中的用户

func (u User) Greet() string
    Greet 向指定的用户问好
```

---

## 6. 文档生成

### 6.1 使用 go doc 工具

```bash
# 查看当前包的文档
go doc

# 查看指定函数的文档
go doc functionName

# 查看标准库包的文档
go doc fmt
go doc fmt.Println

# 查看源代码和文档
go doc -src packageName
go doc -src fmt.Println

# 启动本地文档服务器
godoc -http=:6060
# 然后在浏览器中访问 http://localhost:6060
```

### 6.2 godoc vs go doc

- **go doc**：新工具（1.13+），简化版，推荐使用
- **godoc**：旧工具，功能更完整，`go doc` 在其基础上构建

### 6.3 生成高质量文档的建议

```go
package pokemon

// Trainer 代表一个宝可梦训练师
type Trainer struct {
    // Name 是训练师的名字
    Name string
    
    // Level 是训练师的等级
    Level int
}

// CatchPokemon 捕捉一个宝可梦
//
// 它接收宝可梦的名称，如果捕捉成功返回 true。
// 根据训练师的等级，捕捉成功的概率不同。
func (t *Trainer) CatchPokemon(name string) bool {
    // 实现...
    return true
}

// ExampleCatchPokemon 是一个使用示例
func ExampleCatchPokemon() {
    trainer := &Trainer{Name: "Ash", Level: 50}
    trainer.CatchPokemon("Pikachu")
    // Output: Caught Pikachu!
}
```

---

## 7. 常见问题解答

### 7.1 什么是语句？

**Q: 以下哪个是对语句的正确描述？**

1. 语句指示 Go 做某事 ✅ **正确**
2. 语句产生一个值
3. 语句不能改变执行流

**答案**：第 1 项

**解释**：
- ✅ 语句指示 Go 执行某项操作
- ❌ 产生值的是表达式，不是语句
- ❌ 语句可以改变执行流（如 if、for、switch）

```go
var x int              // 语句：声明
x = 10                 // 语句：赋值
if x > 5 {             // 语句：改变执行流
    fmt.Println("big")
}
```

---

### 7.2 Go 代码的执行方向

**Q: Go 代码的执行方向是什么？**

1. 从左到右
2. 从上到下 ✅ **正确**
3. 从右到左
4. 从下到上

**答案**：第 2 项

**解释**：
- Go 按照代码从上到下的顺序执行，一次一条语句
- 除非有控制流语句（if、for 等）改变流程

```go
func main() {
    fmt.Println("第一步")   // 第 1 个执行
    fmt.Println("第二步")   // 第 2 个执行
    fmt.Println("第三步")   // 第 3 个执行
}
// 输出：
// 第一步
// 第二步
// 第三步
```

---

### 7.3 什么是表达式？

**Q: 以下哪个是对表达式的正确描述？**

1. 表达式指示 Go 做某事
2. 表达式产生一个值 ✅ **正确**
3. 表达式可以改变执行流

**答案**：第 2 项

**解释**：
- ✅ 表达式产生值
- ❌ 指示 Go 做某事的是语句
- ❌ 改变执行流的是语句（如 if、for）

```go
5 + 3              // 表达式，产生值 8
"Hello" + " Go"    // 表达式，产生新字符串
len("Hi")          // 表达式，产生 2
true && false      // 表达式，产生 false
```

---

### 7.4 什么是运算符？

**Q: 以下哪个是对运算符的正确描述？**

1. 运算符指示 Go 做某事
2. 运算符可以改变执行流
3. 运算符可以组合表达式 ✅ **正确**

**答案**：第 3 项

**解释**：
- ✅ 运算符的作用是组合表达式
- ❌ 指示 Go 做某事是语句的作用
- ❌ 改变执行流也是语句的作用

```go
// + 组合两个数字表达式
5 + 3

// && 组合两个布尔表达式
true && false

// + 组合两个字符串表达式
"Hello" + " Go"

// + 组合一个函数调用表达式和一个数字
len("Hi") + 5
```

---

### 7.5 为什么这个程序不工作？

**Q: 以下程序为什么不工作？**

```go
package main
import "fmt"

func main() {
    "Hello"
}
```

**答案**：`"Hello"` 是一个表达式，不能单独在代码行上而不使用它

**解释**：
- `"Hello"` 是表达式，产生值但没有被使用
- 表达式必须被某个语句使用（如赋值、函数参数等）
- 解决方案：
  ```go
  fmt.Println("Hello")       // 作为函数参数
  msg := "Hello"             // 赋值给变量
  ```

---

### 7.6 程序能否使用分号连接语句？

**Q: 以下程序能工作吗？**

```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.NumCPU()); fmt.Println("cpus"); fmt.Println("the machine")
}
```

选项：
1. 可以：表达式可以用分号分隔
2. 不能：语句应该单独在一行
3. 可以：Go 会自动为每条语句添加分号 ✅ **正确**

**答案**：第 3 项

**解释**：
- Go 会自动在语句末尾添加分号，即使你手动添加也没问题
- 分号是语句分隔符，不是必需的（在同一行）
- 这是 Go 的语法特性

```go
// 都是合法的
fmt.Println("Hello"); fmt.Println("World")
fmt.Println("Hello")
fmt.Println("World")
```

---

### 7.7 为什么这个程序能工作？

**Q: 为什么以下程序能工作？**

```go
package main
import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.NumCPU() + 10)
}
```

**答案**：运算符可以组合表达式 ✅

**解释**：
- `runtime.NumCPU()` 是表达式，产生一个整数值
- `+` 是运算符，组合两个表达式：`runtime.NumCPU()` 和 `10`
- 结果是新的表达式，产生它们的和
- 这个表达式被 `Println` 函数使用

```go
runtime.NumCPU() + 10
     ↓                ↓
  表达式            表达式
     └────────┬────────┘
             运算符 +
              ↓
         新表达式
```

---

### 7.8 注释的作用

**Q: 为什么有时候需要使用注释？**

1. 组合不同的表达式
2. 提供解释或生成自动文档 ✅ **正确**
3. 让代码看起来漂亮

**答案**：第 2 项

**解释**：
- 注释的主要目的是解释代码和生成文档
- 不能用注释来组合表达式
- 美观是附加好处，不是主要目的

```go
// ✅ 注释用于解释
// 检查用户是否成人
if age > 18 {
    fmt.Println("adult")
}

// ✅ 文档注释可生成自动文档
// IsAdult 检查用户是否成人
func IsAdult(age int) bool {
    return age > 18
}
```

---

### 7.9 正确的注释用法

**Q: 以下哪个代码是正确的？**

**选项 1**（❌ 错误）：
```go
package main

/ main function is an entry point /
func main() {
    fmt.Println("Hi")
}
```

**选项 2**（✅ 正确）：
```go
package main

// main function is an entry point /*
func main() {
    fmt.Println(/* this will print Hi! */ "Hi")
}
```

**选项 3**（❌ 错误）：
```go
package main

/*
main function is an entry point

It allows Go to find where to start executing an executable program.
*/
func main() {
    fmt.Println(// "this will print Hi!")
}
```

**答案**：第 2 项

**解释**：

**选项 1 的问题**：
- `/` 不是注释符
- 应该用 `//` 或 `/* ... */`

**选项 2 为什么正确**：
- `//` 开始单行注释
- `/*` 开始多行注释
- Go 忽略所有注释
- 代码语法正确

**选项 3 的问题**：
- `//` 会忽略该行的所有内容
- 包括 `"this will print Hi!")`
- 导致字符串没有闭合，编译错误

---

### 7.10 文档注释命名

**Q: 应该如何命名代码以便 Go 自动生成文档？**

1. 注释每一行代码，然后生成文档
2. 以声明的名称开始注释 ✅ **正确**
3. 使用多行注释

**答案**：第 2 项

**解释**：
- 文档注释必须**立即跟在声明前面**
- **注释第一句必须以导出名称开头**
- Go doc 工具会自动识别并生成文档

```go
// ✅ 正确：以函数名开头
// Add 返回两个整数的和
func Add(a, b int) int {
    return a + b
}

// ✅ 正确：以类型名开头
// User 代表一个用户
type User struct {
    Name string
}

// ❌ 错误：不以名称开头
// This function adds two numbers
func Add(a, b int) int {
    return a + b
}
```

运行 `go doc Add` 会显示：
```
func Add(a, b int) int
    Add 返回两个整数的和
```

---

### 7.11 查看文档的工具

**Q: 使用哪个工具从命令行打印文档？**

1. `go build`
2. `go run`
3. `go doctor`
4. `go doc` ✅ **正确**

**答案**：第 4 项

**用法**：
```bash
# 显示当前包的文档
go doc

# 显示特定函数的文档
go doc functionName

# 显示标准库的文档
go doc fmt
go doc fmt.Println

# 查看源代码
go doc -src functionName
```

---

### 7.12 godoc vs go doc 的区别

**Q: `godoc` 和 `go doc` 的区别是什么？**

1. `go doc` 是 `godoc` 背后的真实工具
2. `godoc` 是 `go doc` 背后的真实工具 ✅ **正确**
3. `go` 工具是 `go doc` 背后的真实工具
4. `go` 工具是 `godoc` 背后的真实工具

**答案**：第 2 项

**解释**：
- **godoc**：原始工具，完整且强大
- **go doc**：Go 1.13+ 推荐使用，是 godoc 的简化版本
- `go doc` 在幕后使用 `godoc` 实现

**对比**：

| 特性 | go doc | godoc |
|------|--------|-------|
| **推荐程度** | ✅ 推荐 | 旧版本 |
| **功能** | 基础 | 完整 |
| **易用性** | ✅ 简单 | 复杂 |
| **命令** | `go doc` | `godoc` |
| **Web 服务** | 需要 godoc | `godoc -http` |

---

## 本包的示例代码说明

### main.go
包含了各种常见的表达式示例：
- 字符串拼接表达式
- 函数调用表达式
- 算术表达式（加减乘除取模）
- 比较表达式
- 逻辑表达式
- 变量引用表达式
- 函数调用表达式

### comments.go
演示了：
- 文件头注释（版权和许可证信息）
- 包级注释
- 多行注释

### if/main.go
展示了：
- if 语句如何改变执行流
- 比较表达式的实际使用

### semicolons/main.go
演示了：
- Go 自动分号插入机制
- 多条语句可以用分号连接在一行

---

## 总结

### 核心概念速记

| 概念 | 定义 | 例子 | 产生值 |
|------|------|------|--------|
| **表达式** | 产生值的代码 | `5 + 3` | ✅ 是 |
| **语句** | 指示 Go 做某事 | `if`、`for` | ❌ 否 |
| **运算符** | 组合表达式的符号 | `+`、`&&` | - |
| **注释** | 解释代码的文本 | `//`、`/* */` | - |

### 执行规则

1. **从上到下** - Go 按顺序执行语句
2. **表达式产生值** - 必须被使用（赋值、参数等）
3. **运算符组合表达式** - 创建新表达式
4. **注释被忽略** - Go 不执行注释中的任何内容

### 最佳实践

- ✅ 使用清晰的表达式名称
- ✅ 为复杂表达式添加注释
- ✅ 为导出的声明添加文档注释
- ✅ 保持注释与代码同步
- ❌ 避免过度注释
- ❌ 避免用运算符组合语句

这就是 Go 表达式和注释的完整指南！🎯
