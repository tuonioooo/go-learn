# Go strconv 库常用方法指南

## 简介

`strconv` 是 Go 语言标准库中用于字符串和基本数据类型之间相互转换的包。它提供了高效、类型安全的转换方法。

**官方文档**: [https://pkg.go.dev/strconv](https://pkg.go.dev/strconv)

---

## 常用方法

### 1. 字符串转整数

#### `Atoi` - 字符串转 int
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("转换失败:", err)
        return
    }
    fmt.Printf("结果: %d, 类型: %T\n", num, num)
    // 输出: 结果: 123, 类型: int
}
```

#### `ParseInt` - 字符串转指定类型整数
```go
// ParseInt(s string, base int, bitSize int) (i int64, err error)
// base: 进制 (2-36)，0 表示自动识别
// bitSize: 结果类型的位数 (0, 8, 16, 32, 64)

str := "1010"
num, err := strconv.ParseInt(str, 2, 64)  // 二进制转换
fmt.Println(num)  // 输出: 10

hexStr := "FF"
num, err = strconv.ParseInt(hexStr, 16, 64)  // 十六进制转换
fmt.Println(num)  // 输出: 255
```

#### `ParseUint` - 字符串转无符号整数
```go
str := "123"
num, err := strconv.ParseUint(str, 10, 64)
fmt.Printf("结果: %d, 类型: %T\n", num, num)
// 输出: 结果: 123, 类型: uint64
```

---

### 2. 整数转字符串

#### `Itoa` - int 转字符串
```go
num := 123
str := strconv.Itoa(num)
fmt.Printf("结果: %s, 类型: %T\n", str, str)
// 输出: 结果: 123, 类型: string
```

#### `FormatInt` - 整数转指定进制字符串
```go
// FormatInt(i int64, base int) string
// base: 目标进制 (2-36)

num := int64(255)
binary := strconv.FormatInt(num, 2)    // 转二进制
fmt.Println(binary)  // 输出: 11111111

hex := strconv.FormatInt(num, 16)      // 转十六进制
fmt.Println(hex)     // 输出: ff
```

#### `FormatUint` - 无符号整数转字符串
```go
num := uint64(255)
str := strconv.FormatUint(num, 10)
fmt.Println(str)  // 输出: 255
```

---

### 3. 字符串转浮点数

#### `ParseFloat` - 字符串转浮点数
```go
// ParseFloat(s string, bitSize int) (float64, error)
// bitSize: 32 或 64，表示 float32 或 float64

str := "3.14159"
num, err := strconv.ParseFloat(str, 64)
if err != nil {
    fmt.Println("转换失败:", err)
    return
}
fmt.Printf("结果: %f, 类型: %T\n", num, num)
// 输出: 结果: 3.141590, 类型: float64

// 科学计数法
sciStr := "1.23e-4"
num, _ = strconv.ParseFloat(sciStr, 64)
fmt.Println(num)  // 输出: 0.000123
```

---

### 4. 浮点数转字符串

#### `FormatFloat` - 浮点数转字符串
```go
// FormatFloat(f float64, fmt byte, prec, bitSize int) string
// fmt: 格式化方式 ('f', 'e', 'E', 'g', 'G', 'b')
// prec: 精度 (-1 表示最少位数)
// bitSize: 32 或 64

num := 3.14159265359

// 'f' 格式: 小数形式
str := strconv.FormatFloat(num, 'f', 2, 64)
fmt.Println(str)  // 输出: 3.14

// 'e' 格式: 科学计数法
str = strconv.FormatFloat(num, 'e', 2, 64)
fmt.Println(str)  // 输出: 3.14e+00

// 'g' 格式: 自动选择最紧凑的表示
str = strconv.FormatFloat(num, 'g', 5, 64)
fmt.Println(str)  // 输出: 3.1416
```

---

### 5. 字符串转布尔值

#### `ParseBool` - 字符串转布尔值
```go
// 接受的值: "1", "t", "T", "true", "TRUE", "True" (返回 true)
//          "0", "f", "F", "false", "FALSE", "False" (返回 false)

str1 := "true"
b1, _ := strconv.ParseBool(str1)
fmt.Println(b1)  // 输出: true

str2 := "1"
b2, _ := strconv.ParseBool(str2)
fmt.Println(b2)  // 输出: true

str3 := "false"
b3, _ := strconv.ParseBool(str3)
fmt.Println(b3)  // 输出: false
```

---

### 6. 布尔值转字符串

#### `FormatBool` - 布尔值转字符串
```go
b := true
str := strconv.FormatBool(b)
fmt.Println(str)  // 输出: true

b = false
str = strconv.FormatBool(b)
fmt.Println(str)  // 输出: false
```

---

### 7. 引用和转义

#### `Quote` - 为字符串添加引号并转义
```go
str := `Hello "World"\n`
quoted := strconv.Quote(str)
fmt.Println(quoted)  
// 输出: "Hello \"World\"\\n"
```

#### `Unquote` - 解析引号字符串
```go
quoted := `"Hello \"World\""`
unquoted, err := strconv.Unquote(quoted)
if err == nil {
    fmt.Println(unquoted)  // 输出: Hello "World"
}
```

#### `QuoteToASCII` - 转为 ASCII 可打印字符
```go
str := "Hello 世界"
ascii := strconv.QuoteToASCII(str)
fmt.Println(ascii)  
// 输出: "Hello \u4e16\u754c"
```

---

### 8. 判断字符串是否可转换

#### `CanBackquote` - 判断是否可以用反引号表示
```go
str1 := "Hello World"
fmt.Println(strconv.CanBackquote(str1))  // 输出: true

str2 := "Hello\nWorld"
fmt.Println(strconv.CanBackquote(str2))  // 输出: false
```

---

## 错误处理

所有 Parse 系列函数都会返回 `error`，应当正确处理：

```go
str := "abc"
num, err := strconv.Atoi(str)
if err != nil {
    // 类型断言获取具体错误信息
    if numErr, ok := err.(*strconv.NumError); ok {
        fmt.Printf("转换失败: %s\n", numErr.Err)
        // 输出: 转换失败: invalid syntax
    }
}
```

---

## 性能提示

1. `Atoi` 和 `Itoa` 是最常用的转换方法，性能较好
2. 对于大量转换操作，考虑使用 `AppendInt`、`AppendFloat` 等 Append 系列函数，可以减少内存分配
3. `FormatInt` 和 `FormatUint` 比 `Itoa` 更灵活，但略慢

```go
// 使用 AppendInt 避免额外的内存分配
buf := make([]byte, 0, 64)
buf = strconv.AppendInt(buf, 123, 10)
fmt.Println(string(buf))  // 输出: 123
```

---

## 总结

`strconv` 包提供了完整的字符串和基本类型之间的转换功能：

- **Parse 系列**: 字符串 → 其他类型
- **Format 系列**: 其他类型 → 字符串
- **Append 系列**: 高性能的格式化追加
- **Quote 系列**: 字符串引用和转义

记得始终检查 Parse 系列函数的错误返回值！

**官方文档**: [https://pkg.go.dev/strconv](https://pkg.go.dev/strconv)