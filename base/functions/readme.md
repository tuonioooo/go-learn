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