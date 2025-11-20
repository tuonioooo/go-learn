# Go fmt 包完全讲解

`fmt` 包提供了格式化 I/O（输入/输出）的功能，是 Go 程序中最常用的包之一。

---

## 1. fmt 包概览

### 1.1 什么是 fmt 包？

`fmt` 包用于：
- ✅ 格式化输出（打印到控制台）
- ✅ 格式化输入（从用户获取输入）
- ✅ 格式化字符串（生成格式化文本）
- ✅ 字符串扫描（解析字符串数据）

### 1.2 导入 fmt 包

```go
import "fmt"
```

---

## 第一部分：输出函数

## 2. Print 系列函数

### 2.1 fmt.Print() - 输出数据（无换行）

**语法**：
```go
func Print(a ...interface{}) (n int, err error)
```

**特点**：
- 输出值之间**没有分隔符**
- 末尾**没有换行符**
- 返回写入的字节数和错误

**示例**：
```go
package main

import "fmt"

func main() {
    fmt.Print("Hello")
    fmt.Print("World")
    // 输出: HelloWorld（没有空格和换行）
}
```

**应用场景**：
```go
// 逐字输出
fmt.Print("a")
fmt.Print("b")
fmt.Print("c")
// 输出: abc
```

---

### 2.2 fmt.Println() - 输出数据（带换行）

**语法**：
```go
func Println(a ...interface{}) (n int, err error)
```

**特点**：
- 输出值之间用**空格分隔**
- 末尾**自动添加换行符**（`\n`）
- 返回写入的字节数和错误

**示例**：
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello", "World")
    fmt.Println("Go", "is", "awesome")
}
```

**输出**：
```
Hello World
Go is awesome
```

**最常见的用法**：
```go
fmt.Println("Hello, World!")
fmt.Println(42)
fmt.Println(3.14)
fmt.Println(true)
```

---

### 2.3 fmt.Printf() - 格式化输出

**语法**：
```go
func Printf(format string, a ...interface{}) (n int, err error)
```

**特点**：
- 按照**格式字符串**输出
- **不自动添加换行符**
- 需要显式添加 `\n` 换行

**格式动词（Format Verbs）**：

| 动词 | 说明 | 示例 |
|------|------|------|
| `%v` | 默认格式 | `fmt.Printf("%v\n", 42)` → `42` |
| `%T` | 打印类型 | `fmt.Printf("%T\n", 42)` → `int` |
| `%t` | 打印布尔类型 | `fmt.Printf("%t\n", 42>1)` → `true` |
| `%d` | 十进制整数 | `fmt.Printf("%d\n", 42)` → `42` |
| `%b` | 二进制 | `fmt.Printf("%b\n", 42)` → `101010` |
| `%o` | 八进制 | `fmt.Printf("%o\n", 42)` → `52` |
| `%x` | 十六进制（小写） | `fmt.Printf("%x\n", 255)` → `ff` |
| `%X` | 十六进制（大写） | `fmt.Printf("%X\n", 255)` → `FF` |
| `%f` | 浮点数 | `fmt.Printf("%f\n", 3.14)` → `3.140000` |
| `%.2f` | 浮点数（2位小数） | `fmt.Printf("%.2f\n", 3.14159)` → `3.14` |
| `%e` | 科学计数法 | `fmt.Printf("%e\n", 3.14)` → `3.140000e+00` |
| `%g` | 浮点数（自动选择） | `fmt.Printf("%g\n", 3.14)` → `3.14` |
| `%s` | 字符串 | `fmt.Printf("%s\n", "hello")` → `hello` |
| `%c` | 单个字符 | `fmt.Printf("%c\n", 65)` → `A` |
| `%q` | 带引号的字符串 | `fmt.Printf("%q\n", "hello")` → `"hello"` |
| `%%` | 百分号 | `fmt.Printf("100%%\n")` → `100%` |

**常见示例**：

```go
package main

import "fmt"

func main() {
    // 整数格式化
    fmt.Printf("十进制: %d\n", 42)           // 输出: 十进制: 42
    fmt.Printf("二进制: %b\n", 42)           // 输出: 二进制: 101010
    fmt.Printf("十六进制: %x\n", 255)        // 输出: 十六进制: ff
    
    // 浮点数格式化
    fmt.Printf("浮点数: %f\n", 3.14159)     // 输出: 浮点数: 3.141590
    fmt.Printf("2位小数: %.2f\n", 3.14159) // 输出: 2位小数: 3.14
    
    // 字符串格式化
    fmt.Printf("字符串: %s\n", "Hello")    // 输出: 字符串: Hello
    
    // 类型信息
    fmt.Printf("值: %v, 类型: %T\n", 42, 42)  // 输出: 值: 42, 类型: int
}
```

---

### 2.4 输出函数对比

| 函数 | 分隔符 | 换行符 | 格式化 | 常见用途 |
|------|--------|--------|--------|----------|
| `Print()` | ❌ | ❌ | ❌ | 简单输出 |
| `Println()` | ✅ | ✅ | ❌ | 最常用 |
| `Printf()` | ❌ | ❌ | ✅ | 格式化输出 |

**选择指南**：
```go
// 简单输出多个值 → 用 Println
fmt.Println("Name:", "Alice", "Age:", 25)

// 需要格式化 → 用 Printf
fmt.Printf("Price: $%.2f\n", 19.99)

// 逐个输出，无空格 → 用 Print
fmt.Print("Loading"), fmt.Print("..."), fmt.Print("Done\n")
```

---

## 3. Sprintf 系列 - 格式化字符串

### 3.1 fmt.Sprintf() - 返回格式化字符串

**语法**：
```go
func Sprintf(format string, a ...interface{}) string
```

**特点**：
- **返回字符串**，不输出到控制台
- 用于生成格式化的字符串
- 支持所有的格式动词

**示例**：
```go
package main

import "fmt"

func main() {
    // 生成格式化字符串
    message := fmt.Sprintf("Price: $%.2f", 19.99)
    fmt.Println(message)  // 输出: Price: $19.99
    
    // 用于拼接字符串
    name := "Alice"
    age := 25
    info := fmt.Sprintf("Name: %s, Age: %d", name, age)
    fmt.Println(info)  // 输出: Name: Alice, Age: 25
    
    // 生成日志消息
    status := "success"
    code := 200
    log := fmt.Sprintf("[%s] HTTP %d", status, code)
    fmt.Println(log)  // 输出: [success] HTTP 200
}
```

**常见应用**：
```go
// 1. 动态生成错误信息
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        errMsg := fmt.Sprintf("cannot divide %f by zero", a)
        return 0, fmt.Errorf(errMsg)
    }
    return a / b, nil
}

// 2. 拼接 URL
func MakeURL(domain, path string, id int) string {
    return fmt.Sprintf("https://%s/%s/%d", domain, path, id)
}

// 3. 生成日志
func LogAction(user string, action string, timestamp string) string {
    return fmt.Sprintf("[%s] User %s performed: %s", timestamp, user, action)
}
```

---

## 第二部分：输入函数

## 4. Scan 系列函数 - 读取输入

### 4.1 fmt.Scan() - 扫描空格分隔的值

**语法**：
```go
func Scan(a ...interface{}) (n int, err error)
```

**特点**：
- 从标准输入读取数据
- **按空格分隔**多个值
- 需要传入**指针**
- 返回读取的参数个数和错误

**示例**：
```go
package main

import "fmt"

func main() {
    var name string
    var age int
    
    fmt.Print("Enter name and age: ")
    fmt.Scan(&name, &age)  // 需要传入指针
    
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

**使用场景**：
```go
// 输入示例: Alice 25
// 输入后，name="Alice", age=25
```

---

### 4.2 fmt.Scanln() - 扫描到换行符

**语法**：
```go
func Scanln(a ...interface{}) (n int, err error)
```

**特点**：
- 从标准输入读取数据
- 读到**换行符**停止
- 需要传入**指针**
- 自动忽略行末的换行符

**示例**：
```go
package main

import "fmt"

func main() {
    var name string
    var age int
    
    fmt.Print("Enter name: ")
    fmt.Scanln(&name)
    
    fmt.Print("Enter age: ")
    fmt.Scanln(&age)
    
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

**与 Scan 的区别**：

| 函数 | 读取方式 | 示例 |
|------|----------|------|
| `Scan()` | 按空格分隔 | `fmt.Scan(&a, &b)` 输入: `5 10` |
| `Scanln()` | 按换行分隔 | `fmt.Scanln(&a)` 输入: `Alice` (回车) |

---

### 4.3 fmt.Scanf() - 格式化扫描

**语法**：
```go
func Scanf(format string, a ...interface{}) (n int, err error)
```

**特点**：
- 按照**格式字符串**读取输入
- 支持所有的格式动词
- 需要传入**指针**

**示例**：
```go
package main

import "fmt"

func main() {
    var name string
    var age int
    var height float64
    
    fmt.Print("Enter (name age height): ")
    fmt.Scanf("%s %d %f", &name, &age, &height)
    
    fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
}
```

**使用格式化动词**：
```go
// 输入示例: Alice 25 1.70
// name="Alice", age=25, height=1.70

// 带格式的输入
fmt.Print("Enter price ($): ")
var price float64
fmt.Scanf("$%.2f", &price)  // 输入: $19.99
```

---

## 5. Scan 系列完整对比

| 函数 | 分隔符 | 格式化 | 适用场景 |
|------|--------|--------|----------|
| `Scan()` | 空格 | ❌ | 多个值输入 |
| `Scanln()` | 换行 | ❌ | 单行输入 |
| `Scanf()` | 自定义 | ✅ | 格式化输入 |

---

## 6. ScanString 和其他扫描函数

### 6.1 fmt.Sscan() - 从字符串扫描

**语法**：
```go
func Sscan(str string, a ...interface{}) (n int, err error)
```

**特点**：
- 从**字符串**而不是标准输入读取
- 按空格分隔

**示例**：
```go
package main

import "fmt"

func main() {
    input := "Alice 25 3.14"
    
    var name string
    var age int
    var height float64
    
    fmt.Sscan(input, &name, &age, &height)
    
    fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
    // 输出: Name: Alice, Age: 25, Height: 3.14
}
```

---

### 6.2 fmt.Sscanf() - 格式化扫描字符串

**语法**：
```go
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
```

**示例**：
```go
package main

import "fmt"

func main() {
    input := "2025-11-19"
    
    var year, month, day int
    
    fmt.Sscanf(input, "%d-%d-%d", &year, &month, &day)
    
    fmt.Printf("Year: %d, Month: %d, Day: %d\n", year, month, day)
    // 输出: Year: 2025, Month: 11, Day: 19
}
```

---

## 第三部分：高级用法

## 7. 宽度和精度控制

### 7.1 宽度（Width）

用于指定最小字段宽度：

```go
package main

import "fmt"

func main() {
    // 右对齐（默认）
    fmt.Printf("|%5d|\n", 42)      // 输出: |   42|
    fmt.Printf("|%5s|\n", "hi")    // 输出: |   hi|
    
    // 左对齐（使用 -）
    fmt.Printf("|%-5d|\n", 42)     // 输出: |42   |
    fmt.Printf("|%-5s|\n", "hi")   // 输出: |hi   |
    
    // 用 0 填充（仅整数）
    fmt.Printf("|%05d|\n", 42)     // 输出: |00042|
}
```

**输出**：
```
|   42|
|   hi|
|42   |
|hi   |
|00042|
```

---

### 7.2 精度（Precision）

用于指定小数位数或字符串长度：

```go
package main

import "fmt"

func main() {
    // 浮点数精度
    fmt.Printf("%.2f\n", 3.14159)   // 输出: 3.14
    fmt.Printf("%.4f\n", 3.14159)   // 输出: 3.1416
    
    // 字符串长度（最多字符数）
    fmt.Printf("%.3s\n", "hello")   // 输出: hel
    fmt.Printf("%.5s\n", "hi")      // 输出: hi
}
```

---

### 7.3 组合宽度和精度

```go
package main

import "fmt"

func main() {
    // 宽度 + 精度
    fmt.Printf("|%8.2f|\n", 3.14159)  // 输出: |    3.14|
    fmt.Printf("|%-8.2f|\n", 3.14159) // 输出: |3.14    |
    
    // 动态宽度和精度
    width := 10
    precision := 2
    fmt.Printf("|%*.*f|\n", width, precision, 3.14159)
    // 输出: |      3.14|
}
```

---

## 8. 特殊格式化

### 8.1 指针和数据结构

```go
package main

import "fmt"

func main() {
    // 指针
    x := 42
    fmt.Printf("%p\n", &x)  // 输出: 0xc0000a0008（内存地址）
    
    // 结构体（默认格式）
    type Person struct {
        Name string
        Age  int
    }
    
    p := Person{"Alice", 25}
    fmt.Printf("%v\n", p)     // 输出: {Alice 25}
    fmt.Printf("%+v\n", p)    // 输出: {Name:Alice Age:25}
    fmt.Printf("%#v\n", p)    // 输出: main.Person{Name:"Alice", Age:25}
}
```

---

### 8.2 布尔值

```go
package main

import "fmt"

func main() {
    fmt.Printf("%v\n", true)   // 输出: true
    fmt.Printf("%t\n", true)   // 输出: true
    fmt.Printf("%t\n", false)  // 输出: false
}
```

---

## 9. 实战示例

### 9.1 简单计算器

```go
package main

import "fmt"

func main() {
    var a, b float64
    var operation string
    
    fmt.Print("Enter first number: ")
    fmt.Scanln(&a)
    
    fmt.Print("Enter operation (+, -, *, /): ")
    fmt.Scanln(&operation)
    
    fmt.Print("Enter second number: ")
    fmt.Scanln(&b)
    
    var result float64
    
    switch operation {
    case "+":
        result = a + b
    case "-":
        result = a - b
    case "*":
        result = a * b
    case "/":
        if b == 0 {
            fmt.Println("Error: cannot divide by zero")
            return
        }
        result = a / b
    default:
        fmt.Println("Error: invalid operation")
        return
    }
    
    fmt.Printf("%.2f %s %.2f = %.2f\n", a, operation, b, result)
}
```

---

### 9.2 格式化输出表格

```go
package main

import "fmt"

func main() {
    type Product struct {
        Name  string
        Price float64
        Stock int
    }
    
    products := []Product{
        {"Apple", 1.50, 100},
        {"Banana", 0.75, 200},
        {"Orange", 2.00, 150},
    }
    
    // 打印表头
    fmt.Printf("%-10s %10s %10s\n", "Name", "Price", "Stock")
    fmt.Println(string(make([]byte, 32)))
    
    // 打印数据
    for _, p := range products {
        fmt.Printf("%-10s $%9.2f %10d\n", p.Name, p.Price, p.Stock)
    }
}
```

**输出**：
```
Name           Price      Stock
                                
Apple           $1.50        100
Banana          $0.75        200
Orange          $2.00        150
```

---

### 9.3 日志格式化

```go
package main

import (
    "fmt"
    "time"
)

func LogInfo(level, message string) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    fmt.Printf("[%s] %s: %s\n", timestamp, level, message)
}

func main() {
    LogInfo("INFO", "Application started")
    LogInfo("WARNING", "Low memory")
    LogInfo("ERROR", "Database connection failed")
}
```

**输出**：
```
[2025-11-19 14:30:45] INFO: Application started
[2025-11-19 14:30:46] WARNING: Low memory
[2025-11-19 14:30:47] ERROR: Database connection failed
```

---

## 10. fmt 函数总结表

### 输出函数

| 函数 | 用途 | 分隔符 | 换行符 | 示例 |
|------|------|--------|--------|------|
| `Print()` | 输出 | ❌ | ❌ | `fmt.Print("Hello")` |
| `Println()` | 输出（换行） | ✅ | ✅ | `fmt.Println("Hello")` |
| `Printf()` | 格式化输出 | - | ❌ | `fmt.Printf("%d\n", 42)` |
| `Sprintf()` | 格式化字符串 | - | ❌ | `s := fmt.Sprintf("%d", 42)` |

### 输入函数

| 函数 | 分隔符 | 格式化 | 示例 |
|------|--------|--------|------|
| `Scan()` | 空格 | ❌ | `fmt.Scan(&x, &y)` |
| `Scanln()` | 换行 | ❌ | `fmt.Scanln(&x)` |
| `Scanf()` | 自定义 | ✅ | `fmt.Scanf("%d", &x)` |
| `Sscan()` | 空格（从字符串） | ❌ | `fmt.Sscan(str, &x)` |
| `Sscanf()` | 自定义（从字符串） | ✅ | `fmt.Sscanf(str, "%d", &x)` |

---

## 11. 常见错误和解决方案

### 11.1 忘记传入指针

❌ **错误**：
```go
var x int
fmt.Scan(x)  // 错误：需要指针
```

✅ **正确**：
```go
var x int
fmt.Scan(&x)  // 正确：传入指针
```

---

### 11.2 格式动词和类型不匹配

❌ **错误**：
```go
age := 25
fmt.Printf("Age: %s\n", age)  // %s 用于字符串，不用于整数
```

✅ **正确**：
```go
age := 25
fmt.Printf("Age: %d\n", age)  // %d 用于整数
```

---

### 11.3 忘记添加换行符

❌ **错误**：
```go
fmt.Printf("Line 1")
fmt.Printf("Line 2")
// 输出会在一行: Line 1Line 2
```

✅ **正确**：
```go
fmt.Printf("Line 1\n")
fmt.Printf("Line 2\n")
// 输出在两行
```

---

## 12. 快速速查表

### 格式动词速查

| 动词 | 类型 | 说明 | 示例 |
|------|------|------|------|
| `%v` | 任意 | 默认格式 | `42` |
| `%T` | 任意 | 类型名 | `int` |
| `%d` | 整数 | 十进制 | `42` |
| `%x` | 整数 | 十六进制 | `2a` |
| `%f` | 浮点 | 浮点数 | `3.14` |
| `%.2f` | 浮点 | 2位小数 | `3.14` |
| `%s` | 字符串 | 字符串 | `hello` |
| `%q` | 字符串 | 带引号 | `"hello"` |
| `%c` | 整数 | 字符 | `A` |
| `%%` | - | 百分号 | `%` |

---

## 总结

掌握 `fmt` 包是 Go 编程的基础。记住这些要点：

1. **输出用 `Println()`** - 最常用，自动分隔和换行
2. **格式化用 `Printf()`** - 需要控制格式时使用
3. **生成字符串用 `Sprintf()`** - 不输出到控制台
4. **读取输入用 `Scan*()`** - 记得传入指针
5. **学会格式动词** - `%d`, `%s`, `%f` 最常用

现在你已经掌握了 `fmt` 包的核心功能，可以开始使用它进行各种输入输出操作了！🎉
