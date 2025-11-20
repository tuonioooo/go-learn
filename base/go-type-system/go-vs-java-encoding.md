# Go vs Java 字符串编码完全指南

本教程将深入对比 Go 和 Java 中字符串的编码方式、内存占用和实际字节数计算。

## 核心区别

| 语言 | 内部编码 | 特点 |
|------|---------|------|
| **Go** | UTF-8 | 变长编码，节省内存，ASCII 占 1 字节 |
| **Java** | UTF-16 | 变长编码，BMP 字符占 2 字节 |

---

## 一、Go 语言的字符串编码（UTF-8）

### UTF-8 编码规则

Go 使用 **UTF-8 变长编码**，不同字符占用的字节数不同：

| Unicode 范围 | 字节数 | 示例字符 |
|-------------|--------|---------|
| U+0000 - U+007F | 1 字节 | ASCII：`A`、`1`、`!` |
| U+0080 - U+07FF | 2 字节 | 希腊字母、阿拉伯字母 |
| U+0800 - U+FFFF | 3 字节 | 中文：`中`、`你`、`好` |
| U+10000 - U+10FFFF | 4 字节 | Emoji：`😀`、`🌍`、`🎉` |

### 实例：`"A中😀"` 在 Go 中

```go
package main

import "fmt"

func main() {
    s := "A中😀"
    
    fmt.Printf("字符串: %s\n", s)
    fmt.Printf("字节数 len(): %d\n", len(s))                    // 8 字节
    fmt.Printf("字符数 len([]rune()): %d\n", len([]rune(s)))   // 3 个字符
    
    // 逐个字符分析
    for i, r := range s {
        fmt.Printf("字符 '%c' (U+%04X): 字节位置 %d, UTF-8 长度 %d 字节\n", 
            r, r, i, len(string(r)))
    }
}
```

**输出：**
```
字符串: A中😀
字节数 len(): 8
字符数 len([]rune()): 3
字符 'A' (U+0041): 字节位置 0, UTF-8 长度 1 字节
字符 '中' (U+4E2D): 字节位置 1, UTF-8 长度 3 字节
字符 '😀' (U+1F600): 字节位置 4, UTF-8 长度 4 字节
```

### 字节数计算

| 字符 | Unicode 码点 | UTF-8 字节数 |
|-----|-------------|-------------|
| `A` | U+0041 | 1 字节 |
| `中` | U+4E2D | 3 字节 |
| `😀` | U+1F600 | 4 字节 |
| **总计** | - | **8 字节** |

**计算公式：** `1 + 3 + 4 = 8 字节`

### Go 字符串的内存结构

```go
type StringHeader struct {
    Data uintptr  // 8 字节 (64位) - 指向数据的指针
    Len  int      // 8 字节 (64位) - 字符串字节长度
}
```

完整内存占用 = **字符串内容字节数** + **16 字节**（字符串头部）

---

## 二、Java 语言的字符串编码（UTF-16）

### UTF-16 编码规则

Java 使用 **UTF-16 变长编码**：

| Unicode 范围 | char 数量 | 字节数 | 说明 |
|-------------|----------|--------|------|
| U+0000 - U+FFFF | 1 个 | 2 字节 | 基本多文种平面（BMP） |
| U+10000 - U+10FFFF | 2 个 | 4 字节 | 需要代理对（Surrogate Pair） |

### 实例：`"A中😀"` 在 Java 中

```java
public class StringEncodingDemo {
    public static void main(String[] args) throws Exception {
        String s = "A中😀";
        
        System.out.println("字符串: " + s);
        System.out.println("length() (char数): " + s.length());
        System.out.println("codePointCount() (真实字符数): " + 
            s.codePointCount(0, s.length()));
        
        // 不同编码的字节数
        System.out.println("\n=== 字节数 ===");
        System.out.println("UTF-8:     " + s.getBytes("UTF-8").length + " 字节");
        System.out.println("UTF-16:    " + s.getBytes("UTF-16").length + " 字节 (含BOM)");
        System.out.println("UTF-16BE:  " + s.getBytes("UTF-16BE").length + " 字节 (无BOM)");
        
        // 逐个字符分析
        System.out.println("\n=== 字符分析 ===");
        for (int i = 0; i < s.length(); ) {
            int codePoint = s.codePointAt(i);
            int charCount = Character.charCount(codePoint);
            System.out.printf("字符 '%s' (U+%04X): 占用 %d 个 char, %d 字节\n",
                new String(Character.toChars(codePoint)), 
                codePoint, charCount, charCount * 2);
            i += charCount;
        }
    }
}
```

**输出：**
```
字符串: A中😀
length() (char数): 4
codePointCount() (真实字符数): 3

=== 字节数 ===
UTF-8:     8 字节
UTF-16:    10 字节 (含BOM)
UTF-16BE:  8 字节 (无BOM)

=== 字符分析 ===
字符 'A' (U+0041): 占用 1 个 char, 2 字节
字符 '中' (U+4E2D): 占用 1 个 char, 2 字节
字符 '😀' (U+1F600): 占用 2 个 char, 4 字节
```

### 三种编码方式的字节数详解

#### 1. UTF-8：8 字节

| 字符 | Unicode | UTF-8 字节数 |
|-----|---------|-------------|
| `A` | U+0041 | 1 字节 |
| `中` | U+4E2D | 3 字节 |
| `😀` | U+1F600 | 4 字节 |
| **总计** | - | **8 字节** |

**原理：** UTF-8 根据码点范围动态分配字节数

#### 2. UTF-16：10 字节（含 BOM）

| 部分 | 字节数 | 说明 |
|-----|--------|------|
| BOM | 2 字节 | 字节序标记（`FE FF` 或 `FF FE`） |
| `A` | 2 字节 | BMP 字符 |
| `中` | 2 字节 | BMP 字符 |
| `😀` | 4 字节 | 代理对（高代理 + 低代理） |
| **总计** | **10 字节** | - |

**原理：** UTF-16 默认带 BOM（Byte Order Mark），用于标识字节序

#### 3. UTF-16BE：8 字节（无 BOM）

| 字符 | 字节数 | 说明 |
|-----|--------|------|
| `A` | 2 字节 | BMP 字符 |
| `中` | 2 字节 | BMP 字符 |
| `😀` | 4 字节 | 代理对 |
| **总计** | **8 字节** | - |

**原理：** UTF-16BE（Big Endian）是固定字节序，不需要 BOM

### 为什么 Emoji 在 UTF-16 占 4 字节？

**代理对（Surrogate Pair）机制：**

1. Emoji `😀` 的码点是 `U+1F600`，超出 BMP 范围（U+0000 - U+FFFF）
2. UTF-16 使用**代理对**编码：
   - **高代理**（High Surrogate）：U+D800 - U+DBFF
   - **低代理**（Low Surrogate）：U+DC00 - U+DFFF
3. 每个代理占 1 个 `char`（2 字节），共 2 个 `char`（4 字节）

```java
String emoji = "😀";
System.out.println(emoji.length());  // 输出: 2 (两个 char)
System.out.println(emoji.codePointCount(0, emoji.length()));  // 输出: 1 (一个字符)
```

---

## 三、核心对比总结

### 字符串 `"A中😀"` 的编码对比

| 编码方式 | 字节数 | 组成 |
|---------|-------|------|
| **Go (UTF-8)** | 8 字节 | A(1) + 中(3) + 😀(4) |
| **Java UTF-8** | 8 字节 | A(1) + 中(3) + 😀(4) |
| **Java UTF-16** | 10 字节 | BOM(2) + A(2) + 中(2) + 😀(4) |
| **Java UTF-16BE** | 8 字节 | A(2) + 中(2) + 😀(4) |

### 关键差异

| 特性 | Go | Java |
|-----|-----|------|
| **内部编码** | UTF-8 | UTF-16 |
| **ASCII 字符** | 1 字节 | 2 字节（内存中） |
| **中文字符** | 3 字节 | 2 字节（内存中） |
| **Emoji** | 4 字节 | 4 字节（代理对） |
| **`len()` / `length()`** | 返回字节数 | 返回 char 数量 |
| **获取真实字符数** | `len([]rune(s))` | `s.codePointCount(0, s.length())` |
| **内存效率** | ASCII 文本更高效 | BMP 字符更紧凑 |

### 适用场景

**选择 Go (UTF-8)：**
- 处理大量英文文本（节省内存）
- 网络传输（UTF-8 是 Web 标准）
- 跨平台兼容性

**选择 Java (UTF-16)：**
- 处理大量 BMP 字符（中文、日文、韩文）
- 随机访问字符（大部分字符占 2 字节，更容易计算位置）
- Java 生态系统内部处理

---

## 四、实用代码示例

### Go 完整示例

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func analyzeString(s string) {
    fmt.Printf("字符串: %q\n", s)
    fmt.Printf("字节数: %d\n", len(s))
    fmt.Printf("字符数: %d\n", utf8.RuneCountInString(s))
    fmt.Println("详细分析:")
    
    for i, r := range s {
        size := utf8.RuneLen(r)
        fmt.Printf("  位置 %d: '%c' (U+%04X) - %d 字节\n", 
            i, r, r, size)
    }
    fmt.Println()
}

func main() {
    analyzeString("Hello")
    analyzeString("你好")
    analyzeString("A中😀")
}
```

### Java 完整示例

```java
import java.nio.charset.StandardCharsets;

public class StringAnalyzer {
    public static void analyzeString(String s) throws Exception {
        System.out.println("字符串: " + s);
        System.out.println("char 数量: " + s.length());
        System.out.println("真实字符数: " + s.codePointCount(0, s.length()));
        
        System.out.println("字节数:");
        System.out.println("  UTF-8:     " + s.getBytes("UTF-8").length);
        System.out.println("  UTF-16:    " + s.getBytes("UTF-16").length);
        System.out.println("  UTF-16BE:  " + s.getBytes("UTF-16BE").length);
        
        System.out.println("详细分析:");
        for (int i = 0; i < s.length(); ) {
            int cp = s.codePointAt(i);
            int charCount = Character.charCount(cp);
            System.out.printf("  位置 %d: '%s' (U+%04X) - %d char, %d 字节\n",
                i, new String(Character.toChars(cp)), cp, 
                charCount, charCount * 2);
            i += charCount;
        }
        System.out.println();
    }
    
    public static void main(String[] args) throws Exception {
        analyzeString("Hello");
        analyzeString("你好");
        analyzeString("A中😀");
    }
}
```

---

## 五、常见陷阱

### Go 陷阱

❌ **错误：** 使用索引访问"字符"
```go
s := "你好"
fmt.Println(s[0])  // 输出: 228 (第一个字节，不是完整字符)
```

✅ **正确：** 使用 `range` 遍历字符
```go
s := "你好"
for _, r := range s {
    fmt.Printf("%c ", r)  // 输出: 你 好
}
```

### Java 陷阱

❌ **错误：** 使用 `length()` 作为字符数
```go
String s = "😀😀😀";
System.out.println(s.length());  // 输出: 6 (不是 3!)
```

✅ **正确：** 使用 `codePointCount()`
```go
String s = "😀😀😀";
System.out.println(s.codePointCount(0, s.length()));  // 输出: 3
```

---

## 总结

1. **Go 使用 UTF-8**，内存效率高，适合 ASCII 为主的场景
2. **Java 使用 UTF-16**，BMP 字符访问效率高，适合中日韩文本
3. **Emoji 等补充字符**在两种编码中都需要额外处理
4. **理解编码差异**是跨语言文本处理的关键
5. **永远使用正确的 API** 获取字符数量，而不是字节/char 数量

---

**参考资源：**
- [UTF-8 规范](https://www.rfc-editor.org/rfc/rfc3629)
- [UTF-16 规范](https://www.rfc-editor.org/rfc/rfc2781)
- [Go Strings, bytes, runes and characters](https://go.dev/blog/strings)
- [Java Character 文档](https://docs.oracle.com/javase/8/docs/api/java/lang/Character.html)