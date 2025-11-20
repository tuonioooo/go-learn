# Go 并发（Concurrency）完整教程

## 目录
1. [并发基础概念](#1-并发基础概念)
2. [Goroutine - 轻量级线程](#2-goroutine---轻量级线程)
3. [Channel - 通信机制](#3-channel---通信机制)
4. [代码详解：CheckWebsites](#4-代码详解checkwebsites)
5. [并发模式](#5-并发模式)
6. [最佳实践](#6-最佳实践)

---

## 1. 并发基础概念

### 1.1 并发 vs 并行

```
并发（Concurrency）：多个任务在交替执行
┌─────────────┬─────────────┬─────────────┐
│   Task A    │   Task B    │   Task A    │
└─────────────┴─────────────┴─────────────┘
      时间

并行（Parallelism）：多个任务同时执行（多核）
┌─────────────────────────────────────────┐
│   Task A                                 │
├─────────────┬─────────────┬─────────────┤
│   Task B    │   Task C    │   Task D    │
└─────────────┴─────────────┴─────────────┘
      时间
```

**关键区别**：
- **并发**：一个处理器上，时间片轮转
- **并行**：多个处理器上，真正同时运行

### 1.2 为什么需要并发？

```go
// ❌ 串行执行 - 太慢
func CheckWebsitesSerial(urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = checkURL(url)  // 一个接一个，依次等待
	}
	return results
}
// 如果每个请求耗时 1 秒，10 个 URL 就需要 10 秒

// ✅ 并发执行 - 快速
func CheckWebsitesConcurrent(urls []string) map[string]bool {
	// 所有 URL 并发检查，总耗时约 1 秒（最慢的那个）
	// ...
}
```

### 1.3 Go 的并发模型

Go 使用**共享内存通过通信**（Share Memory By Communicating）的理念：
- 不共享内存来通信
- 通过通信来共享内存

---

## 2. Goroutine - 轻量级线程

### 2.1 什么是 Goroutine？

Goroutine 是由 Go 运行时管理的轻量级线程。

```go
// 普通函数
func sayHello(name string) {
	fmt.Println("Hello, " + name)
}

// 在 Goroutine 中执行
go sayHello("Alice")  // 非阻塞地启动一个 Goroutine
```

### 2.2 Goroutine vs 线程

| 特性 | Goroutine | 线程 |
|------|----------|------|
| 创建成本 | 非常低（微秒） | 高（毫秒） |
| 内存消耗 | ~2KB | ~1-2MB |
| 数量 | 可创建成千上万 | 通常只有几十个 |
| 调度 | Go 运行时 | 操作系统 |
| 通信 | Channel | 共享内存+锁 |

### 2.3 创建和管理 Goroutine

```go
// 方式1：执行函数
go func1()

// 方式2：执行方法
go receiver.method()

// 方式3：执行匿名函数
go func() {
	fmt.Println("Anonymous goroutine")
}()

// 方式4：执行带参数的匿名函数
go func(name string) {
	fmt.Println("Hello, " + name)
}("Alice")

// 方式5：循环中启动 Goroutine
for _, item := range items {
	go processItem(item)
}
```

### 2.4 等待 Goroutine 完成

```go
import "sync"

// 方式1：使用 WaitGroup（推荐）
var wg sync.WaitGroup

wg.Add(1)
go func() {
	defer wg.Done()
	fmt.Println("Task 1")
}()

wg.Add(1)
go func() {
	defer wg.Done()
	fmt.Println("Task 2")
}()

wg.Wait()  // 等待所有 Goroutine 完成
fmt.Println("All tasks done")

// 方式2：使用 Channel（见下节）
resultChannel := make(chan string)

go func() {
	resultChannel <- "Task 1 done"
}()

result := <-resultChannel  // 等待结果
```

---

## 3. Channel - 通信机制

### 3.1 什么是 Channel？

Channel 是 Goroutine 之间的通信管道。使用 `chan` 关键字定义 Channel 类型。

```go
// 创建 Channel
resultChannel := make(chan string)

// 发送数据
resultChannel <- "Hello"

// 接收数据
message := <-resultChannel
```

### 3.1.1 `chan` 关键字详解

```go
chan result
 ↑    ↑
 |    └── Channel 传输的数据类型
 └────── Channel 关键字

// 完整含义：创建一个可以传输 result 类型数据的 Channel
```

**Channel 类型语法**：
```go
make(chan DataType)
make(chan DataType, BufferSize)
```

| 创建方式 | 说明 | 例子 |
|---------|------|------|
| `chan int` | 双向 Channel，传输 int | `ch := make(chan int)` |
| `chan string` | 双向 Channel，传输 string | `ch := make(chan string)` |
| `chan result` | 双向 Channel，传输 result | `ch := make(chan result)` |
| `<-chan int` | 只读 Channel | `readCh := make(<-chan int)` |
| `chan<- int` | 只写 Channel | `writeCh := make(chan<- int)` |
| `chan int, 5` | 有缓冲 Channel（容量 5） | `ch := make(chan int, 5)` |

**Channel 的三个组成部分**：
1. `chan` - 关键字，声明这是一个 Channel
2. 数据类型 - 这个 Channel 能传输什么类型的数据
3. 缓冲大小（可选）- 无缓冲（默认）或有缓冲

### 3.2 Channel 的类型

```go
// 无缓冲 Channel（阻塞式）
ch1 := make(chan int)

// 有缓冲 Channel（非阻塞式，容量为 5）
ch2 := make(chan int, 5)

// 只读 Channel
readOnlyCh := make(<-chan int)

// 只写 Channel
writeOnlyCh := make(chan<- int)
```

**Channel 类型对比表**：

| Channel 类型 | 可以发送 | 可以接收 | 缓冲 | 使用场景 |
|-------------|---------|---------|------|---------|
| `chan T` | ✅ 是 | ✅ 是 | ✅ 无 | 一般通信 |
| `chan T, n` | ✅ 是 | ✅ 是 | ✅ 有 | 减少 Goroutine 阻塞 |
| `<-chan T` | ❌ 否 | ✅ 是 | - | 只读接口 |
| `chan<- T` | ✅ 是 | ❌ 否 | - | 只写接口 |

### 3.3 无缓冲 vs 有缓冲

```go
// 无缓冲 Channel - 发送者和接收者必须同时准备好
ch := make(chan int)

go func() {
	ch <- 5  // 发送，等待接收者准备好
}()

value := <-ch  // 接收
fmt.Println(value)  // 5

// 有缓冲 Channel - 可以先发送，后接收
ch := make(chan int, 2)
ch <- 5   // 发送（不阻塞）
ch <- 10  // 发送（不阻塞）
fmt.Println(<-ch)  // 5
fmt.Println(<-ch)  // 10
```

**无缓冲 Channel 的阻塞行为**：

上面的代码**会有阻塞，但不会死锁**。执行流程如下：

```
时间轴：
T0  ├─ 执行 ch := make(chan int)，创建无缓冲 Channel
    ├─ 执行 go func() {...}()，启动 Goroutine
    └─ 继续执行主 Goroutine

T0+  ├─ 子 Goroutine 执行到 ch <- 5
    │  └─ ⏸️ 【阻塞】：尝试发送，但没有接收者准备好
    │
    └─ 主 Goroutine 继续执行，到达 value := <-ch
       └─ ✅ 成为接收者，子 Goroutine 的发送解除阻塞

T0++ ├─ Channel 完成数据传输
     ├─ fmt.Println(value)  输出 5
     └─ 程序结束
```

**关键理解**：

| 阶段 | Goroutine | 状态 | 说明 |
|------|-----------|------|------|
| T0 | 子 Goroutine | 运行中 | 执行到 ch <- 5 |
| T0+ | 子 Goroutine | **阻塞** | 等待接收者（无接收者） |
| T0+ | 主 Goroutine | 运行中 | 继续执行到 <-ch |
| T0++ | 子 Goroutine | **继续** | 发送完成，解除阻塞 |
| T0++ | 主 Goroutine | **继续** | 接收完成，获得值 5 |

**为什么不会死锁**：

```go
// ❌ 会死锁：主 Goroutine 在发送处阻塞，没有其他 Goroutine 接收
ch := make(chan int)
ch <- 5  // 永久阻塞！程序 panic: all goroutines are asleep - deadlock!

// ❌ 会死锁：主 Goroutine 是唯一的，尝试从空 Channel 接收
ch := make(chan int)
value := <-ch  // 永久阻塞！没有 Goroutine 来发送

// ✅ 你的代码：不会死锁
ch := make(chan int)
go func() {
	ch <- 5  // 这里阻塞是【临时的】
}()
value := <-ch  // 主 Goroutine 继续运行，成为接收者
```

**无缓冲 Channel 的同步特性**：

无缓冲 Channel 通过**阻塞**实现了 Goroutine 之间的**同步**：
- 发送方必须等到接收方准备好
- 接收方必须等到发送方准备好
- 这种"握手"行为确保了数据的可靠传输

这就是 CheckWebsites 中使用无缓冲 Channel 来等待所有结果的原因！

### 3.4 关闭 Channel

```go
ch := make(chan int)

go func() {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)  // 关闭 Channel
}()

for value := range ch {
	fmt.Println(value)  // 1, 2, 3
}

// 尝试从关闭的 Channel 接收
value, ok := <-ch
fmt.Println(value, ok)  // 0, false
```

### 3.5 Channel 操作规则

| 操作 | 无缓冲 | 有缓冲满 | 有缓冲空 | 已关闭 |
|------|-------|---------|---------|--------|
| 发送 | 阻塞直到有接收者 | 阻塞 | 发送 | ❌ Panic |
| 接收 | 阻塞直到有发送者 | 接收 | 阻塞 | 返回零值 |
| 关闭 | 关闭 | 关闭 | 关闭 | ❌ Panic |

---

## 4. 代码详解：CheckWebsites

### 4.1 完整代码

```go
package concurrency

// WebsiteChecker 是一个函数类型，用于检查网站是否正常
type WebsiteChecker func(string) bool

// result 结构体存储检查结果
type result struct {
	string  // 网址
	bool    // 是否正常
}

// CheckWebsites 并发检查多个网站
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	// 1. 创建结果 map，存储最终结果
	results := make(map[string]bool)
	
	// 2. 创建无缓冲 Channel，用于接收检查结果
	// chan result 表示：
	//   - chan：这是一个 Channel
	//   - result：这个 Channel 只能传输 result 类型的数据
	//   - make(chan result)：创建一个无缓冲的 result Channel
	resultChannel := make(chan result)

	// 3. 为每个 URL 启动一个 Goroutine
	for _, url := range urls {
		// ❌ 常见错误：直接使用 url
		// go func() {
		//     resultChannel <- result{url, wc(url)}
		// }()
		// 问题：闭包捕获的 url 会改变，所有 Goroutine 最后都用最后一个 URL

		// ✅ 正确：将 url 作为参数传递
		go func(u string) {
			// 创建 result 值并发送到 Channel
			resultChannel <- result{u, wc(u)}
			// result{u, wc(u)} 是 result 类型的值
			// <- 是发送操作符，将这个值发送到 resultChannel
		}(url)
	}

	// 4. 从 Channel 接收所有结果
	for i := 0; i < len(urls); i++ {
		// result := <-resultChannel 从 Channel 接收一个 result 值
		// <- 是接收操作符
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
```

### 4.1.1 Channel 操作详解

```go
// ========== 发送操作 ==========
resultChannel <- result{url, true}
//            ↑
//         发送操作符
// 将 result 值发送到 resultChannel

// ========== 接收操作 ==========
r := <-resultChannel
//   ↑
// 接收操作符
// 从 resultChannel 接收一个 result 值，赋值给 r

// ========== Channel 类型 ==========
resultChannel := make(chan result)
//                        ↑
//                  数据类型
// 这个 Channel 只能传输 result 类型的数据

// ========== 发送/接收必须匹配类型 ==========
intCh := make(chan int)
intCh <- 5                    // ✅ 正确，5 是 int
// intCh <- "hello"           // ❌ 错误，"hello" 是 string，不匹配

resultCh := make(chan result)
resultCh <- result{"url", true}  // ✅ 正确，是 result 类型
// resultCh <- "url"          // ❌ 错误，"url" 是 string，不匹配 chan result
```

### 4.2 代码流程图

```
CheckWebsites(wc, ["url1", "url2", "url3"])
│
├─► 创建 results map 和 resultChannel
│
├─► 启动 Goroutine 1: go func(u="url1") { resultChannel <- result{"url1", wc("url1")} }(url)
├─► 启动 Goroutine 2: go func(u="url2") { resultChannel <- result{"url2", wc("url2")} }(url)
├─► 启动 Goroutine 3: go func(u="url3") { resultChannel <- result{"url3", wc("url3")} }(url)
│
├─► 主 Goroutine 等待接收结果
│   result := <-resultChannel  // 接收 Goroutine 1 的结果
│   result := <-resultChannel  // 接收 Goroutine 2 的结果
│   result := <-resultChannel  // 接收 Goroutine 3 的结果
│
└─► 返回 results map
```

### 4.3 关键点详解

#### 4.3.1 为什么要传递参数？

```go
// ❌ 错误做法：闭包陷阱
for _, url := range urls {
	go func() {
		// url 是对循环变量的引用
		// 当 Goroutine 执行时，url 可能已经改变
		resultChannel <- result{url, wc(url)}
	}()
}
// 所有 Goroutine 都会使用最后一个 url！

// ✅ 正确做法：参数传递
for _, url := range urls {
	go func(u string) {
		// u 是参数，每个 Goroutine 有自己的副本
		resultChannel <- result{u, wc(u)}
	}(url)  // 将 url 作为参数传递
}
```

#### 4.3.2 无缓冲 Channel 的作用

```go
resultChannel := make(chan result)  // 无缓冲

// 第一个 Goroutine 发送结果
// 如果没有接收者，会阻塞

// 主 Goroutine 接收
for i := 0; i < len(urls); i++ {
	result := <-resultChannel
	// 接收到数据后，之前阻塞的 Goroutine 继续执行
	results[result.string] = result.bool
}
```

#### 4.3.3 struct 的简写

```go
type result struct {
	string  // 嵌入字段（匿名字段）
	bool
}

// 创建 result
r := result{"http://example.com", true}

// 访问字段（两种方式等价）
r.string  // "http://example.com"
r.bool    // true
```

### 4.4 执行流程时间图

```
时间轴：
T0  ├─ 启动 Goroutine 1（检查 url1，耗时 1s）
    ├─ 启动 Goroutine 2（检查 url2，耗时 2s）
    ├─ 启动 Goroutine 3（检查 url3，耗时 1.5s）
    └─ 主 Goroutine 阻塞在 <-resultChannel
    
T1s ├─ Goroutine 1 完成 ──► 发送结果 ──► 主 Goroutine 接收
    └─ 主 Goroutine 接收第一个结果，继续等待

T1.5s ├─ Goroutine 3 完成 ──► 发送结果 ──► 主 Goroutine 接收
      └─ 主 Goroutine 接收第二个结果，继续等待

T2s  ├─ Goroutine 2 完成 ──► 发送结果 ──► 主 Goroutine 接收
     └─ 主 Goroutine 接收第三个结果，返回

总耗时：2s（最慢的 Goroutine 的耗时）
如果是串行：1s + 2s + 1.5s = 4.5s
```

---

## 5. 并发模式

### 5.1 Fan-Out/Fan-In 模式

```go
// Fan-Out：一个任务分发给多个 Goroutine
// Fan-In：多个 Goroutine 的结果汇聚到一个 Channel

func processParallel(items []string) []string {
	// 创建 Channel
	resultCh := make(chan string)
	
	// Fan-Out：启动多个 Goroutine
	for _, item := range items {
		go func(i string) {
			result := processItem(i)
			resultCh <- result  // 发送结果
		}(item)
	}
	
	// Fan-In：收集所有结果
	var results []string
	for i := 0; i < len(items); i++ {
		results = append(results, <-resultCh)
	}
	
	return results
}
```

### 5.2 Worker 池模式

```go
func workerPool(jobs <-chan int, numWorkers int) {
	// 创建 results Channel
	results := make(chan int)
	
	// 启动固定数量的 Worker
	for i := 0; i < numWorkers; i++ {
		go worker(jobs, results)
	}
	
	// 处理结果
	for j := 0; j < 10; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * 2
	}
}
```

### 5.3 Timeout 模式

```go
import "time"

ch := make(chan string)

go func() {
	time.Sleep(2 * time.Second)
	ch <- "Result"
}()

// 设置 5 秒超时
select {
case result := <-ch:
	fmt.Println("Got result:", result)
case <-time.After(5 * time.Second):
	fmt.Println("Timeout!")
}
```

---

## 6. 最佳实践

### 6.1 Channel 的最佳实践

```go
// ✅ 好的做法：及时关闭 Channel
func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)  // 发送者负责关闭
}

func main() {
	ch := make(chan int)
	go producer(ch)
	
	for value := range ch {
		fmt.Println(value)
	}
}

// ❌ 避免：多个发送者共享 Channel
// 因为发送者无法确定何时关闭
```

### 6.2 Goroutine 的最佳实践

```go
// ✅ 好的做法：使用 WaitGroup 管理 Goroutine
import "sync"

var wg sync.WaitGroup

for _, item := range items {
	wg.Add(1)
	go func(i string) {
		defer wg.Done()
		process(i)
	}(item)
}

wg.Wait()

// ❌ 避免：启动过多 Goroutine
for i := 0; i < 1000000; i++ {
	go doSomething()  // 会耗尽系统资源
}
```

### 6.3 死锁避免

```go
// ❌ 死锁：发送到无缓冲 Channel 但没有接收者
ch := make(chan int)
ch <- 5  // Deadlock!

// ✅ 正确：有接收者
ch := make(chan int)
go func() {
	fmt.Println(<-ch)
}()
ch <- 5

// ✅ 或使用有缓冲 Channel
ch := make(chan int, 1)
ch <- 5
fmt.Println(<-ch)
```

### 6.4 竞态条件避免

```go
// ❌ 有竞态条件：多个 Goroutine 访问共享变量
var count int
for i := 0; i < 10; i++ {
	go func() {
		count++  // 竞态条件！
	}()
}

// ✅ 使用 Channel 避免竞态条件
countCh := make(chan int)
for i := 0; i < 10; i++ {
	go func(i int) {
		countCh <- i
	}(i)
}

total := 0
for i := 0; i < 10; i++ {
	total += <-countCh
}

// ✅ 或使用 Mutex
import "sync"

var mu sync.Mutex
var count int

for i := 0; i < 10; i++ {
	go func() {
		mu.Lock()
		defer mu.Unlock()
		count++
	}()
}
```

---

## 7. CheckWebsites 的改进版本

### 7.1 添加超时支持

```go
func CheckWebsitesWithTimeout(wc WebsiteChecker, urls []string, timeout time.Duration) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	
	for _, url := range urls {
		go func(u string) {
			// 在 Goroutine 中添加超时
			done := make(chan bool, 1)
			go func() {
				done <- wc(u)
			}()
			
			select {
			case isUp := <-done:
				resultChannel <- result{u, isUp}
			case <-time.After(timeout):
				resultChannel <- result{u, false}
			}
		}(url)
	}
	
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	
	return results
}
```

### 7.2 使用 WaitGroup 改进

```go
import "sync"

func CheckWebsitesWithWaitGroup(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	
	var wg sync.WaitGroup
	
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	
	// 等待所有 Goroutine 完成
	go func() {
		wg.Wait()
		close(resultChannel)
	}()
	
	for result := range resultChannel {
		results[result.string] = result.bool
	}
	
	return results
}
```

---

## 总结

| 概念 | 说明 |
|------|------|
| **Goroutine** | 轻量级线程，由 Go 运行时管理 |
| **Channel** | Goroutine 之间的通信管道 |
| **并发** | 多个任务交替执行 |
| **并行** | 多个任务同时执行 |
| **Fan-Out/In** | 一个任务分发给多个 Goroutine，结果汇聚 |
| **无缓冲 Channel** | 发送者和接收者必须同时准备好 |
| **有缓冲 Channel** | 可以先发送，后接收（容量限制） |
| **闭包陷阱** | 循环中启动 Goroutine 要传递参数 |

**关键代码模式**：
```go
resultChannel := make(chan result)

// 启动多个 Goroutine
for _, item := range items {
	go func(i string) {
		resultChannel <- processItem(i)
	}(item)
}

// 收集所有结果
for i := 0; i < len(items); i++ {
	result := <-resultChannel
	// 处理结果
}
```

这就是 Go 高效并发编程的核心！👍
