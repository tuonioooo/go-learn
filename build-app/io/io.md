# Go 文件 I/O 与持久化存储完整教程

## 目录
1. [核心概念](#1-核心概念)
2. [架构设计](#2-架构设计)
3. [关键组件详解](#3-关键组件详解)
4. [JSON 常用方法详解](#4-json-常用方法详解)
5. [数据持久化流程](#5-数据持久化流程)
6. [工作流程](#6-工作流程)

---

## 1. 核心概念

### 1.1 什么是持久化存储？

**问题**：内存中的数据在程序重启后就丢失了。

```
程序启动 → 数据存入内存 → 程序运行 → 程序退出
                              ↓
                         所有数据丢失 ❌
```

**解决**：将数据保存到文件系统中。

```
程序启动
    ↓
从文件读取数据
    ↓
修改数据
    ↓
写回文件
    ↓
程序退出
    ↓
重启程序
    ↓
文件中的数据仍然存在 ✅
```

### 1.2 本项目的存储架构

本项目实现了**两种存储方式**，供用户选择：

```
PlayerStore 接口
├─ InMemoryPlayerStore（内存存储）
│  └─ 数据存储在 map 中，程序退出后丢失
│
└─ FileSystemPlayerStore（文件存储）
   └─ 数据保存在 JSON 文件中，持久化
```

### 1.3 为什么使用接口？

使用 `PlayerStore` 接口，让 HTTP 服务器无需关心数据从哪里来。

```go
type PlayerStore interface {
    GetPlayerScore(name string) int      // 获取玩家得分
    RecordWin(name string)               // 记录胜利
    GetLeague() League                   // 获取排行榜
}
```

**优势**：
- 服务器代码不变
- 只需替换存储实现
- 便于测试（可用 Stub 替代真实存储）

---

## 2. 架构设计

### 2.1 分层架构

```
┌─────────────────────────────┐
│   HTTP 服务器               │
│   (PlayerServer)            │
│   - 处理请求                │
│   - 返回响应                │
└──────────────┬──────────────┘
               │ 依赖
               ↓
┌─────────────────────────────┐
│   PlayerStore 接口          │
│   - GetPlayerScore()        │
│   - RecordWin()             │
│   - GetLeague()             │
└──────────────┬──────────────┘
               │ 可有多个实现
       ┌───────┴─────────┐
       ↓                 ↓
┌─────────────┐   ┌──────────────────┐
│   内存存储   │   │   文件存储       │
│  In Memory  │   │   File System    │
│   Store     │   │   Store          │
└─────────────┘   └──────────────────┘
```

### 2.2 文件结构

```
io/
├─ server.go                  ← HTTP 处理器
├─ in_memory_player_store.go  ← 内存存储实现
├─ file_system_store.go       ← 文件存储实现
├─ league.go                  ← 排行榜数据结构
├─ tape.go                    ← 文件写入工具
├─ main.go                    ← 程序入口
├─ server_test.go             ← 单元测试
├─ server_integration_test.go ← 集成测试
├─ file_system_store_test.go  ← 文件存储测试
├─ tape_test.go               ← 文件写入测试
└─ io.md                       ← 本文档
```

---

## 3. 关键组件详解

### 3.1 League（排行榜）

```go
// League 是玩家列表的类型别名
type League []Player

// Player 存储玩家信息
type Player struct {
    Name string  // 玩家名称
    Wins int     // 胜利次数
}

// Find 方法：从排行榜中查找玩家
func (l League) Find(name string) *Player {
    for i, p := range l {
        if p.Name == name {
            return &l[i]  // 返回指针，便于修改
        }
    }
    return nil  // 找不到返回 nil
}

// NewLeague 方法：从 JSON 读取排行榜数据
func NewLeague(rdr io.Reader) (League, error) {
    var league []Player
    err := json.NewDecoder(rdr).Decode(&league)
    return league, err
}
```

**数据结构示例**：

```json
[
  {"Name": "Alice", "Wins": 20},
  {"Name": "Bob", "Wins": 15},
  {"Name": "Charlie", "Wins": 10}
]
```

**Find 方法的作用**：

```go
league := League{
    {"Alice", 20},
    {"Bob", 15},
}

player := league.Find("Bob")
// ↓
// &Player{"Bob", 15}  ← 返回指针

player.Wins++  // 可以直接修改
```

### 3.2 tape（文件写入工具）

**问题**：如何让文件内容被完全替换？

```go
type tape struct {
    file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
    t.file.Truncate(0)      // ① 清空文件
    t.file.Seek(0, 0)        // ② 移动指针到开头
    return t.file.Write(p)  // ③ 写入新数据
}
```

**执行流程**：

```
原文件内容：
{"Name": "Alice", "Wins": 5}

调用 tape.Write([]byte(`[{"Name":"Alice","Wins":6}]`))
    ↓
1. Truncate(0)：清空文件
   文件内容：（空）

2. Seek(0, 0)：指针移到开头
   文件指针：位置 0

3. Write(...)：写入新数据
   文件内容：[{"Name":"Alice","Wins":6}]

结果 ✅
```

**为什么需要 Truncate？**

```go
// ❌ 不使用 Truncate 的问题
file.Seek(0, 0)
file.Write([]byte("123"))  // 写入 "123"

// 如果原文件内容是 "abcde"（5 字节）
// 写入 "123" 后：file = "123de" ❌ 数据混乱

// ✅ 使用 Truncate 的正确做法
file.Truncate(0)          // 清空文件
file.Seek(0, 0)
file.Write([]byte("123")) // 写入 "123"
// 结果：file = "123" ✅ 正确
```

### 3.3 InMemoryPlayerStore（内存存储）

```go
type InMemoryPlayerStore struct {
    store map[string]int
    lock  sync.RWMutex  // 并发安全
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
    i.lock.Lock()
    defer i.lock.Unlock()
    i.store[name]++
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
    i.lock.RLock()
    defer i.lock.RUnlock()
    return i.store[name]
}
```

**特点**：
- 数据存储在内存 map 中
- 使用 RWMutex 确保并发安全
- 程序退出时数据丢失

**使用场景**：
- 测试
- 临时数据
- 服务器重启不重要的数据

### 3.4 FileSystemPlayerStore（文件存储）

#### 初始化

```go
func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
    // 1. 初始化数据库文件
    err := initialisePlayerDBFile(file)
    if err != nil {
        return nil, fmt.Errorf("problem initialising player db file, %v", err)
    }

    // 2. 从文件读取现有数据
    league, err := NewLeague(file)
    if err != nil {
        return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
    }

    // 3. 创建编码器用于写入
    return &FileSystemPlayerStore{
        database: json.NewEncoder(&tape{file}),
        league:   league,
    }, nil
}

func initialisePlayerDBFile(file *os.File) error {
    file.Seek(0, 0)
    info, err := file.Stat()
    
    if err != nil {
        return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
    }

    // 如果文件为空，写入初始化数据 []
    if info.Size() == 0 {
        file.Write([]byte("[]"))
        file.Seek(0, 0)
    }

    return nil
}
```

**初始化流程**：

```
打开数据库文件
    ↓
检查文件是否为空
├─ 是：写入 "[]"（空数组）
└─ 否：继续
    ↓
从文件读取 JSON 数据
    ↓
解析为 League
    ↓
创建 FileSystemPlayerStore
    ↓
返回
```

#### 读取操作

```go
// GetLeague 返回排序后的排行榜
func (f *FileSystemPlayerStore) GetLeague() League {
    sort.Slice(f.league, func(i, j int) bool {
        return f.league[i].Wins > f.league[j].Wins  // 按胜利次数降序
    })
    return f.league
}

// GetPlayerScore 获取单个玩家得分
func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
    player := f.league.Find(name)
    if player != nil {
        return player.Wins
    }
    return 0
}
```

#### 写入操作

```go
func (f *FileSystemPlayerStore) RecordWin(name string) {
    // 1. 从内存中的 league 查找玩家
    player := f.league.Find(name)
    
    if player != nil {
        // 2. 玩家存在，增加胜利次数
        player.Wins++
    } else {
        // 3. 玩家不存在，添加新玩家
        f.league = append(f.league, Player{name, 1})
    }
    
    // 4. 将更新后的 league 写回文件
    f.database.Encode(f.league)
}
```

**数据流**：

```
RecordWin("Alice")
    ↓
查找 Alice
├─ 找到：Wins++
└─ 未找到：append 新玩家
    ↓
使用 json.NewEncoder 编码 league
    ↓
通过 tape 写入文件
    ↓
tape 清空文件 → 写入新数据
    ↓
文件已更新 ✅
```

---

## 4. JSON 常用方法详解

### 4.1 JSON 序列化和反序列化基础

**什么是 JSON？**

JSON（JavaScript Object Notation）是一种轻量级的数据交换格式。

```go
// Go 结构体
type Player struct {
    Name string
    Wins int
}

// 对应的 JSON
{"Name":"Alice","Wins":20}

// JSON 数组
[{"Name":"Alice","Wins":20},{"Name":"Bob","Wins":15}]
```

### 4.2 json.Encoder - 编码（Go → JSON）

**作用**：将 Go 对象转换为 JSON 格式并写入 I/O。

#### 基础使用

```go
import (
    "encoding/json"
    "os"
)

type Player struct {
    Name string
    Wins int
}

// 创建编码器，写入到文件
file, _ := os.Create("output.json")
encoder := json.NewEncoder(file)

// 编码单个对象
player := Player{"Alice", 20}
encoder.Encode(player)
// 文件内容：{"Name":"Alice","Wins":20}

// 编码数组
players := []Player{
    {"Alice", 20},
    {"Bob", 15},
}
encoder.Encode(players)
// 文件内容：[{"Name":"Alice","Wins":20},{"Name":"Bob","Wins":15}]
```

#### 编码流程

```
encoder.Encode(data)
    ↓
1. 遍历结构体字段
2. 将每个字段转换为 JSON 格式
3. 通过 Write 写入到文件
4. 自动添加换行符 \n
```

#### 与 tape 结合使用

```go
// FileSystemPlayerStore 中的用法
database: json.NewEncoder(&tape{file})

// 当执行 database.Encode(f.league) 时：
// 1. json.NewEncoder 创建编码器
// 2. 将 league 编码为 JSON
// 3. 通过 &tape{file} 的 Write 方法写入
//    ├─ tape.Write() 先清空文件（Truncate）
//    ├─ 移动指针到开头（Seek）
//    └─ 写入新数据（Write）
```

**完整示例**：

```go
type FileSystemPlayerStore struct {
    database *json.Encoder
    league   League
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
    player := f.league.Find(name)
    
    if player != nil {
        player.Wins++
    } else {
        f.league = append(f.league, Player{name, 1})
    }
    
    // 使用编码器写入
    f.database.Encode(f.league)
    // ↓
    // 文件内容自动更新为最新的 JSON 数据
}
```

### 4.3 json.Decoder - 解码（JSON → Go）

**作用**：从 I/O 读取 JSON 并转换为 Go 对象。

#### 基础使用

```go
import (
    "encoding/json"
    "os"
)

type Player struct {
    Name string
    Wins int
}

// 打开文件
file, _ := os.Open("players.json")

// 创建解码器
decoder := json.NewDecoder(file)

// 解码单个对象
var player Player
decoder.Decode(&player)
// 从文件读取 JSON，转换为 Player 结构体

// 解码数组
var players []Player
decoder.Decode(&players)
// 从文件读取 JSON 数组，转换为 []Player
```

#### 解码流程

```
decoder.Decode(&data)
    ↓
1. 从文件读取 JSON 文本
2. 解析 JSON 格式
3. 转换为 Go 对象
4. 存储到指定的地址 (&data)
```

#### 为什么使用指针？

```go
// ❌ 错误：Decode 需要指针
var player Player
decoder.Decode(player)  // 编译错误

// ✅ 正确：使用指针
var player Player
decoder.Decode(&player)  // OK

// 原因：Decode 需要修改 player 的值
// 所以必须传递指针，这样才能赋值
```

#### League 中的应用

```go
// league.go 中的应用
func NewLeague(rdr io.Reader) (League, error) {
    var league []Player
    
    // 创建解码器
    err := json.NewDecoder(rdr).Decode(&league)
    
    if err != nil {
        err = fmt.Errorf("problem parsing league, %v", err)
    }
    
    return league, err
}

// 使用示例
file, _ := os.Open("game.db.json")
league, err := NewLeague(file)
// ↓
// league 包含从文件读取的所有玩家数据
```

### 4.4 json.Marshal - 编码为字节

**作用**：将 Go 对象转换为 JSON 字节数组。

```go
import "encoding/json"

type Player struct {
    Name string
    Wins int
}

player := Player{"Alice", 20}

// 编码为字节
jsonBytes, err := json.Marshal(player)

// jsonBytes = []byte(`{"Name":"Alice","Wins":20}`)
// 字符串形式：{"Name":"Alice","Wins":20}

// 适合：
// - 网络传输
// - 保存到字符串
// - 发送 HTTP 响应
```

#### vs Encoder 的区别

```go
// json.Marshal - 返回字节
jsonBytes, err := json.Marshal(player)
// 返回：[]byte, error

// json.Encoder - 直接写入 I/O
encoder := json.NewEncoder(writer)
encoder.Encode(player)
// 不返回字节，直接写入

// 对比：
// Marshal：内存 → 字节 → 再处理
// Encoder：内存 → 直接写入 I/O（更高效）
```

**实际使用**：

```go
// HTTP 响应中的使用
func sendJSON(w http.ResponseWriter, data interface{}) {
    // 方式 1：使用 Encoder（推荐，效率高）
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
    
    // 方式 2：使用 Marshal
    w.Header().Set("Content-Type", "application/json")
    jsonBytes, _ := json.Marshal(data)
    w.Write(jsonBytes)
}
```

### 4.5 json.Unmarshal - 解码从字节

**作用**：将 JSON 字节数组转换为 Go 对象。

```go
import "encoding/json"

type Player struct {
    Name string
    Wins int
}

// JSON 字节
jsonBytes := []byte(`{"Name":"Alice","Wins":20}`)

// 解码为对象
var player Player
err := json.Unmarshal(jsonBytes, &player)

// player = Player{"Alice", 20}

// 适合：
// - 接收 HTTP 请求体
// - 处理已加载的字节数据
// - API 响应解析
```

#### vs Decoder 的区别

```go
// json.Unmarshal - 从字节解码
var player Player
err := json.Unmarshal([]byte(`{"Name":"Alice","Wins":20}`), &player)

// json.Decoder - 从 I/O 解码
var player Player
err := json.NewDecoder(reader).Decode(&player)

// 对比：
// Unmarshal：所有数据都在内存中
// Decoder：流式处理，逐步读取（大文件更节省内存）
```

### 4.6 json 标签（Tags）

**作用**：控制 JSON 字段的名称、忽略字段、默认值等。

```go
type Player struct {
    // 标签：`json:"name"`
    // 含义：JSON 中的字段名为 "name"（默认是 "Name"）
    Name string `json:"name"`
    Wins int    `json:"wins"`
}

player := Player{"Alice", 20}
jsonBytes, _ := json.Marshal(player)

// jsonBytes = []byte(`{"name":"Alice","wins":20}`)
// 注意：JSON 中是小写的 "name" 和 "wins"
```

#### 常见标签选项

```go
type Player struct {
    // 标准名称
    Name string `json:"name"`
    
    // 忽略该字段（JSON 中不包含）
    Password string `json:"password"`
    Secret   string `json:"-"`
    
    // 如果为空值则忽略
    Email string `json:"email,omitempty"`
    
    // 字段必填
    ID int `json:"id"`
}

// 示例：
player := Player{
    Name:     "Alice",
    Password: "secret",
    Secret:   "hidden",
    Email:    "",  // 空值
    ID:       1,
}

jsonBytes, _ := json.Marshal(player)
// 输出：{"name":"Alice","password":"secret","id":1}
// 解析：
// - Secret 被 "-" 忽略（不出现）
// - Email 被 "omitempty" 忽略（值为空）
```

### 4.7 JSON 编码和解码的完整流程

```
写入文件：
Go 对象 → json.Marshal → []byte → 写入文件
        或 json.Encoder 直接写入

读取文件：
文件内容 → []byte → json.Unmarshal → Go 对象
      或 json.Decoder 直接读取
```

**本项目的实现**：

```
程序启动：
game.db.json → json.NewDecoder(file) → json.Decode(&league) → League

修改数据：
RecordWin("Alice")
    ↓
f.league[...].Wins++
    ↓
json.NewEncoder(&tape{file}) → .Encode(f.league)
    ↓
league → JSON → tape 清空文件 → 写入文件

读取排行榜：
GET /league
    ↓
f.league（已在内存中）→ json.NewEncoder(w) → .Encode(f.league)
    ↓
league → JSON → 返回给客户端
```

### 4.8 常见错误和最佳实践

#### 错误 1：未导出的字段

```go
type Player struct {
    name string  // ❌ 小写，无法序列化
    wins int
}

// 解决：大写首字母（导出）
type Player struct {
    Name string  // ✅
    Wins int
}
```

#### 错误 2：忘记指针

```go
var player Player

// ❌ Decode 需要指针
json.Unmarshal(data, player)

// ✅ 使用指针
json.Unmarshal(data, &player)
```

#### 错误 3：忽视错误处理

```go
// ❌ 忽视可能的解析错误
var player Player
json.Unmarshal(data, &player)

// ✅ 检查错误
var player Player
if err := json.Unmarshal(data, &player); err != nil {
    log.Fatalf("Failed to parse JSON: %v", err)
}
```

#### 最佳实践

```go
// 1. 总是检查错误
league, err := NewLeague(file)
if err != nil {
    return nil, fmt.Errorf("problem loading league: %v", err)
}

// 2. 使用 Encoder/Decoder 处理流（高效）
encoder := json.NewEncoder(file)
encoder.Encode(league)

// 3. 对大结构体，使用流式处理避免加载整个文件到内存
decoder := json.NewDecoder(file)
decoder.Decode(&league)

// 4. 使用 JSON 标签控制序列化
type Player struct {
    Name string `json:"name"`
    Wins int    `json:"wins"`
}

// 5. 验证数据有效性
if len(league) == 0 {
    return nil, fmt.Errorf("no players in league")
}
```

---

## 5. 数据持久化流程

### 5.1 程序启动流程

```go
func main() {
    // 1. 打开（或创建）数据库文件
    db, err := os.OpenFile(
        "game.db.json",
        os.O_RDWR|os.O_CREATE,  // 读写模式，文件不存在则创建
        0666,                    // 权限
    )
    if err != nil {
        log.Fatalf("problem opening %s %v", dbFileName, err)
    }

    // 2. 创建文件存储
    store, err := NewFileSystemPlayerStore(db)
    if err != nil {
        log.Fatalf("problem creating file system player store, %v ", err)
    }

    // 3. 创建 HTTP 服务器
    server := NewPlayerServer(store)

    // 4. 启动服务器
    if err := http.ListenAndServe(":5000", server); err != nil {
        log.Fatalf("could not listen on port 5000 %v", err)
    }
}
```

**流程图**：

```
启动程序
    ↓
打开文件 game.db.json
    ├─ 文件存在：打开
    └─ 文件不存在：创建
    ↓
初始化 FileSystemPlayerStore
    ├─ 检查文件是否为空
    ├─ 若为空，写入 "[]"
    └─ 从文件读取数据到内存
    ↓
创建 HTTP 服务器
    ├─ 配置路由
    └─ 保存 store 引用
    ↓
启动服务器监听 :5000
    ↓
等待客户端请求
```

### 4.2 读取数据的流程

```
客户端请求：GET /league
    ↓
服务器接收请求
    ↓
调用 store.GetLeague()
    ↓
FileSystemPlayerStore.GetLeague()
    ├─ 对内存中的 league 排序
    └─ 返回排序后的列表
    ↓
转换为 JSON 格式
    ↓
返回给客户端
```

**关键点**：数据已经在程序启动时加载到内存中，读取非常快。

### 5.2 写入数据的流程

```
客户端请求：POST /players/Alice
    ↓
服务器接收请求
    ↓
调用 store.RecordWin("Alice")
    ↓
FileSystemPlayerStore.RecordWin()
    ├─ 从内存 league 中查找 Alice
    ├─ 若存在：Wins++
    ├─ 若不存在：append 新玩家
    │
    └─ 使用 json.NewEncoder(tape) 将 league 写入文件
       ├─ tape.Write() 清空文件
       ├─ 写入新的 JSON 数据
       └─ 文件保存成功 ✅
    ↓
返回 202 Accepted 给客户端
```

**重要**：每次 RecordWin 都会同步写入文件，确保数据不丢失。

---

## 6. 工作流程

### 6.1 完整的增删改查流程

```
程序启动
    ├─ 打开 game.db.json
    ├─ 从文件读取所有玩家数据
    └─ 加载到内存

客户端请求 1：POST /players/Alice
    ↓
RecordWin("Alice")
    ├─ 内存：league.append({"Alice", 1})
    └─ 文件：写入 [{"Name":"Alice","Wins":1}]

客户端请求 2：POST /players/Alice
    ↓
RecordWin("Alice")
    ├─ 内存：league[0].Wins = 2
    └─ 文件：写入 [{"Name":"Alice","Wins":2}]

客户端请求 3：POST /players/Bob
    ↓
RecordWin("Bob")
    ├─ 内存：league.append({"Bob", 1})
    └─ 文件：写入 [{"Name":"Alice","Wins":2},{"Name":"Bob","Wins":1}]

客户端请求 4：GET /league
    ↓
GetLeague()
    ├─ 内存排序：[{"Alice", 2}, {"Bob", 1}]
    └─ 返回 JSON 给客户端

程序关闭
    ↓
所有数据保存在 game.db.json 中 ✅

下次启动程序
    ↓
从文件重新加载数据
    ↓
Alice 仍有 2 次胜利，Bob 仍有 1 次 ✅
```

### 6.2 集成测试流程

```go
func TestRecordingWinsAndRetrievingThem(t *testing.T) {
    // 1. 创建临时文件，初始内容为 "[]"
    database, cleanDatabase := createTempFile(t, `[]`)
    defer cleanDatabase()  // 测试完成后删除临时文件
    
    // 2. 创建文件存储
    store, err := NewFileSystemPlayerStore(database)
    assertNoError(t, err)
    
    // 3. 创建 HTTP 服务器
    server := NewPlayerServer(store)
    player := "Pepper"
    
    // 4. 模拟 POST 请求 3 次（Pepper 胜利 3 次）
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
    
    // 5. 验证 GET /players/Pepper
    t.Run("get score", func(t *testing.T) {
        response := httptest.NewRecorder()
        server.ServeHTTP(response, newGetScoreRequest(player))
        assertStatus(t, response.Code, http.StatusOK)
        assertResponseBody(t, response.Body.String(), "3")  // ✅ 得分为 3
    })
    
    // 6. 验证 GET /league
    t.Run("get league", func(t *testing.T) {
        response := httptest.NewRecorder()
        server.ServeHTTP(response, newLeagueRequest())
        assertStatus(t, response.Code, http.StatusOK)
        
        got := getLeagueFromResponse(t, response.Body)
        want := []Player{{"Pepper", 3}}
        assertLeague(t, got, want)  // ✅ 排行榜正确
    })
}
```

**测试验证**：
- 数据成功写入文件
- 数据成功从文件读取
- 多次操作的累积正确

### 6.3 单元测试 vs 集成测试

**单元测试**（server_test.go）：
```go
// 使用 StubPlayerStore（模拟存储）
// 只测试 HTTP 处理逻辑
store := StubPlayerStore{...}
server := NewPlayerServer(&store)
server.ServeHTTP(response, request)
```

**集成测试**（server_integration_test.go）：
```go
// 使用真实的 FileSystemPlayerStore
// 测试文件读写和 HTTP 处理的整体流程
database, cleanDatabase := createTempFile(t, `[]`)
store, err := NewFileSystemPlayerStore(database)
server := NewPlayerServer(store)
server.ServeHTTP(response, request)
```

---

## 总结

### 核心要点

| 组件 | 作用 | 特点 |
|------|------|------|
| **PlayerStore** | 存储接口 | 多个实现可选 |
| **InMemoryPlayerStore** | 内存存储 | 快，不持久 |
| **FileSystemPlayerStore** | 文件存储 | 持久，可恢复 |
| **League** | 排行榜数据 | 玩家列表 + Find 方法 |
| **tape** | 文件写工具 | Truncate + Write |
| **json.Encoder/Decoder** | JSON 序列化 | 编码/解码数据 |

### 关键流程

```
写入流程：
修改内存 → 编码为 JSON → 清空文件 → 写入文件 ✓ 持久化

读取流程：
打开文件 → 读取 JSON → 解码为对象 → 返回给客户端

重启后：
从文件读取 → 恢复所有数据 ✓ 持久化成功
```

### 最佳实践

1. **使用接口**：隐藏存储实现细节
2. **同步写入**：每次修改都立即写入文件，避免数据丢失
3. **并发安全**：使用 RWMutex 保护并发访问
4. **错误处理**：初始化时检查文件和解析错误
5. **测试充分**：单元测试和集成测试结合

这就是 Go 文件 I/O 和持久化存储的完整实现！🚀
