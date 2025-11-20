# JSON 路由处理完整教程

## 目录
1. [核心概念](#1-核心概念)
2. [代码架构](#2-代码架构)
3. [详细讲解](#3-详细讲解)
4. [new() vs 指针赋值](#4-new-vs-指针赋值)
5. [工作流程](#5-工作流程)

---

## 1. 核心概念

### 1.1 什么是路由？

**路由**就是将不同的 HTTP 请求分配到不同的处理函数。

```
客户端请求                路由器              处理函数
   │                      │                    │
GET /league        →    路由器分析       →   leagueHandler()
                         ├─ /league
POST /players/Pepper →   ├─ /players/     →  playersHandler()
```

### 1.2 什么是 JSON？

JSON 是一种数据格式，用来在网络上传输结构化数据。

```
Go 结构体：
type Player struct {
    Name string
    Wins int
}
player := Player{"Alice", 20}

转换为 JSON：
{"Name":"Alice","Wins":20}

发送给客户端 ← JSON 格式
```

---

## 2. 代码架构

### 2.1 核心结构定义

```go
// PlayerStore 接口定义数据存储操作
type PlayerStore interface {
    GetPlayerScore(name string) int    // 获取玩家得分
    RecordWin(name string)             // 记录玩家胜利
    GetLeague() []Player               // 获取排行榜
}

// Player 结构体存储玩家信息
type Player struct {
    Name string  // 玩家名称
    Wins int     // 胜利次数
}

// PlayerServer 是 HTTP 服务器
type PlayerServer struct {
    store   PlayerStore   // 数据存储
    http.Handler          // 路由处理器
}
```

**架构图**：

```
PlayerServer
├─ store (PlayerStore)
│  ├─ GetPlayerScore()
│  ├─ RecordWin()
│  └─ GetLeague()
│
└─ Handler (路由器)
   ├─ /league → leagueHandler()
   └─ /players/ → playersHandler()
```

### 2.2 路由配置

```go
func NewPlayerServer(store PlayerStore) *PlayerServer {
    p := new(PlayerServer)
    
    p.store = store
    
    // 创建路由器
    router := http.NewServeMux()
    
    // 配置路由
    router.Handle("/league", http.HandlerFunc(p.leagueHandler))
    router.Handle("/players/", http.HandlerFunc(p.playersHandler))
    
    // 将路由器赋值给 PlayerServer
    p.Handler = router
    
    return p
}
```

**配置流程**：

```
1. 创建 PlayerServer 实例
   ↓
2. 保存 store（数据存储）
   ↓
3. 创建路由器（http.ServeMux）
   ↓
4. 注册路由：
   - /league → leagueHandler
   - /players/ → playersHandler
   ↓
5. 将路由器赋值给 p.Handler
   ↓
6. 返回配置完成的 PlayerServer
```

---

## 3. 详细讲解

### 3.1 leagueHandler - 获取排行榜（JSON 格式）

```go
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
    // 设置响应内容类型为 JSON
    w.Header().Set("content-type", jsonContentType)
    
    // 使用 JSON 编码器将数据编码并发送
    json.NewEncoder(w).Encode(p.store.GetLeague())
}
```

**逐步解析**：

```
1. w.Header().Set("content-type", jsonContentType)
   └─ 告诉客户端：我发送的是 JSON 格式的数据
   └─ Header: content-type: application/json

2. json.NewEncoder(w).Encode(p.store.GetLeague())
   ├─ p.store.GetLeague() 获取排行榜数据
   │  └─ 返回 []Player{{"Alice", 20}, {"Bob", 10}}
   │
   ├─ json.NewEncoder(w) 创建 JSON 编码器
   │  └─ w 是 ResponseWriter，编码器将数据写入响应
   │
   └─ .Encode(...) 将数据编码为 JSON 并写入 w
      └─ 输出：[{"Name":"Alice","Wins":20},{"Name":"Bob","Wins":10}]
```

**HTTP 响应示例**：

```
HTTP/1.1 200 OK
Content-Type: application/json

[{"Name":"Alice","Wins":20},{"Name":"Bob","Wins":10}]
```

**客户端测试**：

```bash
curl http://localhost:5000/league

# 输出：
# [{"Name":"Alice","Wins":20},{"Name":"Bob","Wins":10}]
```

### 3.2 playersHandler - 处理玩家请求（GET/POST）

```go
func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
    // 从 URL 路径提取玩家名称
    player := r.URL.Path[len("/players/"):]
    
    // 根据 HTTP 方法进行不同的处理
    switch r.Method {
    case http.MethodPost:
        p.processWin(w, player)
    case http.MethodGet:
        p.showScore(w, player)
    }
}
```

**路径解析示例**：

```
请求：GET /players/Alice

r.URL.Path = "/players/Alice"
r.URL.Path[len("/players/"):] = r.URL.Path[9:] = "Alice"

player = "Alice"
```

**处理流程**：

```
HTTP 请求
    ↓
路由器匹配 /players/
    ↓
调用 playersHandler()
    ↓
提取玩家名称
    ↓
switch r.Method
├─ POST → processWin()      (记录胜利)
└─ GET → showScore()        (显示得分)
```

### 3.3 showScore - 返回玩家得分

```go
func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
    score := p.store.GetPlayerScore(player)
    
    // 如果没有这个玩家，返回 404
    if score == 0 {
        w.WriteHeader(http.StatusNotFound)
    }
    
    fmt.Fprint(w, score)
}
```

**逐步解析**：

```
1. score := p.store.GetPlayerScore(player)
   └─ 从存储中获取玩家得分

2. if score == 0
   └─ 检查玩家是否存在（得分为 0 表示不存在）

3. w.WriteHeader(http.StatusNotFound)
   └─ 设置状态码 404

4. fmt.Fprint(w, score)
   └─ 写入响应体（得分值）
```

**HTTP 响应示例**：

```
请求 1：GET /players/Alice（Alice 的得分为 20）
响应：
HTTP/1.1 200 OK
20

请求 2：GET /players/Unknown（Unknown 不存在）
响应：
HTTP/1.1 404 Not Found
0
```

### 3.4 processWin - 记录玩家胜利

```go
func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
    p.store.RecordWin(player)
    w.WriteHeader(http.StatusAccepted)
}
```

**逐步解析**：

```
1. p.store.RecordWin(player)
   └─ 在存储中记录玩家的胜利
   └─ 例如：Alice 的胜利次数 + 1

2. w.WriteHeader(http.StatusAccepted)
   └─ 设置状态码 202（Accepted）
   └─ 表示请求已被接受并处理
```

**HTTP 请求和响应**：

```
请求：POST /players/Alice
响应：
HTTP/1.1 202 Accepted
```

---

## 4. new() vs 指针赋值

### 4.1 两种创建指针的方式

```go
// 方式 1：使用 new()
p := new(PlayerServer)

// 方式 2：使用 & 操作符
ps := PlayerServer{}
p := &ps

// 方式 3：结合使用（推荐）
p := &PlayerServer{
    store: store,
}
```

### 4.2 详细对比

#### new() 方式

```go
p := new(PlayerServer)

// 1. 分配内存，创建零值实例
// 2. 返回指向该实例的指针 *PlayerServer
// 3. 字段初始化为零值：
//    - store: nil
//    - Handler: nil
```

**优点**：
- 简洁明了
- 只需一行代码
- 适合后续逐个赋值

**缺点**：
- 字段初始化为零值
- 需要分开赋值：`p.store = store`

#### & 操作符方式

```go
ps := PlayerServer{}
p := &ps

// 或者更简洁：
p := &PlayerServer{}

// 也可以在创建时初始化：
p := &PlayerServer{
    store: store,
}
```

**优点**：
- 可以同时初始化字段
- 更易读（看得出初始值）

**缺点**：
- 多一行代码
- 需要创建临时变量

### 4.3 性能和实际使用

**性能完全相同**：

```go
// 都会最终生成相同的机器码
p := new(PlayerServer)

p := &PlayerServer{}

p := &PlayerServer{store: store}
```

**代码比较**：

```go
// 使用 new() - 逐个赋值
func NewPlayerServer(store PlayerStore) *PlayerServer {
    p := new(PlayerServer)
    p.store = store
    
    router := http.NewServeMux()
    router.Handle("/league", http.HandlerFunc(p.leagueHandler))
    router.Handle("/players/", http.HandlerFunc(p.playersHandler))
    
    p.Handler = router
    return p
}

// 使用 & 操作符 - 但不推荐（路由配置无法在初始化时进行）
func NewPlayerServer(store PlayerStore) *PlayerServer {
    return &PlayerServer{
        store: store,
        // ❌ 无法在这里配置 router
    }
}
```

### 4.4 最佳实践

**在构造函数中：**

```go
// ✅ 推荐 1：使用 new() + 逐个赋值（当需要后续操作时）
func NewPlayerServer(store PlayerStore) *PlayerServer {
    p := new(PlayerServer)
    p.store = store
    
    router := http.NewServeMux()
    router.Handle("/league", http.HandlerFunc(p.leagueHandler))
    router.Handle("/players/", http.HandlerFunc(p.playersHandler))
    
    p.Handler = router
    return p
}

// ✅ 推荐 2：使用 & 操作符（当能直接初始化时）
type Config struct {
    Name string
}

func NewConfig(name string) *Config {
    return &Config{
        Name: name,
    }
}

// ❌ 不推荐：new() 但不赋值
p := new(PlayerServer)  // p.store 是 nil，容易出错
```

### 4.5 零值 vs 初始化值

```go
// new() 创建的是零值实例
p := new(PlayerServer)
// ├─ p.store = nil
// ├─ p.Handler = nil
// └─ 需要手动赋值

// & 可以同时初始化
p := &PlayerServer{
    store: store,
    Handler: router,  // 可选
}
// ├─ p.store = store
// ├─ p.Handler = router
// └─ 初始化完整
```

---

## 5. 工作流程

### 5.1 完整 HTTP 请求处理流程

```
客户端发送请求
    ↓
GET /league

    ↓

Go HTTP 服务器接收请求
    ├─ 解析请求路径 = "/league"
    └─ 解析 HTTP 方法 = GET

    ↓

路由器（http.ServeMux）匹配路由
    ├─ 检查 "/league" 是否在路由表中
    └─ 找到对应的处理函数 → leagueHandler

    ↓

调用 leagueHandler(w, r)
    ├─ 从存储获取排行榜数据
    │  └─ []Player{{"Alice", 20}, {"Bob", 10}}
    │
    ├─ 设置响应头
    │  └─ Content-Type: application/json
    │
    └─ 编码为 JSON 并写入响应
       └─ [{"Name":"Alice","Wins":20},{"Name":"Bob","Wins":10}]

    ↓

返回 HTTP 响应给客户端
    HTTP/1.1 200 OK
    Content-Type: application/json
    
    [{"Name":"Alice","Wins":20},{"Name":"Bob","Wins":10}]
```

### 5.2 POST 请求处理流程

```
客户端发送请求
    ↓
POST /players/Alice

    ↓

路由器匹配
    ├─ 路径 = "/players/Alice"
    ├─ 匹配路由 "/players/"
    └─ 调用 playersHandler(w, r)

    ↓

playersHandler 处理
    ├─ 提取玩家名称：player = "Alice"
    ├─ 检查 HTTP 方法 = POST
    └─ 调用 processWin(w, "Alice")

    ↓

processWin 处理
    ├─ p.store.RecordWin("Alice")
    │  └─ 存储中 Alice 的胜利次数 + 1
    │
    └─ w.WriteHeader(http.StatusAccepted)
       └─ 设置状态码 202

    ↓

返回响应
    HTTP/1.1 202 Accepted
```

### 5.3 测试流程

```
单元测试：TestLeague
    ↓
创建模拟数据存储 (StubPlayerStore)
    ├─ scores: nil
    ├─ winCalls: nil
    └─ league: [{"Cleo", 32}, {"Chris", 20}, {"Tiest", 14}]

    ↓

创建 PlayerServer
    └─ server := NewPlayerServer(&store)

    ↓

创建 HTTP 请求
    └─ request := newLeagueRequest()
    └─ 等价于：GET /league

    ↓

创建响应记录器
    └─ response := httptest.NewRecorder()

    ↓

直接调用 ServeHTTP（不启动真实服务器）
    └─ server.ServeHTTP(response, request)

    ↓

验证结果
    ├─ 验证状态码 = 200
    ├─ 验证响应头 Content-Type = application/json
    └─ 验证响应体 = JSON 格式的排行榜

    ↓

测试通过 ✅
```

### 5.4 集成测试

```
集成测试：TestRecordingWinsAndRetrievingThem
    ↓
创建真实数据存储
    └─ store := NewInMemoryPlayerStore()

    ↓

创建 PlayerServer
    └─ server := NewPlayerServer(store)

    ↓

模拟多个 POST 请求
    ├─ POST /players/Pepper (第 1 次)
    ├─ POST /players/Pepper (第 2 次)
    └─ POST /players/Pepper (第 3 次)
    └─ Pepper 的胜利次数 = 3

    ↓

验证 GET 请求
    ├─ GET /players/Pepper
    └─ 返回 "3"

    ↓

验证 GET /league 请求
    ├─ GET /league
    └─ 返回 JSON：[{"Name":"Pepper","Wins":3}]

    ↓

测试通过 ✅
```

---

## 总结

### 核心要点

| 概念 | 说明 |
|------|------|
| **路由** | 将 HTTP 请求分配到不同的处理函数 |
| **PlayerServer** | 包含数据存储和路由的 HTTP 服务器 |
| **leagueHandler** | 处理 GET /league，返回 JSON 格式的排行榜 |
| **playersHandler** | 处理 GET/POST /players/{name}，获取/记录玩家信息 |
| **json.NewEncoder** | 将 Go 对象编码为 JSON 并写入响应 |

### new() vs & 的使用场景

| 方式 | 使用场景 | 例子 |
|------|----------|------|
| `new(T)` | 需要后续操作和赋值 | NewPlayerServer 中使用 new() |
| `&T{}` | 可以直接初始化所有字段 | &Config{Name: "abc"} |
| `&T{Field: value}` | 推荐方式，既清晰又高效 | &PlayerServer{store: s} |

### 关键代码模式

```go
// 构造函数模式
func NewPlayerServer(store PlayerStore) *PlayerServer {
    p := new(PlayerServer)      // 1. 创建指针
    p.store = store              // 2. 赋值字段
    
    // 3. 复杂初始化
    router := http.NewServeMux()
    router.Handle("/league", http.HandlerFunc(p.leagueHandler))
    router.Handle("/players/", http.HandlerFunc(p.playersHandler))
    p.Handler = router
    
    return p                      // 4. 返回指针
}

// 处理器模式
func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")
    json.NewEncoder(w).Encode(p.store.GetLeague())
}

// 路由分发模式
func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
    player := r.URL.Path[len("/players/"):]
    switch r.Method {
    case http.MethodPost:
        p.processWin(w, player)
    case http.MethodGet:
        p.showScore(w, player)
    }
}
```

这就是 JSON 路由处理的核心！🚀
