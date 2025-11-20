# Go 语言数值类型详解

本教程将详细介绍 Go 语言中的数值类型，包括它们的取值范围和内存占用。

## 整数类型 (Integer Types)

### 有符号整数 (Signed Integers)

有符号整数可以表示正数、负数和零。Go 提供了四种不同大小的有符号整数类型：

| 类型 | 最小值 | 最大值 | 内存占用 |
|------|--------|--------|----------|
| `int8` | -128 | 127 | 1 字节 |
| `int16` | -32,768 | 32,767 | 2 字节 |
| `int32` | -2,147,483,648 | 2,147,483,647 | 4 字节 |
| `int64` | -9,223,372,036,854,775,808 | 9,223,372,036,854,775,807 | 8 字节 |

```go
fmt.Println("int8   :", math.MinInt8, math.MaxInt8)
fmt.Println("int16  :", math.MinInt16, math.MaxInt16)
fmt.Println("int32  :", math.MinInt32, math.MaxInt32)
fmt.Println("int64  :", math.MinInt64, math.MaxInt64)
```

### 无符号整数 (Unsigned Integers)

无符号整数只能表示零和正整数，不能表示负数。由于不需要存储符号位，无符号整数可以表示更大的正数：

| 类型 | 最小值 | 最大值 | 内存占用 |
|------|--------|--------|----------|
| `uint8` | 0 | 255 | 1 字节 |
| `uint16` | 0 | 65,535 | 2 字节 |
| `uint32` | 0 | 4,294,967,295 | 4 字节 |
| `uint64` | 0 | 18,446,744,073,709,551,615 | 8 字节 |

```go
fmt.Println("uint8  :", 0, math.MaxUint8)
fmt.Println("uint16 :", 0, math.MaxUint16)
fmt.Println("uint32 :", 0, math.MaxUint32)
```

**注意**：`math.MaxUint64` 不能直接打印，需要进行类型转换：

```go
fmt.Println("uint64 :", 0, uint64(math.MaxUint64))
```

### 平台相关的整数类型

`int` 和 `uint` 的大小取决于你的操作系统架构：
- 在 32 位系统上：4 字节
- 在 64 位系统上：8 字节

```go
fmt.Println("int    :", unsafe.Sizeof(int(1)), "bytes")
fmt.Println("uint   :", unsafe.Sizeof(uint(1)), "bytes")
```

## 浮点数类型 (Floating-Point Types)

浮点数用于表示带小数点的数值。Go 提供两种浮点数类型：

| 类型 | 最小非零值 | 最大值 | 内存占用 | 精度 |
|------|------------|--------|----------|------|
| `float32` | ~1.4e-45 | ~3.4e38 | 4 字节 | 单精度 |
| `float64` | ~4.9e-324 | ~1.8e308 | 8 字节 | 双精度 |

```go
fmt.Println("float32:", math.SmallestNonzeroFloat32, math.MaxFloat32)
fmt.Println("float64:", math.SmallestNonzeroFloat64, math.MaxFloat64)
```

### 科学计数法说明

在浮点数中，`e` 表示 10 的幂次方（科学计数法）：
- `1e1` = 10
- `1e2` = 100
- `12e1` = 120
- `1.4e-45` = 1.4 × 10⁻⁴⁵

## 内存占用 (Memory Costs)

使用 `unsafe.Sizeof()` 可以查看不同类型在内存中占用的字节数：

```go
fmt.Println("int8   :", unsafe.Sizeof(int8(1)), "bytes")    // 1 字节
fmt.Println("int16  :", unsafe.Sizeof(int16(1)), "bytes")   // 2 字节
fmt.Println("int32  :", unsafe.Sizeof(int32(1)), "bytes")   // 4 字节
fmt.Println("int64  :", unsafe.Sizeof(int64(1)), "bytes")   // 8 字节
fmt.Println("float32:", unsafe.Sizeof(float32(1)), "bytes") // 4 字节
fmt.Println("float64:", unsafe.Sizeof(float64(1)), "bytes") // 8 字节
```

## 字符串的内存占用

字符串在 Go 中的内存占用 = 字符串长度 + 8 字节（字符串头部信息）：

```go
fmt.Println("hello  :", len("hello")+8, "bytes") // 5 + 8 = 13 字节
fmt.Println("hi     :", len("hi")+8, "bytes")    // 2 + 8 = 10 字节
```

这 8 字节存储了字符串的指针和长度信息。

## 如何选择合适的类型？

1. **优先使用 `int`**：对于一般的整数运算，使用 `int` 类型即可
2. **需要节省内存时**：如果处理大量数据，可以选择更小的类型（如 `int8`、`int16`）
3. **需要明确范围时**：当你知道数值不会超过某个范围时，选择对应的类型
4. **浮点数默认用 `float64`**：除非有特殊的内存限制，否则推荐使用 `float64` 以获得更高精度
5. **需要与外部系统交互时**：根据外部系统的要求选择固定大小的类型（如 `int32`、`uint64`）

## 完整示例代码

```go
package main

import (
    "fmt"
    "math"
    "unsafe"
)

func main() {
    // 查看整数类型的范围
    fmt.Println("=== 有符号整数 ===")
    fmt.Println("int8   :", math.MinInt8, "到", math.MaxInt8)
    fmt.Println("int16  :", math.MinInt16, "到", math.MaxInt16)
    fmt.Println("int32  :", math.MinInt32, "到", math.MaxInt32)
    fmt.Println("int64  :", math.MinInt64, "到", math.MaxInt64)
    
    fmt.Println("\n=== 无符号整数 ===")
    fmt.Println("uint8  :", 0, "到", math.MaxUint8)
    fmt.Println("uint16 :", 0, "到", math.MaxUint16)
    fmt.Println("uint32 :", 0, "到", math.MaxUint32)
    fmt.Println("uint64 :", 0, "到", uint64(math.MaxUint64))
    
    fmt.Println("\n=== 浮点数 ===")
    fmt.Println("float32:", math.SmallestNonzeroFloat32, "到", math.MaxFloat32)
    fmt.Println("float64:", math.SmallestNonzeroFloat64, "到", math.MaxFloat64)
    
    fmt.Println("\n=== 内存占用 ===")
    fmt.Println("int8   :", unsafe.Sizeof(int8(1)), "字节")
    fmt.Println("int64  :", unsafe.Sizeof(int64(1)), "字节")
    fmt.Println("float32:", unsafe.Sizeof(float32(1)), "字节")
    fmt.Println("float64:", unsafe.Sizeof(float64(1)), "字节")
}
```

## 总结

- Go 提供了丰富的数值类型，从 1 字节的 `int8` 到 8 字节的 `int64`
- 选择合适的类型可以在性能和内存使用之间取得平衡
- 使用 `math` 包可以方便地获取各类型的边界值
- 使用 `unsafe.Sizeof()` 可以查看类型的内存占用

---

**版权声明**：本教程基于 Inanc Gumus 的 Go 编程课程代码改编  
**许可协议**：[CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/)