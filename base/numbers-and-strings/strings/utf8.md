# Go unicode/utf8 库使用指南

## 📚 官方文档

- **官方文档**: https://pkg.go.dev/unicode/utf8
- **源码地址**: https://github.com/golang/go/tree/master/src/unicode/utf8

## 📖 简介

`unicode/utf8` 包实现了 UTF-8 编码的编码和解码函数。它提供了对 UTF-8 字节序列的验证、计数和操作功能。

## 🔧 常见用法

### 1. 计算字符串中的 UTF-8 字符数量

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    s := "Hello, 世界"
    
    // 使用 len() 获取字节数
    fmt.Printf("字节数: %d\n", len(s))
    
    // 使用 RuneCountInString() 获取字符数
    fmt.Printf("字符数: %d\n", utf8.RuneCountInString(s))
}

// 输出:
// 字节数: 13
// 字符数: 9
```

### 2. 验证 UTF-8 编码的有效性

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    valid := "Hello, 世界"
    invalid := string([]byte{0xff, 0xfe, 0xfd})
    
    fmt.Printf("%q 是否有效: %v\n", valid, utf8.ValidString(valid))
    fmt.Printf("%q 是否有效: %v\n", invalid, utf8.ValidString(invalid))
    
    // 验证字节切片
    bytes := []byte("测试")
    fmt.Printf("字节切片是否有效: %v\n", utf8.Valid(bytes))
}

// 输出:
// "Hello, 世界" 是否有效: true
// "\xff\xfe\xfd" 是否有效: false
// 字节切片是否有效: true
```

### 3. 解码 UTF-8 字符（Rune）

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    s := "世界"
    
    // DecodeRuneInString 返回第一个字符和其字节长度
    r, size := utf8.DecodeRuneInString(s)
    fmt.Printf("第一个字符: %c, 占用字节: %d\n", r, size)
    
    // 遍历字符串中的所有字符
    for i := 0; i < len(s); {
        r, size := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("字符: %c, Unicode: U+%04X, 字节数: %d\n", r, r, size)
        i += size
    }
}

// 输出:
// 第一个字符: 世, 占用字节: 3
// 字符: 世, Unicode: U+4E16, 字节数: 3
// 字符: 界, Unicode: U+754C, 字节数: 3
```

### 4. 编码 Rune 到 UTF-8

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    r := '界'
    
    // 创建足够大的字节切片
    buf := make([]byte, 4)
    
    // EncodeRune 将 rune 编码到字节切片
    n := utf8.EncodeRune(buf, r)
    fmt.Printf("字符 %c 编码为 %d 个字节: %v\n", r, n, buf[:n])
    
    // 获取 rune 编码所需的字节数
    size := utf8.RuneLen(r)
    fmt.Printf("字符 %c 需要 %d 个字节\n", r, size)
}

// 输出:
// 字符 界 编码为 3 个字节: [231 149 140]
// 字符 界 需要 3 个字节
```

### 5. 获取字符串第一个和最后一个字符

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    s := "Hello, 世界"
    
    // 获取第一个字符
    firstRune, _ := utf8.DecodeRuneInString(s)
    fmt.Printf("第一个字符: %c\n", firstRune)
    
    // 获取最后一个字符
    lastRune, _ := utf8.DecodeLastRuneInString(s)
    fmt.Printf("最后一个字符: %c\n", lastRune)
}

// 输出:
// 第一个字符: H
// 最后一个字符: 界
```

### 6. 检查字节是否是 UTF-8 字符的起始字节

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    s := "世界"
    bytes := []byte(s)
    
    for i, b := range bytes {
        if utf8.RuneStart(b) {
            fmt.Printf("索引 %d 的字节 0x%X 是字符起始字节\n", i, b)
        }
    }
}

// 输出:
// 索引 0 的字节 0xE4 是字符起始字节
// 索引 3 的字节 0xE7 是字符起始字节
```

### 7. 完整示例：字符串处理工具

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func analyzeString(s string) {
    fmt.Printf("\n分析字符串: %q\n", s)
    fmt.Printf("字节长度: %d\n", len(s))
    fmt.Printf("字符数量: %d\n", utf8.RuneCountInString(s))
    fmt.Printf("是否有效的 UTF-8: %v\n", utf8.ValidString(s))
    
    fmt.Println("\n字符详情:")
    for i, r := range s {
        fmt.Printf("  位置 %d: 字符 %c (U+%04X), 字节数: %d\n", 
            i, r, r, utf8.RuneLen(r))
    }
}

func main() {
    analyzeString("Go语言")
    analyzeString("Hello🌍")
}
```

## 📝 重要函数速查表

| 函数 | 功能 |
|------|------|
| `RuneCountInString(s string) int` | 返回字符串中的 UTF-8 字符数量 |
| `RuneCount(p []byte) int` | 返回字节切片中的 UTF-8 字符数量 |
| `ValidString(s string) bool` | 检查字符串是否为有效的 UTF-8 编码 |
| `Valid(p []byte) bool` | 检查字节切片是否为有效的 UTF-8 编码 |
| `DecodeRuneInString(s string) (r rune, size int)` | 解码字符串的第一个字符 |
| `DecodeLastRuneInString(s string) (r rune, size int)` | 解码字符串的最后一个字符 |
| `DecodeRune(p []byte) (r rune, size int)` | 解码字节切片的第一个字符 |
| `EncodeRune(p []byte, r rune) int` | 将字符编码到字节切片 |
| `RuneLen(r rune) int` | 返回编码该字符需要的字节数 |
| `RuneStart(b byte) bool` | 判断字节是否是字符的起始字节 |
| `FullRune(p []byte) bool` | 判断字节切片是否包含完整的 UTF-8 字符 |

## 💡 最佳实践

1. **使用 `range` 遍历字符串**: Go 的 `range` 关键字自动处理 UTF-8 解码
   ```go
   for i, r := range "世界" {
       fmt.Printf("%d: %c\n", i, r)
   }
   ```

2. **不要使用索引直接访问多字节字符**: 避免 `s[i]`，使用 `utf8.DecodeRuneInString()` 或 `range`

3. **计数使用 `RuneCountInString`**: 而不是 `len()`，除非你需要字节数

4. **验证外部输入**: 使用 `ValidString()` 验证来自外部的 UTF-8 数据

## ⚠️ 常见陷阱

```go
s := "世界"

// ❌ 错误：直接索引可能截断字符
firstByte := s[0] // 只得到第一个字节，不是完整字符

// ✅ 正确：使用 UTF-8 解码
firstRune, _ := utf8.DecodeRuneInString(s)

// ❌ 错误：使用 len() 获取字符数
charCount := len(s) // 返回 6（字节数）

// ✅ 正确：使用 RuneCountInString
charCount = utf8.RuneCountInString(s) // 返回 2（字符数）
```

## 🔗 相关资源

- UTF-8 标准: https://en.wikipedia.org/wiki/UTF-8
- Go 字符串和 rune: https://go.dev/blog/strings
- Unicode 包: https://pkg.go.dev/unicode