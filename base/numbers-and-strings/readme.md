# Go 字符串练习题集

## 1. 字符串字面量与转义字符

### 问题：以下表达式的结果是什么？
```go
"\"Hello\\"" + ` \"World\"`
```

**选项：**
1. `"Hello" "World"`
2. `"Hello" \"World\"` ✅
3. `"Hello" `"World"``
4. `"\"Hello\" `\"World\""`

**解析：**
- Go 会将双引号字符串中的 `\"` 解释为 `"`
- 但在反引号（原始字符串字面量）中不会解释转义序列
- 因此 `` ` \"World\"` `` 中的 `\"` 保持原样

---

## 2. 多行字符串的最佳实践

### 问题：表示以下 XML 文本的最佳方式是什么？
```xml
<xml>
  <items>
    <item>"Teddy Bear"</item>
  </items>
</xml>
```

**正确答案：** ✅
```go
`<xml>
    <items>
        <item>"Teddy Bear"</item>
    </items>
</xml>`
```

**错误选项：**

2-3. 普通字符串字面量不能跨多行：
```go
"<xml>
    <items>
        <item>"Teddy Bear"</item>
    </items>
</xml>"
```

4. 在原始字符串字面量中不需要转义：
```go
`<xml>
    <items>
        <item>\"Teddy Bear\"</item>
    </items>
</xml>`
```

**要点：**
- 使用反引号 `` ` `` 创建多行字符串
- 反引号内不需要转义双引号

---

## 3. 字符串长度基础

### 问题：以下表达式的结果是什么？
```go
len("lovely")
```

**选项：**
1. 7
2. 8
3. 6 ✅
4. 0

**解析：**
- `len()` 函数返回字符串的**字节数**
- "lovely" 包含 6 个 ASCII 字符，每个 1 字节 = 6 字节

---

## 4. 原始字符串的长度计算

### 问题：以下表达式的结果是什么？
```go
len("very") + len(`\"cool\"`)
```

**选项：**
1. 8
2. 12 ✅
3. 16
4. 10

**解析：**
- `"very"` = 4 字节
- `` `\"cool\"` `` = 8 字节（Go 不解释原始字符串中的转义序列，`\` 和 `"` 都算独立字符）
- 总计：4 + 8 = 12

---

## 5. 转义字符串的长度计算

### 问题：以下表达式的结果是什么？
```go
len("very") + len("\"cool\"")
```

**选项：**
1. 8
2. 12
3. 16
4. 10 ✅

**解析：**
- `"very"` = 4 字节
- `"\"cool\""` 中，`\"` 被解释为 `"`（1 字节），所以是 `"cool"` = 6 字节
- 总计：4 + 6 = 10

---

## 6. UTF-8 字符的字节长度

### 问题：以下表达式的结果是什么？
```go
len("péripatéticien")
```

**提示：** é 占用 2 字节，`len()` 计算字节数而非字符数

**选项：**
1. 14
2. 16 ✅
3. 18
4. 20

**解析：**
- 英文字母 = 1 字节/个
- é = 2 字节/个
- "péripatéticien" 有 12 个英文字母 + 2 个 é
- 总计：12×1 + 2×2 = 16 字节

---

## 7. 正确计算字符数量

### 问题：如何正确计算此字符串的字符数量？
```go
"péripatéticien"
```

**选项：**
1. `len(péripatéticien)`
2. `len("péripatéticien")`
3. `utf8.RuneCountInString("péripatéticien")` ✅
4. `unicode/utf8.RuneCountInString("péripatéticien")`

**解析：**
- `len()` 只能计算字节数
- `utf8.RuneCountInString()` 计算 rune（字符）数量
- 包名是 `utf8`，不是 `unicode/utf8`

---

## 8. Rune 计数

### 问题：以下表达式的结果是什么？
```go
utf8.RuneCountInString("péripatéticien")
```

**选项：**
1. 16
2. 14 ✅
3. 18
4. 20

**解析：**
- 16 是字节数
- 14 是字符（rune）数
- `RuneCountInString()` 返回 Unicode 字符数量，不是字节数

---

## 9. 字符串操作包

### 问题：哪个包包含字符串操作函数？

**选项：**
1. string
2. unicode/utf8
3. strings ✅
4. unicode/strings

**解析：**
- `strings` 包提供字符串搜索、替换、分割等功能
- `unicode/utf8` 处理 UTF-8 编码
- 没有 `string` 或 `unicode/strings` 包

---

## 10. Repeat 函数

### 问题：以下表达式的结果是什么？
```go
strings.Repeat("*x", 3) + "*"
```

**选项：**
1. `*x*x*x`
2. `x*x*x`
3. `*x3`
4. `*x*x*x*` ✅

**解析：**
- `strings.Repeat("*x", 3)` 重复字符串 3 次 = `*x*x*x`
- 连接 `"*"` = `*x*x*x*`

---

## 11. ToUpper 函数

### 问题：以下表达式的结果是什么？
```go
strings.ToUpper("bye bye ") + "see you!"
```

**选项：**
1. `bye bye see you!`
2. `BYE BYE SEE YOU!`
3. `bye bye + see you!`
4. `BYE BYE see you!` ✅

**解析：**
- `strings.ToUpper("bye bye ")` 只转换第一部分 = `BYE BYE `
- 后面的 `"see you!"` 保持原样
- 结果：`BYE BYE see you!`

---

## 字符串的底层结构

Go 中的字符串在底层是一个结构体，定义在 `runtime` 包中：

```go
type stringStruct struct {
    str unsafe.Pointer  // 指向底层字节数组的指针
    len int             // 字符串的长度（字节数）
}
```

## 占用的字节数

**字符串类型本身占用 16 个字节**（在 64 位系统上）：

- **指针部分**：8 字节（`unsafe.Pointer`）
- **长度部分**：8 字节（`int` 类型在 64 位系统上是 8 字节）

在 32 位系统上则占用 8 个字节（指针 4 字节 + int 4 字节）。

## 关键特性

1. **字符串头部是固定大小**：无论字符串内容多长，字符串变量本身始终占用 16 字节（64位系统）

2. **实际数据存储在别处**：指针指向的底层字节数组存储在堆或只读数据段中，不计入字符串结构体本身的大小

3. **不可变性**：Go 的字符串是不可变的，多个字符串可以安全地共享同一个底层字节数组

4. **高效传递**：由于字符串只是 16 字节的结构体，传递字符串时只复制这个小结构，不会复制底层数据

示例验证：

```go
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var s1 string = "hello"
    var s2 string = "this is a very long string"
    
    fmt.Println("Size of string type:", unsafe.Sizeof(s1))  // 输出: 16
    fmt.Println("Size of string type:", unsafe.Sizeof(s2))  // 输出: 16
    fmt.Println("Length of s1:", len(s1))                   // 输出: 5
    fmt.Println("Length of s2:", len(s2))                   // 输出: 26
}
```

---

## 关键要点总结

### 字符串字面量
- **双引号字符串**：解释转义序列（`\n`, `\"`, `\t` 等）
- **反引号字符串**：原始字符串，不解释转义序列

### 长度计算
- `len()`：返回**字节数**
- `utf8.RuneCountInString()`：返回**字符数**（rune）

### 常用包
- `strings`：字符串操作（Repeat, ToUpper, Split 等）
- `unicode/utf8`：UTF-8 编码处理

### 注意事项
- ASCII 字符 = 1 字节
- 非 ASCII 字符（如 é）= 2+ 字节
- 多行字符串优先使用反引号