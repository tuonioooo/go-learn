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