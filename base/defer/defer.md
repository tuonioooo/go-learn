# Go Defer 完全指南

## 目录
- [什么是 Defer](#什么是-defer)
- [基本用法](#基本用法)
- [执行时机](#执行时机)
- [核心特性](#核心特性)
- [常见应用场景](#常见应用场景)
- [高级用法](#高级用法)
- [常见陷阱](#常见陷阱)
- [性能优化](#性能优化)
- [最佳实践](#最佳实践)

---

## 什么是 Defer

`defer` 是 Go 语言的关键字,用于延迟函数调用的执行。被 defer 修饰的函数会在包含它的函数返回之前执行,无论函数是正常返回还是发生 panic。

**核心概念:**
- defer 语句会将函数调用压入一个栈中
- 当外层函数返回时,这些延迟函数按照后进先出(LIFO)的顺序执行
- defer 常用于资源清理、错误处理和性能监控等场景

---

## 基本用法

### 简单示例

```go
package main

import "fmt"

func main() {
    fmt.Println("开始")
    defer fmt.Println("结束")
    fmt.Println("中间")
}

// 输出:
// 开始
// 中间
// 结束
```

### 多个 Defer 语句

```go
func example() {
    defer fmt.Println("第一个 defer")
    defer fmt.Println("第二个 defer")
    defer fmt.Println("第三个 defer")
    fmt.Println("函数体")
}

// 输出:
// 函数体
// 第三个 defer
// 第二个 defer
// 第一个 defer
```

---

## 执行时机

### 1. 正常返回时执行

```go
func normalReturn() {
    defer fmt.Println("defer 执行")
    fmt.Println("正常执行")
    return
}
```

### 2. Panic 时也会执行

```go
func panicCase() {
    defer fmt.Println("defer 仍然会执行")
    fmt.Println("即将 panic")
    panic("发生错误")
    fmt.Println("这行不会执行")
}

// 输出:
// 即将 panic
// defer 仍然会执行
// panic: 发生错误
```

### 3. 作用域规则

defer 只在声明它的函数内生效:

```go
func outer() {
    defer fmt.Println("outer defer")
    inner()
    fmt.Println("outer 继续执行")
}

func inner() {
    defer fmt.Println("inner defer")
    fmt.Println("inner 执行")
}

// 输出:
// inner 执行
// inner defer
// outer 继续执行
// outer defer
```

---

## 核心特性

### 特性 1: 参数立即求值

defer 语句的参数在声明时就已经求值,而非执行时:

```go
func main() {
    x := 10
    defer fmt.Println("Deferred:", x)  // x 在这里就被求值为 10
    x = 20
    fmt.Println("Current:", x)
}

// 输出:
// Current: 20
// Deferred: 10
```

### 特性 2: 可以修改命名返回值

```go
func compute() (result int) {
    defer func() {
        result *= 2  // 修改返回值
    }()
    result = 5
    return  // 返回前会执行 defer,result 变成 10
}

func main() {
    fmt.Println("Result:", compute())  // 输出: Result: 10
}
```

### 特性 3: LIFO 执行顺序

```go
func stack() {
    for i := 0; i < 3; i++ {
        defer fmt.Println(i)
    }
}

// 输出:
// 2
// 1
// 0
```

---

## 常见应用场景

### 1. 文件资源管理

```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // 确保文件一定会被关闭
    
    // 读取文件内容
    data := make([]byte, 100)
    _, err = file.Read(data)
    return err
}
```

### 2. 互斥锁管理

```go
var mu sync.Mutex

func updateSharedResource() {
    mu.Lock()
    defer mu.Unlock()  // 防止死锁
    
    // 临界区代码
    fmt.Println("更新共享资源")
}
```

### 3. 数据库连接管理

```go
func queryDatabase() error {
    db, err := sql.Open("mysql", "dsn")
    if err != nil {
        return err
    }
    defer db.Close()
    
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        return err
    }
    defer rows.Close()
    
    // 处理查询结果
    return nil
}
```

### 4. HTTP 响应体关闭

```go
func fetchURL(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    // 处理响应
    body, err := io.ReadAll(resp.Body)
    fmt.Println(string(body))
    return err
}
```

### 5. 函数执行时间监控

```go
func measureTime() {
    start := time.Now()
    defer func() {
        elapsed := time.Since(start)
        fmt.Printf("函数执行时间: %v\n", elapsed)
    }()
    
    // 模拟耗时操作
    time.Sleep(2 * time.Second)
}
```

### 6. Panic 恢复

```go
func safeDivide(a, b int) (result int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("捕获到 panic:", r)
            result = 0
        }
    }()
    
    result = a / b  // 如果 b 为 0 会 panic
    return
}
```

---

## 高级用法

### 1. 利用闭包修改返回值

```go
func increment() (i int) {
    defer func() {
        i++
    }()
    return 1  // 实际返回 2
}
```

### 2. 链式 Defer 用于追踪

```go
func trace(name string) func() {
    fmt.Println("进入:", name)
    return func() {
        fmt.Println("离开:", name)
    }
}

func a() {
    defer trace("a")()
    fmt.Println("执行 a")
    b()
}

func b() {
    defer trace("b")()
    fmt.Println("执行 b")
}
```

### 3. 条件性 Defer

```go
func conditionalDefer(needCleanup bool) {
    var cleanup func()
    
    if needCleanup {
        cleanup = func() {
            fmt.Println("执行清理")
        }
        defer cleanup()
    }
    
    // 业务逻辑
}
```

---

## 常见陷阱

### 陷阱 1: 在循环中使用 Defer

**❌ 错误示例:**

```go
func processFiles() {
    for i := 0; i < 100; i++ {
        file, _ := os.Open(fmt.Sprintf("file%d.txt", i))
        defer file.Close()  // 所有文件会在函数结束时才关闭,可能导致资源耗尽
    }
}
```

**✅ 正确做法:**

```go
func processFiles() {
    for i := 0; i < 100; i++ {
        func() {
            file, _ := os.Open(fmt.Sprintf("file%d.txt", i))
            defer file.Close()  // 在匿名函数返回时立即关闭
            // 处理文件
        }()
    }
}
```

### 陷阱 2: 忽略 Defer 函数的错误

**❌ 错误示例:**

```go
func writeFile() {
    file, _ := os.Create("test.txt")
    defer file.Close()  // 忽略了 Close 可能返回的错误
    
    file.WriteString("hello")
}
```

**✅ 正确做法:**

```go
func writeFile() (err error) {
    file, err := os.Create("test.txt")
    if err != nil {
        return err
    }
    
    defer func() {
        closeErr := file.Close()
        if err == nil {
            err = closeErr
        }
    }()
    
    _, err = file.WriteString("hello")
    return err
}
```

### 陷阱 3: Defer 与 nil 指针

**❌ 错误示例:**

```go
func riskyCode() {
    var file *os.File
    defer file.Close()  // file 是 nil,会导致 panic
    
    var err error
    file, err = os.Open("test.txt")
    if err != nil {
        return
    }
}
```

**✅ 正确做法:**

```go
func safeCode() error {
    file, err := os.Open("test.txt")
    if err != nil {
        return err
    }
    defer file.Close()
    
    // 处理文件
    return nil
}
```

### 陷阱 4: 闭包变量捕获问题

**❌ 错误示例:**

```go
func printNumbers() {
    for i := 0; i < 3; i++ {
        defer func() {
            fmt.Println(i)  // 都会打印 3
        }()
    }
}
```

**✅ 正确做法:**

```go
func printNumbers() {
    for i := 0; i < 3; i++ {
        defer func(n int) {
            fmt.Println(n)  // 正确打印 2, 1, 0
        }(i)
    }
}
```

---

## 性能优化

### Defer 的性能开销

Go 1.13 之前,defer 有较大的性能开销。Go 1.13+ 引入了开放编码优化,大幅降低了开销。

### 三种 Defer 实现方式

1. **Open-coded defer** (最快): 编译器直接内联
2. **Stack-allocated defer** (中等): 在栈上分配
3. **Heap-allocated defer** (最慢): 在堆上分配

### 性能对比示例

```go
// 不使用 defer
func withoutDefer() {
    mu.Lock()
    // 操作
    mu.Unlock()
}

// 使用 defer (Go 1.13+ 性能几乎无差异)
func withDefer() {
    mu.Lock()
    defer mu.Unlock()
    // 操作
}
```

### 优化建议

- 在性能关键路径上,如果 defer 带来明显开销,考虑手动管理
- 对于绝大多数场景,defer 的可维护性优势远大于微小的性能损失
- 使用基准测试验证性能影响

---

## 最佳实践

### 1. 紧邻资源获取处使用 Defer

```go
func goodPractice() error {
    file, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer file.Close()  // 立即 defer,代码更清晰
    
    // 使用 file
    return nil
}
```

### 2. 使用命名返回值处理错误

```go
func robustFunction() (err error) {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    
    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            err = fmt.Errorf("panic: %v", p)
        } else if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()
    
    // 业务逻辑
    return nil
}
```

### 3. 避免过度使用

```go
// ❌ 不必要的 defer
func unnecessary() {
    defer fmt.Println("end")
    fmt.Println("start")
    // 只有两行代码,不需要 defer
}

// ✅ 合理使用
func necessary() error {
    file, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer file.Close()
    
    // 复杂的处理逻辑
    // 多个可能的返回路径
    return nil
}
```

### 4. 使用 Defer 简化多返回路径

```go
func complexLogic(id int) error {
    resource := acquireResource(id)
    defer resource.Release()
    
    if condition1 {
        return errors.New("error 1")
    }
    
    if condition2 {
        return errors.New("error 2")
    }
    
    if condition3 {
        return errors.New("error 3")
    }
    
    return nil
    // 无论从哪里返回,资源都会被释放
}
```

### 5. WaitGroup 配合 Defer

```go
func concurrentTasks() {
    var wg sync.WaitGroup
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()  // 确保一定会调用 Done
            
            // 任务逻辑
            processTask(id)
        }(i)
    }
    
    wg.Wait()
}
```

---

## 总结

### Q: defer 什么时候执行？

A: 在函数 return 的那一刻，就在函数真正退出之前

### Q: 多个 defer 怎么排序？

A: 倒着来！最后写的 defer 最先执行

### Q: defer 能捕获错误吗？

A: 可以配合 recover() 捕获 panic
```go
func 安全执行() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("捕获到错误:", err)
        }
    }()
    
    panic("出错了")  // 会被上面的 defer 捕获
}
```

📌 记住这个口诀

> "defer 是延迟执行，return 前必执行"
> "先进后出像堆栈，清理资源不会忘"
---

## 参考资源

- [Go 官方博客: Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
- [Effective Go: Defer](https://go.dev/doc/effective_go#defer)
- Go 1.13+ 版本的 defer 性能优化文档

---

*最后更新: 2024*