# Go HTTP 服务器完整教程

## 目录
1. [HTTP 基础概念](#1-http-基础概念)
2. [核心工具详解](#2-核心工具详解)
3. [代码详解](#3-代码详解)
4. [完整工作流](#4-完整工作流)
5. [实践：扩展功能](#5-实践扩展功能)

---

## 1. HTTP 基础概念

### 1.1 什么是 HTTP 服务器？

HTTP 服务器就是一个**监听网络请求，然后返回响应**的程序。

```
客户端 (Client)          网络          服务器 (Server)
   │                     │                  │
   │   HTTP 请求    →    │    ─────────→  │
   │  (GET /player)      │                 │
   │                     │              处理请求
   │                     │            返回响应
   │   HTTP 响应    ←    │    ←─────────  │
   │  (200 OK, "20")     │                 │
   │                     │                  │
```

### 1.2 HTTP 请求和响应

```
HTTP 请求格式：
┌─────────────────────────────────┐
│ GET /player HTTP/1.1            │  ← 请求行：方法、路径、版本
├─────────────────────────────────┤
│ Host: localhost:5000            │  ← 请求头
│ User-Agent: Mozilla/5.0         │
├─────────────────────────────────┤
│                                 │  ← 请求体（可能为空）
└─────────────────────────────────┘

HTTP 响应格式：
┌─────────────────────────────────┐
│ HTTP/1.1 200 OK                 │  ← 状态行：版本、状态码、状态消息
├─────────────────────────────────┤
│ Content-Type: text/plain        │  ← 响应头
│ Content-Length: 2               │
├─────────────────────────────────┤
│ 20                              │  ← 响应体
└─────────────────────────────────┘
```

---

## 2. 核心工具详解

### 2.1 http.ResponseWriter

**作用**：写入 HTTP 响应内容

```go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "20")  // ← 使用 ResponseWriter 写入响应
}
```

**关键方法**：

| 方法 | 作用 | 例子 |
|------|------|------|
| `w.Write()` | 写入响应体（字节） | `w.Write([]byte("Hello"))` |
| `fmt.Fprint(w, ...)` | 写入响应体（格式化） | `fmt.Fprint(w, "20")` |
| `w.Header()` | 设置响应头 | `w.Header().Set("Content-Type", "text/plain")` |
| `w.WriteHeader()` | 设置 HTTP 状态码 | `w.WriteHeader(http.StatusOK)` |

**示例**：

```go
// ✅ 设置状态码和响应头
func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, `{"score": 20}`)
}

// 返回响应：
// HTTP/1.1 200 OK
// Content-Type: application/json
//
// {"score": 20}
```

### 2.2 *http.Request

**作用**：表示客户端发送的 HTTP 请求

```go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    // r 就是 HTTP 请求对象
    fmt.Println(r.Method)   // GET、POST、PUT 等
    fmt.Println(r.URL.Path) // 请求路径 (/player)
    fmt.Println(r.Header)   // 请求头
}
```

**关键字段和方法**：

| 字段/方法 | 作用 | 例子 |
|----------|------|------|
| `r.Method` | HTTP 方法 | "GET", "POST", "PUT", "DELETE" |
| `r.URL.Path` | 请求路径 | "/player", "/scores" |
| `r.Header` | 请求头 | `r.Header.Get("Content-Type")` |
| `r.Body` | 请求体 | 读取 POST 数据 |
| `r.FormValue()` | 获取表单参数 | `r.FormValue("name")` |

**示例**：

```go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        // 处理 GET 请求
        path := r.URL.Path  // 获取路径
        if path == "/player" {
            fmt.Fprint(w, "20")
        }
    } else if r.Method == http.MethodPost {
        // 处理 POST 请求
        fmt.Fprint(w, "Created")
    }
}
```

### 2.3 http.HandlerFunc

**作用**：将函数转换为 HTTP 处理器

```go
// 函数签名：func(http.ResponseWriter, *http.Request)
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "20")
}

// 转换为处理器
handler := http.HandlerFunc(PlayerServer)

// 等价于：
handler := http.Handler(&PlayerServer)
```

**为什么需要 HandlerFunc？**

```go
// ❌ 不能直接使用函数
http.ListenAndServe(":5000", PlayerServer)  // 编译错误

// ✅ 转换为处理器
http.ListenAndServe(":5000", http.HandlerFunc(PlayerServer))
```

**接口本质**：

```go
// http.Handler 接口定义
type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

// http.HandlerFunc 类型转换
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)  // 调用函数
}

// 所以 HandlerFunc 是将【函数】适配为【接口】的工具
```

### 2.4 http.ListenAndServe

**作用**：启动 HTTP 服务器，监听并处理请求

```go
func main() {
    handler := http.HandlerFunc(PlayerServer)
    http.ListenAndServe(":5000", handler)
}
```

**工作流程**：

```
http.ListenAndServe(":5000", handler)
    ↓
启动服务器
    ├─ 绑定端口 5000
    ├─ 进入监听循环
    └─ 等待客户端连接
    
客户端请求来临
    ↓
创建 *http.Request 对象
    ↓
调用 handler.ServeHTTP(w, r)
    ↓
执行 PlayerServer(w, r)
    ↓
返回响应给客户端
    ↓
继续等待下一个请求
```

**参数说明**：

```go
http.ListenAndServe(":5000", handler)
                    ↑         ↑
                   地址      处理器
                   
// ":5000" 表示：
// - 监听本机的所有网卡（0.0.0.0）
// - 端口 5000

// handler 是 http.Handler：
// 当有请求时，调用 handler.ServeHTTP(w, r)
```

### 2.5 httptest 包（测试工具）

**作用**：不启动真实服务器，直接测试 HTTP 处理器

#### httptest.NewRecorder

**用途**：记录响应结果

```go
response := httptest.NewRecorder()

PlayerServer(response, request)

// 获取响应内容
response.Code           // HTTP 状态码（200）
response.Body.String()  // 响应体内容（"20"）
response.Header()       // 响应头
```

**工作原理**：

```
真实场景：
客户端 ─HTTP请求─→ 服务器 ─HTTP响应─→ 客户端

测试场景：
测试代码 ─直接调用─→ PlayerServer(responseRecorder, request)
                         ↓
                    写入响应到 responseRecorder
                         ↓
测试代码 ─读取─→ responseRecorder.Body
```

#### http.NewRequest

**用途**：创建模拟的 HTTP 请求

```go
request, _ := http.NewRequest(http.MethodGet, "/", nil)
//                             ↑             ↑    ↑
//                         HTTP 方法      路径  请求体

response := httptest.NewRecorder()

// 直接调用处理器，不需要启动服务器
PlayerServer(response, request)
```

**参数说明**：

```go
http.NewRequest(
    http.MethodGet,    // HTTP 方法：GET、POST 等
    "/",               // 请求路径
    nil                // 请求体：GET 通常为 nil
)
```

---

## 3. 代码详解

### 3.1 server.go - HTTP 处理器

```go
package main

import (
    "fmt"
    "net/http"
)

// PlayerServer 处理所有请求，返回玩家得分
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "20")
}
```

**逐行解析**：

```go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    ↑              ↑  ↑
    函数名         响应  请求
    
    // w：ResponseWriter，用来写入响应
    // r：*Request，表示客户端的请求
}

fmt.Fprint(w, "20")
↑          ↑  ↑
格式化输出  目标 内容

// fmt.Fprint 的作用：
// 1. 获取参数 "20"
// 2. 将其转换为字符串
// 3. 写入到 w（ResponseWriter）
```

**执行流程**：

```
POST /player HTTP/1.1
Host: localhost:5000

    ↓

调用 PlayerServer(responseWriter, request)

    ↓

fmt.Fprint(responseWriter, "20")
    ├─ 将 "20" 写入 responseWriter
    └─ responseWriter 收集响应数据

    ↓

HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8

20
```

### 3.2 server_test.go - 测试代码

```go
func TestGETPlayers(t *testing.T) {
    // 1. 创建模拟请求
    request, _ := http.NewRequest(http.MethodGet, "/", nil)
    
    // 2. 创建响应记录器
    response := httptest.NewRecorder()

    // 3. 直接调用处理器（不启动服务器）
    PlayerServer(response, request)

    // 4. 验证响应内容
    t.Run("returns Pepper's score", func(t *testing.T) {
        got := response.Body.String()
        want := "20"

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })
}
```

**逐步解析**：

```
第 1 步：创建请求
┌────────────────────────────────────────┐
│ request, _ := http.NewRequest(...)     │
├────────────────────────────────────────┤
│ 创建一个模拟的 HTTP 请求               │
│ ├─ Method: GET                         │
│ ├─ Path: /                             │
│ └─ Body: nil                           │
└────────────────────────────────────────┘

第 2 步：创建响应记录器
┌────────────────────────────────────────┐
│ response := httptest.NewRecorder()     │
├────────────────────────────────────────┤
│ 创建一个响应记录器                     │
│ ├─ Code: 0（初始）                    │
│ ├─ Body: []byte{}（初始）             │
│ └─ Header: map{}（初始）              │
└────────────────────────────────────────┘

第 3 步：调用处理器
┌────────────────────────────────────────┐
│ PlayerServer(response, request)        │
├────────────────────────────────────────┤
│ ├─ 执行 fmt.Fprint(response, "20")   │
│ └─ response.Body 现在包含 "20"       │
└────────────────────────────────────────┘

第 4 步：验证结果
┌────────────────────────────────────────┐
│ got := response.Body.String()          │
│ want := "20"                           │
│ if got != want { ... }                 │
├────────────────────────────────────────┤
│ got = "20"  ✅                         │
│ want = "20"                            │
│ 相等，测试通过！                      │
└────────────────────────────────────────┘
```

### 3.3 main.go - 启动服务器

```go
package main

import (
    "log"
    "net/http"
)

func main() {
    // 1. 将函数转换为处理器
    handler := http.HandlerFunc(PlayerServer)
    
    // 2. 启动服务器
    if err := http.ListenAndServe(":5000", handler); err != nil {
        log.Fatalf("could not listen on port 5000 %v", err)
    }
}
```

**逐步解析**：

```
第 1 步：转换为处理器
handler := http.HandlerFunc(PlayerServer)
    └─ PlayerServer 是函数
    └─ HandlerFunc 将其转换为 Handler 接口

第 2 步：启动服务器
http.ListenAndServe(":5000", handler)
    ├─ ":5000" 绑定 5000 端口
    ├─ handler 处理请求
    └─ 阻塞运行，持续监听

第 3 步：错误处理
if err := ... {
    log.Fatalf(...)  // 如果启动失败，打印错误并退出
}
```

---

## 4. 完整工作流

### 4.1 开发工作流

```
1. 编写处理器 (server.go)
   └─ func PlayerServer(w http.ResponseWriter, r *http.Request)

2. 编写测试 (server_test.go)
   └─ 创建请求、调用处理器、验证响应
   └─ go test 运行测试

3. 启动服务器 (main.go)
   └─ http.ListenAndServe(":5000", handler)

4. 测试服务器
   └─ curl http://localhost:5000
```

### 4.2 HTTP 请求处理流程

```
客户端发送请求
    │
    ├─ 请求行：GET /player HTTP/1.1
    ├─ 请求头：Host: localhost:5000
    └─ 请求体：（空）

                ↓

Go HTTP 服务器收到请求
    │
    ├─ 解析请求
    ├─ 创建 *http.Request 对象
    └─ 创建 http.ResponseWriter 对象

                ↓

调用处理器
    │
    └─ handler.ServeHTTP(w, r)
       └─ PlayerServer(w, r)
          └─ fmt.Fprint(w, "20")

                ↓

生成响应
    │
    ├─ 状态码：200
    ├─ 响应头：Content-Type: text/plain
    └─ 响应体：20

                ↓

返回给客户端
    │
    └─ HTTP/1.1 200 OK
       Content-Type: text/plain; charset=utf-8
       
       20
```

### 4.3 测试流程（不启动真实服务器）

```
测试代码
    │
    ├─ 创建模拟请求：http.NewRequest(...)
    ├─ 创建响应记录器：httptest.NewRecorder()
    └─ 直接调用处理器：PlayerServer(recorder, request)

                ↓

处理器执行
    │
    └─ fmt.Fprint(recorder, "20")
       └─ 将 "20" 写入 recorder

                ↓

验证结果
    │
    ├─ recorder.Body.String() = "20"
    ├─ recorder.Code = 200
    └─ 测试通过

优点：
✅ 速度快（不需要网络）
✅ 不需要启动真实服务器
✅ 便于单元测试
```

---

## 5. 实践：扩展功能

### 5.1 处理不同的路径

```go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    
    if path == "/player" {
        fmt.Fprint(w, "20")
    } else if path == "/leaderboard" {
        fmt.Fprint(w, "Alice: 100\nBob: 90")
    } else {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprint(w, "Not Found")
    }
}
```

**测试**：

```go
t.Run("returns Pepper's score", func(t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/player", nil)
    response := httptest.NewRecorder()
    PlayerServer(response, request)
    
    got := response.Body.String()
    if got != "20" {
        t.Errorf("got %q, want %q", got, "20")
    }
})

t.Run("returns leaderboard", func(t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/leaderboard", nil)
    response := httptest.NewRecorder()
    PlayerServer(response, request)
    
    got := response.Body.String()
    if !strings.Contains(got, "Alice") {
        t.Errorf("expected Alice in leaderboard")
    }
})

t.Run("returns 404 for invalid path", func(t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/invalid", nil)
    response := httptest.NewRecorder()
    PlayerServer(response, request)
    
    if response.Code != http.StatusNotFound {
        t.Errorf("got status %d, want %d", response.Code, http.StatusNotFound)
    }
})
```

### 5.2 处理 POST 请求

```go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // 处理 POST：添加得分
        playerName := r.URL.Query().Get("name")
        score := r.URL.Query().Get("score")
        
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprintf(w, `{"player": "%s", "score": "%s"}`, playerName, score)
        return
    }
    
    // GET 请求：返回得分
    fmt.Fprint(w, "20")
}
```

**测试**：

```go
t.Run("records a player score", func(t *testing.T) {
    request, _ := http.NewRequest(
        http.MethodPost,
        "/player?name=Alice&score=100",
        nil,
    )
    response := httptest.NewRecorder()
    PlayerServer(response, request)
    
    if response.Code != http.StatusOK {
        t.Errorf("got status %d, want %d", response.Code, http.StatusOK)
    }
    
    got := response.Body.String()
    want := `{"player": "Alice", "score": "100"}`
    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
})
```

### 5.3 设置响应头

```go
func PlayerServer(w http.ResponseWriter, r *http.Request) {
    // 设置响应头
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("X-Custom-Header", "MyValue")
    
    // 设置状态码（必须在写入响应体前）
    w.WriteHeader(http.StatusOK)
    
    // 写入响应体
    fmt.Fprint(w, `{"score": 20}`)
}
```

**测试**：

```go
t.Run("returns JSON response", func(t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/player", nil)
    response := httptest.NewRecorder()
    PlayerServer(response, request)
    
    // 验证响应头
    got := response.Header().Get("Content-Type")
    want := "application/json"
    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
    
    // 验证状态码
    if response.Code != http.StatusOK {
        t.Errorf("got status %d, want %d", response.Code, http.StatusOK)
    }
})
```

---

## 6. 并发安全与互斥锁

### 6.1 什么是并发问题？

在 Go 中，多个 Goroutine 可能同时访问同一个资源，导致数据竞争（Race Condition）。

**示例问题**：

```go
// 无保护的情况
var store map[string]int
store["Alice"] = 5

// Goroutine A                  Goroutine B
// 1. 读取 5                     1. 读取 5
// 2. 计算 5 + 1 = 6            2. 计算 5 + 1 = 6
// 3. 写入 6                     3. 写入 6
//
// 预期：7 ❌
// 实际：6 ❌ 数据丢失！
```

### 6.2 互斥锁（Mutex）解决方案

**互斥锁**确保同一时间只有一个 Goroutine 能访问共享资源。

```go
import "sync"

type InMemoryPlayerStore struct {
    store map[string]int
    lock  sync.RWMutex  // ← 互斥锁
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
    i.lock.Lock()           // ① 获取锁
    defer i.lock.Unlock()   // ③ 确保释放锁（defer 保证）
    i.store[name]++         // ② 修改数据（原子操作）
}
```

**执行流程**：

```
时间轴：
│
├─ Goroutine A：i.lock.Lock() → 获得锁 ✓
├─ Goroutine B：i.lock.Lock() → 阻塞等待 ⏳
├─ Goroutine A：i.store[name]++ → 执行
├─ Goroutine A：defer 触发，i.lock.Unlock() → 释放锁
├─ Goroutine B：i.lock.Lock() → 获得锁 ✓
├─ Goroutine B：i.store[name]++ → 执行
├─ Goroutine B：defer 触发，i.lock.Unlock() → 释放锁
│
结果：✅ 数据正确！
```

### 6.3 RWMutex（读写锁）

**问题**：普通 Mutex 在读操作时也会阻塞其他读操作，效率低。

**解决**：RWMutex 允许多个 Goroutine 同时读，但写时独占。

```go
type InMemoryPlayerStore struct {
    store map[string]int
    lock  sync.RWMutex  // ← 读写锁
}

// 写操作：需要独占锁
func (i *InMemoryPlayerStore) RecordWin(name string) {
    i.lock.Lock()           // 获取写锁
    defer i.lock.Unlock()
    i.store[name]++
}

// 读操作：可以共享锁
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    i.lock.RLock()           // 获取读锁（可多个 Goroutine 同时持有）
    defer i.lock.RUnlock()
    return i.store[name]
}
```

**对比**：

```
Mutex：
读1 ─┬─ Lock
     ├─ 阻塞
读2 ─┴─ Lock

RWMutex：
读1 ─┬─ RLock ✓
     │
读2 ─┼─ RLock ✓
     │
读3 ─┴─ RLock ✓

并发高效！
```

### 6.4 defer 的重要性

**为什么用 defer 来释放锁？**

```go
// ❌ 不使用 defer - 容易出错
func (i *InMemoryPlayerStore) RecordWin(name string) {
    i.lock.Lock()
    i.store[name]++
    i.lock.Unlock()  // 如果上面 panic，这里永远不会执行！
}

// ✅ 使用 defer - 安全可靠
func (i *InMemoryPlayerStore) RecordWin(name string) {
    i.lock.Lock()
    defer i.lock.Unlock()  // 无论如何都会执行
    i.store[name]++
}
```

**defer 的执行时机**：

```go
defer i.lock.Unlock()  // 注册延迟执行
i.store[name]++        // 正常执行

// 函数返回时
// ↓
// i.lock.Unlock() 自动执行 ✓
```

### 6.5 代码示例详解

```go
// RecordWin 记录玩家胜利（写操作）
func (i *InMemoryPlayerStore) RecordWin(name string) {
    i.lock.Lock()              // ① 获取锁，阻塞其他 Goroutine
    defer i.lock.Unlock()       // ③ 注册解锁操作（保证执行）
    i.store[name]++             // ② 安全地修改数据
}
// 函数返回时，defer 自动调用 Unlock()

// GetPlayerScore 获取玩家得分（读操作）
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    i.lock.RLock()              // ① 获取读锁，允许其他读操作
    defer i.lock.RUnlock()       // ③ 注册解锁操作
    return i.store[name]        // ② 安全地读取数据
}
// 函数返回时，defer 自动调用 RUnlock()
```

**关键要点**：

| 操作 | Mutex | RWMutex | 说明 |
|------|-------|---------|------|
| 写数据 | `Lock()` / `Unlock()` | `Lock()` / `Unlock()` | 独占，其他操作阻塞 |
| 读数据 | `Lock()` / `Unlock()` | `RLock()` / `RUnlock()` | 可并发，效率高 |
| 多个读 | ❌ 阻塞 | ✅ 并发 | RWMutex 优于 Mutex |

---

## 总结

### 核心工具速查表

| 工具 | 作用 | 何时使用 |
|------|------|---------|
| **http.ResponseWriter** | 写入响应 | 在处理器中 |
| **\*http.Request** | 读取请求 | 在处理器中 |
| **http.HandlerFunc** | 转换函数为处理器 | 在 main.go 中 |
| **http.ListenAndServe** | 启动服务器 | 在 main.go 中 |
| **httptest.NewRecorder** | 记录响应 | 在单元测试中 |
| **http.NewRequest** | 创建请求 | 在单元测试中 |

### 关键代码模式

```go
// 处理器模式
func MyHandler(w http.ResponseWriter, r *http.Request) {
    // 读取请求
    path := r.URL.Path
    method := r.Method
    
    // 设置响应
    w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "response body")
}

// 启动服务器模式
func main() {
    handler := http.HandlerFunc(MyHandler)
    http.ListenAndServe(":5000", handler)
}

// 测试模式
func TestMyHandler(t *testing.T) {
    request, _ := http.NewRequest(http.MethodGet, "/", nil)
    response := httptest.NewRecorder()
    MyHandler(response, request)
    
    if response.Body.String() != "expected" {
        t.Fail()
    }
}
```

### 工作流总结

```
编写处理器 → 编写单元测试 → 启动服务器 → 手动测试
    (TDD)      (httptest)      (main)        (curl)
```

这就是 Go HTTP 服务器开发的核心！🚀
