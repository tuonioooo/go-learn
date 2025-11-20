# Go 语言常量类型教程

## 目录
- [什么是常量](#什么是常量)
- [常量的声明](#常量的声明)
- [常量类型](#常量类型)
- [iota 枚举器](#iota-枚举器)
- [常量表达式](#常量表达式)
- [类型化和非类型化常量](#类型化和非类型化常量)
- [最佳实践](#最佳实践)

## 什么是常量

常量是一个简单值的标识符,在程序运行期间不会被修改。常量中的数据类型只能是布尔型、数字型(整数型、浮点型和复数)和字符串型。

## 常量的声明

### 单个常量声明

```go
const Pi = 3.14159
const Language = "Go"
const IsActive = true
```

### 批量声明

```go
const (
    StatusOK = 200
    StatusNotFound = 404
    StatusInternalServerError = 500
)
```

### 显式类型声明

```go
const Pi float64 = 3.14159265359
const MaxConnections int = 100
const AppName string = "MyApp"
```

## 常量类型

### 布尔型常量

```go
const (
    Enabled = true
    Disabled = false
)
```

### 数字型常量

```go
// 整数
const MaxSize = 1024
const MinValue int32 = -100

// 浮点数
const Epsilon = 0.0001
const Rate float32 = 0.05

// 复数
const ComplexNum = 2 + 3i
```

### 字符串常量

```go
const (
    Welcome = "欢迎使用 Go 语言"
    EmptyString = ""
    MultiLine = `这是一个
多行字符串
常量`
)
```

## iota 枚举器

`iota` 是 Go 语言的常量计数器,只能在常量表达式中使用。`iota` 在 const 关键字出现时将被重置为 0,每新增一行常量声明将使 `iota` 计数一次。

### 基本用法

```go
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)
```

### 跳过某些值

```go
const (
    A = iota    // 0
    _           // 1 (跳过)
    C           // 2
    D           // 3
)
```

### 使用表达式

```go
const (
    B1 = 1 << (10 * iota)  // 1 (1 << 0)
    KB                      // 1024 (1 << 10)
    MB                      // 1048576 (1 << 20)
    GB                      // 1073741824 (1 << 30)
)
```

### 自定义起始值

```go
const (
    ErrorCodeBase = 1000
    ErrorNotFound = ErrorCodeBase + iota  // 1000
    ErrorUnauthorized                     // 1001
    ErrorForbidden                        // 1002
)
```

### 多个常量同行

```go
const (
    A, B = iota + 1, iota + 2  // A=1, B=2
    C, D                        // C=2, D=3
    E, F                        // E=3, F=4
)
```

### 常量自动继承

在 `const` 声明块中,如果不显式指定常量的值和类型,该常量会自动继承上一行的表达式和类型。这是 Go 语言常量的一个重要特性。

#### 基本继承

```go
const (
    A = 100      // 100
    B            // 100 (继承 A 的值)
    C            // 100 (继承 A 的值)
)
```

#### 继承表达式

```go
const (
    Size = 10 * 2    // 20
    Limit            // 20 (继承表达式 10 * 2)
    Max              // 20 (继承表达式 10 * 2)
)
```

#### 继承类型

```go
const (
    Pi float64 = 3.14159    // float64 类型
    E                        // float64 类型,继承类型和表达式
    Phi                      // float64 类型,继承类型和表达式
)
```

#### 结合 iota 的继承

```go
const (
    A = iota * 10    // 0 (0 * 10)
    B                // 10 (1 * 10)
    C                // 20 (2 * 10)
    D                // 30 (3 * 10)
)
```

#### 复杂表达式继承

```go
const (
    // 计算字节大小
    _ = iota
    KB = 1 << (10 * iota)  // 1024 (1 << 10)
    MB                      // 1048576 (1 << 20)
    GB                      // 1073741824 (1 << 30)
    TB                      // 1099511627776 (1 << 40)
)
```

#### 多值继承

```go
const (
    X, Y = 1, 2      // X=1, Y=2
    M, N             // M=1, N=2 (继承上一行的表达式)
    P, Q             // P=1, Q=2 (继承上一行的表达式)
)
```

#### 结合 iota 的多值继承

```go
const (
    A, B = iota, iota + 10   // A=0, B=10
    C, D                      // C=1, D=11 (iota 递增)
    E, F                      // E=2, F=12 (iota 递增)
)
```

#### 实际应用示例

```go
// 定义不同级别的日志优先级
const (
    LogLevelDebug = iota    // 0
    LogLevelInfo            // 1
    LogLevelWarn            // 2
    LogLevelError           // 3
    LogLevelFatal           // 4
)

// 定义文件权限(Unix 风格)
const (
    PermOwnerRead = 1 << iota  // 1 (二进制: 001)
    PermOwnerWrite              // 2 (二进制: 010)
    PermOwnerExec               // 4 (二进制: 100)
    PermGroupRead               // 8
    PermGroupWrite              // 16
    PermGroupExec               // 32
    PermOtherRead               // 64
    PermOtherWrite              // 128
    PermOtherExec               // 256
)

// 定义时间间隔
const (
    Second = 1
    Minute = 60 * Second     // 60
    Hour                      // 60 (继承 60 * Second,但 iota 不递增)
    Day                       // 60 (继承 60 * Second)
)

// 正确的时间间隔定义
const (
    _Second = 1
    _Minute = 60 * _Second   // 60
    _Hour = 60 * _Minute     // 3600
    _Day = 24 * _Hour        // 86400
)
```

#### 注意事项

```go
// 错误示例:继承会导致意外结果
const (
    A = 1
    B = 2
    C      // C = 2 (继承 B 的值,而不是递增)
    D      // D = 2 (继承 B 的值)
)

// 正确示例:需要递增时使用 iota
const (
    A = iota + 1  // 1
    B             // 2 (iota 递增)
    C             // 3 (iota 递增)
    D             // 4 (iota 递增)
)
```

#### 完整示例

```go
package main

import "fmt"

const (
    // 基本继承
    BaseValue = 100
    Value1           // 100
    Value2           // 100
)

const (
    // 表达式继承
    Square = 5 * 5
    Area              // 25
    Size              // 25
)

const (
    // 类型继承
    Rate float64 = 0.05
    Tax                  // 0.05, float64 类型
    Fee                  // 0.05, float64 类型
)

const (
    // iota 结合继承
    Code = iota + 100  // 100
    Code2              // 101
    Code3              // 102
)

const (
    // 位运算继承
    Flag = 1 << iota  // 1
    Flag2             // 2
    Flag3             // 4
    Flag4             // 8
)

func main() {
    fmt.Println("基本继承:")
    fmt.Printf("BaseValue=%d, Value1=%d, Value2=%d\n", BaseValue, Value1, Value2)
    
    fmt.Println("\n表达式继承:")
    fmt.Printf("Square=%d, Area=%d, Size=%d\n", Square, Area, Size)
    
    fmt.Println("\n类型继承:")
    fmt.Printf("Rate=%.2f, Tax=%.2f, Fee=%.2f\n", Rate, Tax, Fee)
    
    fmt.Println("\niota 结合继承:")
    fmt.Printf("Code=%d, Code2=%d, Code3=%d\n", Code, Code2, Code3)
    
    fmt.Println("\n位运算继承:")
    fmt.Printf("Flag=%d, Flag2=%d, Flag3=%d, Flag4=%d\n", Flag, Flag2, Flag3, Flag4)
}
```

## 常量表达式

常量可以参与各种表达式计算:

```go
const (
    // 算术运算
    Result = (10 + 5) * 2 - 3  // 27
    
    // 位运算
    Flag = 1 << 2  // 4
    
    // 字符串操作
    FullName = "Go" + " " + "Language"  // "Go Language"
    
    // 函数调用(限于内置函数)
    Length = len("Hello")  // 5
)
```

## 类型化和非类型化常量

### 非类型化常量

没有明确指定类型的常量称为非类型化常量,它们具有更高的灵活性:

```go
const Number = 123  // 非类型化常量

var i int = Number
var f float64 = Number
var c complex128 = Number
// 以上都合法,Number 可以适配不同类型
```

### 类型化常量

明确指定类型的常量:

```go
const TypedNumber int = 123  // 类型化常量

var i int = TypedNumber        // 合法
// var f float64 = TypedNumber // 编译错误:类型不匹配
```

### 数值精度

非类型化常量可以保持高精度:

```go
const (
    // 非类型化常量,保持高精度
    BigNumber = 1e1000
    
    // 使用时才转换为具体类型
    Result = BigNumber / 1e999  // 10
)
```

## 最佳实践

### 1. 使用有意义的命名

```go
// 好的命名
const MaxRetryAttempts = 3
const DefaultTimeout = 30

// 不好的命名
const N = 3
const T = 30
```

### 2. 将相关常量分组

```go
const (
    // HTTP 状态码
    StatusOK = 200
    StatusCreated = 201
    StatusBadRequest = 400
)

const (
    // 配置项
    DefaultPort = 8080
    MaxConnections = 100
)
```

### 3. 使用 iota 创建枚举

```go
type Status int

const (
    StatusPending Status = iota
    StatusRunning
    StatusCompleted
    StatusFailed
)
```

### 4. 避免魔法数字

```go
// 不推荐
if age > 18 {
    // ...
}

// 推荐
const LegalAge = 18

if age > LegalAge {
    // ...
}
```

### 5. 使用常量而非变量

当值不需要改变时,优先使用常量:

```go
// 推荐
const Pi = 3.14159

// 不推荐(对于不变的值)
var pi = 3.14159
```

## 完整示例

```go
package main

import "fmt"

// 定义应用常量
const (
    AppName    = "MyApplication"
    AppVersion = "1.0.0"
    MaxUsers   = 1000
)

// 使用 iota 定义权限级别
type Permission int

const (
    PermissionRead Permission = 1 << iota
    PermissionWrite
    PermissionExecute
    PermissionAll = PermissionRead | PermissionWrite | PermissionExecute
)

// 定义配置常量
const (
    ConfigTimeout    = 30
    ConfigRetries    = 3
    ConfigBufferSize = 1024
)

func main() {
    fmt.Printf("应用名称: %s\n", AppName)
    fmt.Printf("版本: %s\n", AppVersion)
    fmt.Printf("最大用户数: %d\n", MaxUsers)
    
    fmt.Printf("\n权限值:\n")
    fmt.Printf("读取: %d\n", PermissionRead)
    fmt.Printf("写入: %d\n", PermissionWrite)
    fmt.Printf("执行: %d\n", PermissionExecute)
    fmt.Printf("全部: %d\n", PermissionAll)
    
    // 常量可用于数组长度
    var buffer [ConfigBufferSize]byte
    fmt.Printf("\n缓冲区大小: %d 字节\n", len(buffer))
}
```

## 官方文档链接

- **Go 语言规范 - 常量**: https://go.dev/ref/spec#Constants
- **Effective Go - 常量**: https://go.dev/doc/effective_go#constants
- **Go 官方文档**: https://go.dev/doc/
- **Go 博客 - Constants**: https://go.dev/blog/constants

## 总结

常量是 Go 语言中重要的特性,合理使用常量可以:
- 提高代码可读性
- 避免魔法数字
- 提供编译时类型检查
- 优化性能(编译时求值)

掌握常量的使用,特别是 `iota` 的灵活运用,能让你的 Go 代码更加优雅和高效。