# Go 函数声明和使用完整教程

## 1. 函数的基本声明

### 1.1 基础语法
```go
func 函数名(参数列表) 返回值类型 {
    // 函数体
}
```

### 1.2 简单函数示例

#### 无参数、无返回值
```go
func Greet() {
    fmt.Println("Hello, World!")
}

// 调用
Greet()
```

#### 有参数、无返回值
```go
func Greet(name string) {
    fmt.Println("Hello, " + name)
}

// 调用
Greet("Alice")
```

#### 无参数、有返回值
```go
func GetMessage() string {
    return "Hello, World!"
}

// 调用
message := GetMessage()
fmt.Println(message)
```

#### 有参数、有返回值
```go
func Add(a int, b int) int {
    return a + b
}

// 调用
result := Add(5, 3)
fmt.Println(result)  // 输出: 8
```

#### 在go中，函数名称首字母大写是公有方法，首字母小写是私有方法。

```go
// 公共方法可以被外部调用
func Add(a int, b int) int {
    return a + b
}
// 私有方法不可以被外部调用
func add(a int, b int) int {
    return a + b
}

```

### 1.3 参数类型简写
当多个参数类型相同时，可以简写：

```go
// 完整写法
func Add(a int, b int) int {
    return a + b
}

// 简写写法
func Add(a, b int) int {
    return a + b
}
```

## 2. 多返回值函数

Go 语言支持函数返回多个值，这是处理错误的常见模式。

### 2.1 基本用法
```go
// 返回两个值
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// 调用
result, err := Divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)  // 输出: 5
}
```

### 2.2 命名返回值
```go
// 使用命名返回值
func Divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = fmt.Errorf("division by zero")
        return  // 自动返回命名的返回值
    }
    result = a / b
    return
}

// 调用方式相同
result, err := Divide(10, 2)
```

### 2.3 忽略某个返回值
```go
// 使用 _ 忽略不需要的返回值
result, _ := Divide(10, 2)
fmt.Println(result)

// 或者都忽略
_, err := Divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
}
```

## 3. 可变参数函数

使用 `...` 语法定义可接收任意数量参数的函数。

### 3.1 基本用法
```go
// 可接收任意数量的整数
func Sum(numbers ...int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}

// 调用
fmt.Println(Sum(1, 2, 3))           // 输出: 6
fmt.Println(Sum(1, 2, 3, 4, 5))     // 输出: 15
fmt.Println(Sum())                   // 输出: 0
```

### 3.2 混合固定参数和可变参数
```go
// 可变参数必须在最后
func Format(prefix string, items ...string) string {
    result := prefix + ": "
    for i, item := range items {
        if i > 0 {
            result += ", "
        }
        result += item
    }
    return result
}

// 调用
fmt.Println(Format("Items", "apple", "banana", "orange"))
// 输出: Items: apple, banana, orange
```

### 3.3 展开切片作为可变参数
```go
numbers := []int{1, 2, 3, 4, 5}
result := Sum(numbers...)  // 使用 ... 展开切片
fmt.Println(result)  // 输出: 15
```

## 4. 函数作为参数和返回值

Go 函数是一等对象，可以作为参数或返回值。

### 4.1 函数作为参数
```go
// 定义接收函数作为参数的函数
func ApplyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

// 定义操作函数
func Add(a, b int) int {
    return a + b
}

func Multiply(a, b int) int {
    return a * b
}

// 调用
fmt.Println(ApplyOperation(5, 3, Add))       // 输出: 8
fmt.Println(ApplyOperation(5, 3, Multiply))  // 输出: 15
```

### 4.2 匿名函数
```go
// 定义匿名函数作为参数
fmt.Println(ApplyOperation(5, 3, func(a, b int) int {
    return a - b
}))  // 输出: 2
```

### 4.3 函数作为返回值
```go
// 返回一个函数
func MakeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// 调用
double := MakeMultiplier(2)
triple := MakeMultiplier(3)

fmt.Println(double(5))  // 输出: 10
fmt.Println(triple(5))  // 输出: 15
```

## 5. 方法（结构体关联的函数）

方法是与特定类型关联的函数，具体如下：

```
func + 接收器 + 【方法名】 + 参数 + 返回值 + 方法体
```

### 方法 vs 函数对比

```go
// 普通函数 - 没有接收器
func Add(a, b int) int {
    return a + b
}
// 调用: result := Add(5, 3)

// 方法 - 有接收器
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
// 调用: result := rect.Area()
```

| 特性 | 普通函数 | 方法 |
|------|---------|------|
| 有没有接收器 | ❌ 无 | ✅ 有 |
| 定义位置 | 包级别 | 与特定类型关联 |
| 调用方式 | `FunctionName(args)` | `instance.MethodName(args)` |
| 能否修改数据 | 仅当使用指针参数时 | 指针接收器可修改 |
| 示例 | `func Greet(name string)` | `func (p Person) Introduce()` |

### 5.1 值接收器方法
```go
type Rectangle struct {
    Width  float64
    Height float64
}

// 值接收器方法
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// 调用
rect := Rectangle{Width: 4, Height: 5}
fmt.Println(rect.Area())       // 输出: 20
fmt.Println(rect.Perimeter())  // 输出: 18
```

### 5.2 指针接收器方法
```go
type Counter struct {
    count int
}

// 指针接收器方法 - 可以修改接收器
func (c *Counter) Increment() {
    c.count++
}

func (c *Counter) Decrement() {
    c.count--
}

func (c Counter) GetCount() int {
    return c.count
}

// 调用
counter := Counter{}
counter.Increment()
counter.Increment()
fmt.Println(counter.GetCount())  // 输出: 2
```

### 5.3 在结构体上定义多个方法
```go
type Person struct {
    Name string
    Age  int
}

// 方法 1：获取信息
func (p Person) Introduce() string {
    return fmt.Sprintf("I'm %s and I'm %d years old", p.Name, p.Age)
}

// 方法 2：成年检查
func (p Person) IsAdult() bool {
    return p.Age >= 18
}

// 方法 3：修改年龄（指针接收器）
func (p *Person) HaveBirthday() {
    p.Age++
}

// 调用
person := Person{Name: "Alice", Age: 20}
fmt.Println(person.Introduce())  // 输出: I'm Alice and I'm 20 years old
fmt.Println(person.IsAdult())    // 输出: true
person.HaveBirthday()
fmt.Println(person.Age)          // 输出: 21
```

## 6. 闭包

闭包是访问其外层函数作用域中变量的函数。

### 6.1 基本闭包
```go
func MakeCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// 调用
counter1 := MakeCounter()
counter2 := MakeCounter()

fmt.Println(counter1())  // 输出: 1
fmt.Println(counter1())  // 输出: 2
fmt.Println(counter2())  // 输出: 1
fmt.Println(counter1())  // 输出: 3
```

### 6.2 闭包修改外层变量
```go
func MakeAdder(x int) func(int) int {
    // 闭包捕获 x
    return func(y int) int {
        return x + y
    }
}

// 调用
add5 := MakeAdder(5)
add10 := MakeAdder(10)

fmt.Println(add5(3))   // 输出: 8
fmt.Println(add10(3))  // 输出: 13
```

## 7. 实际应用示例

### 7.1 Repeat 函数示例
```go
// 重复字符串指定次数
func Repeat(character string, count int) string {
    var repeated string
    for i := 0; i < count; i++ {
        repeated += character
    }
    return repeated
}

// 调用
fmt.Println(Repeat("a", 5))  // 输出: aaaaa
fmt.Println(Repeat("ab", 3)) // 输出: ababab
```

### 7.2 包含多个结构体方法的完整示例
```go
// 定义结构体
type Account struct {
    Owner   string
    Balance float64
}

// 显示账户信息
func (a Account) Display() string {
    return fmt.Sprintf("Account Owner: %s, Balance: $%.2f", a.Owner, a.Balance)
}

// 存款（指针接收器）
func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return fmt.Errorf("deposit amount must be positive")
    }
    a.Balance += amount
    return nil
}

// 取款（指针接收器）
func (a *Account) Withdraw(amount float64) error {
    if amount <= 0 {
        return fmt.Errorf("withdraw amount must be positive")
    }
    if amount > a.Balance {
        return fmt.Errorf("insufficient funds")
    }
    a.Balance -= amount
    return nil
}

// 获取余额（值接收器）
func (a Account) GetBalance() float64 {
    return a.Balance
}

// 使用示例
func main() {
    account := Account{Owner: "Alice", Balance: 1000}
    
    fmt.Println(account.Display())  // 输出: Account Owner: Alice, Balance: $1000.00
    
    account.Deposit(500)
    fmt.Println(account.GetBalance())  // 输出: 1500
    
    err := account.Withdraw(2000)
    if err != nil {
        fmt.Println("Error:", err)  // 输出: Error: insufficient funds
    }
    
    account.Withdraw(300)
    fmt.Println(account.Display())  // 输出: Account Owner: Alice, Balance: $1200.00
}
```

## 8. 函数声明方式总结表

| 方式 | 语法 | 使用场景 | 示例 |
|------|------|----------|------|
| 普通函数 | `func name(params) return` | 基本功能 | `func Add(a, b int) int` |
| 多返回值 | `func name(params) (type1, type2)` | 错误处理 | `func Divide(a, b float64) (float64, error)` |
| 命名返回值 | `func name(params) (name type)` | 提高可读性 | `func Divide() (result float64, err error)` |
| 可变参数 | `func name(params ...type)` | 不定数量参数 | `func Sum(numbers ...int) int` |
| 函数参数 | `func name(f func(...))` | 回调、装饰器 | `func ApplyOp(a, b int, op func(int, int) int)` |
| 函数返回 | `func name() func(...)` | 工厂函数、闭包 | `func MakeAdder(x int) func(int) int` |
| 方法（值接收器） | `func (r Type) name()` | 只读操作 | `func (r Rectangle) Area() float64` |
| 方法（指针接收器） | `func (r *Type) name()` | 修改状态 | `func (c *Counter) Increment()` |
| 匿名函数 | `func(params) returns { ... }` | 一次性使用 | `func(x int) { fmt.Println(x) }(5)` |
| 闭包 | 访问外层变量的函数 | 状态保存 | `func() { count++ }` |

## 9. 最佳实践

1. **使用清晰的函数名**
   ```go
   // ✅ 好的
   func CalculateTotal(items []Item) float64 { ... }
   
   // ❌ 不好
   func calc(i []Item) float64 { ... }
   ```

2. **错误处理优先**
   ```go
   // ✅ 好的
   func Process(data string) (string, error) {
       if data == "" {
           return "", fmt.Errorf("data cannot be empty")
       }
       // ...
   }
   ```

3. **选择合适的接收器类型**
   ```go
   // ✅ 修改状态时用指针接收器
   func (a *Account) Deposit(amount float64) error { ... }
   
   // ✅ 只读操作用值接收器
   func (a Account) GetBalance() float64 { ... }
   ```

4. **避免过长的参数列表**
   ```go
   // ❌ 参数过多
   func Process(a, b, c, d, e, f int) { ... }
   
   // ✅ 使用结构体
   type Config struct {
       A, B, C, D, E, F int
   }
   func Process(config Config) { ... }
   ```

5. **接口而非具体类型**
   ```go
   // ❌ 依赖具体类型
   func Save(f *os.File, data []byte) error { ... }
   
   // ✅ 依赖接口
   func Save(w io.Writer, data []byte) error { ... }
   ```
---

## init 函数与 import 机制 函数定义规则

### init 函数
- 可以在 `package main` 中定义
- 可以在其他 package 中定义
- 同一个 package 中可以出现多次
- 定义时不能有任何参数和返回值
- **最佳实践**：建议每个文件只写一个 init 函数，以提高可读性和可维护性

### main 函数
- 只能在 `package main` 中定义
- 定义时不能有任何参数和返回值
- 每个可执行程序必须包含一个 main 函数

### 执行顺序

Go 程序的初始化和执行遵循以下顺序：

#### 1. 包导入阶段
- 程序从 `main` 包开始
- 递归导入 main 包依赖的所有包
- 每个包只会被导入一次（即使被多个包同时引用）

#### 2. 初始化顺序（对每个包）
1. 初始化该包导入的其他包（递归进行）
2. 初始化包级常量
3. 初始化包级变量
4. 执行 init 函数（如果存在）

#### 3. main 包执行
1. 初始化 main 包的包级常量和变量
2. 执行 main 包的 init 函数（如果存在）
3. 执行 main 函数

### 执行流程图

```
导入包 A, B, C
    ↓
初始化包 A (常量 → 变量 → init())
    ↓
初始化包 B (常量 → 变量 → init())
    ↓
初始化包 C (常量 → 变量 → init())
    ↓
初始化 main 包 (常量 → 变量 → init())
    ↓
执行 main()
```

### 关键特性

- **自动调用**：init() 和 main() 函数由 Go 运行时自动调用，无需手动调用
- **可选性**：init 函数是可选的，但 package main 必须包含 main 函数
- **执行保证**：所有依赖包的初始化完成后，才会开始执行 main 包的代码
- init()函数只会初始化1次，有多个的话，按顺序执行

### 示例

```go
package main

import "fmt"

// 包级变量初始化
var globalVar = initGlobalVar()

func initGlobalVar() string {
    fmt.Println("初始化全局变量")
    return "global"
}

// init 函数
func init() {
    fmt.Println("执行 init 函数")
}

// main 函数
func main() {
    fmt.Println("执行 main 函数")
}
```

**输出顺序**：
```
初始化全局变量
执行 init 函数
执行 main 函数
```
---

## Go 语言中的指针：引用传递与地址操作

### 基本概念

- **`&` 取地址符**：获取变量的内存地址
- **`*` 指针符号**：
  - 在类型声明中表示指针类型（如 `*int`）
  - 在变量前表示解引用，获取指针指向的值

---

### 示例 1：基本的指针操作

```go
package main

import "fmt"

func main() {
    // 声明一个普通变量
    num := 42
    
    // 使用 & 获取变量的地址
    ptr := &num
    
    fmt.Println("变量的值:", num)           // 输出: 42
    fmt.Println("变量的地址:", &num)        // 输出: 0xc000012028 (示例地址)
    fmt.Println("指针存储的地址:", ptr)      // 输出: 0xc000012028
    fmt.Println("指针指向的值:", *ptr)      // 输出: 42
    
    // 通过指针修改原变量的值
    *ptr = 100
    fmt.Println("修改后的值:", num)         // 输出: 100
}
```

**输出结果**：
```
变量的值: 42
变量的地址: 0xc000012028
指针存储的地址: 0xc000012028
指针指向的值: 42
修改后的值: 100
```

---

### 示例 2：值传递 vs 引用传递

### 值传递（传递副本）

```go
package main

import "fmt"

// 值传递：函数接收变量的副本
func modifyValue(x int) {
    x = 100
    fmt.Println("函数内部的值:", x)
}

func main() {
    num := 42
    modifyValue(num)
    fmt.Println("函数外部的值:", num)  // 原值不变
}
```

**输出结果**：
```
函数内部的值: 100
函数外部的值: 42
```

### 引用传递（传递指针）

```go
package main

import "fmt"

// 引用传递：函数接收变量的地址
func modifyPointer(x *int) {
    *x = 100
    fmt.Println("函数内部的值:", *x)
}

func main() {
    num := 42
    modifyPointer(&num)  // 传递地址
    fmt.Println("函数外部的值:", num)  // 原值被修改
}
```

**输出结果**：
```
函数内部的值: 100
函数外部的值: 100
```

---

### 示例 3：结构体的指针传递

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// 值传递：修改不影响原结构体
func updatePersonValue(p Person) {
    p.Age = 30
    fmt.Printf("函数内: %+v\n", p)
}

// 引用传递：修改会影响原结构体
func updatePersonPointer(p *Person) {
    p.Age = 30
    fmt.Printf("函数内: %+v\n", *p)
}

func main() {
    person1 := Person{Name: "Alice", Age: 25}
    fmt.Println("=== 值传递 ===")
    fmt.Printf("调用前: %+v\n", person1)
    updatePersonValue(person1)
    fmt.Printf("调用后: %+v\n", person1)
    
    fmt.Println("\n=== 引用传递 ===")
    person2 := Person{Name: "Bob", Age: 25}
    fmt.Printf("调用前: %+v\n", person2)
    updatePersonPointer(&person2)
    fmt.Printf("调用后: %+v\n", person2)
}
```

**输出结果**：
```
=== 值传递 ===
调用前: {Name:Alice Age:25}
函数内: {Name:Alice Age:30}
调用后: {Name:Alice Age:25}

=== 引用传递 ===
调用前: {Name:Bob Age:25}
函数内: {Name:Bob Age:30}
调用后: {Name:Bob Age:30}
```

---

### 示例 4：切片的特殊性

```go
package main

import "fmt"

// 切片本身包含指向底层数组的指针
func modifySlice(s []int) {
    s[0] = 999  // 修改元素会影响原切片
    fmt.Println("函数内:", s)
}

func appendSlice(s []int) {
    s = append(s, 100)  // append 可能创建新数组，不影响原切片
    fmt.Println("函数内追加后:", s)
}

func appendSlicePointer(s *[]int) {
    *s = append(*s, 100)  // 通过指针修改原切片
    fmt.Println("函数内追加后:", *s)
}

func main() {
    nums := []int{1, 2, 3}
    
    fmt.Println("=== 修改切片元素 ===")
    fmt.Println("原切片:", nums)
    modifySlice(nums)
    fmt.Println("修改后:", nums)
    
    fmt.Println("\n=== 追加元素（值传递）===")
    nums2 := []int{1, 2, 3}
    fmt.Println("原切片:", nums2)
    appendSlice(nums2)
    fmt.Println("追加后:", nums2)
    
    fmt.Println("\n=== 追加元素（指针传递）===")
    nums3 := []int{1, 2, 3}
    fmt.Println("原切片:", nums3)
    appendSlicePointer(&nums3)
    fmt.Println("追加后:", nums3)
}
```

**输出结果**：
```
=== 修改切片元素 ===
原切片: [1 2 3]
函数内: [999 2 3]
修改后: [999 2 3]

=== 追加元素（值传递）===
原切片: [1 2 3]
函数内追加后: [1 2 3 100]
追加后: [1 2 3]

=== 追加元素（指针传递）===
原切片: [1 2 3]
函数内追加后: [1 2 3 100]
追加后: [1 2 3 100]
```

---

### 示例 5：指针的零值与判断

```go
package main

import "fmt"

func main() {
    var ptr *int  // 声明指针，未初始化
    
    fmt.Println("未初始化的指针:", ptr)  // 输出: <nil>
    
    // 判断指针是否为空
    if ptr == nil {
        fmt.Println("指针为空，不能解引用")
    }
    
    // 初始化指针
    num := 42
    ptr = &num
    
    if ptr != nil {
        fmt.Println("指针不为空，值为:", *ptr)
    }
}
```

**输出结果**：
```
未初始化的指针: <nil>
指针为空，不能解引用
指针不为空，值为: 42
```

---

### 总结对比

| 特性 | 值传递 | 引用传递（指针） |
|------|--------|------------------|
| 参数类型 | `func(x int)` | `func(x *int)` |
| 传递方式 | 传递变量的副本 | 传递变量的地址 |
| 修改影响 | 不影响原变量 | 影响原变量 |
| 内存开销 | 复制整个值 | 只复制地址（8字节） |
| 使用场景 | 小数据、不需修改 | 大数据、需要修改 |
| 调用方式 | `func(value)` | `func(&value)` |

### 使用建议

1. **小数据类型**（int、bool等）：值传递即可
2. **大结构体**：使用指针传递以提高性能
3. **需要修改原数据**：必须使用指针传递
4. **切片、map、channel**：本身已包含引用，通常不需要显式传递指针
5. **避免空指针**：使用指针前务必检查是否为 nil