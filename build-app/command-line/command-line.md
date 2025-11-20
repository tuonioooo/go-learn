# Command-Line Poker 项目说明

## 项目结构概览

```
command-line/
├── CLI.go                          # CLI 包的核心实现
├── CLI_test.go                     # CLI 的单元测试
├── server.go                       # HTTP 服务器实现
├── server_test.go                  # 服务器单元测试
├── file_system_store.go            # 文件系统存储实现
├── file_system_store_test.go       # 存储单元测试
├── tape.go                         # 文件 I/O 辅助工具
├── league.go                       # 玩家排行榜逻辑
├── testing.go                      # 测试辅助函数
├── cli/                            # CLI 应用入口
│   └── main.go                     # 命令行应用程序
└── webserver/                      # Web 服务器应用入口
    └── main.go                     # HTTP Web 服务器程序
```

---

## 核心组件说明

### 1. CLI.go - 命令行交互器

**作用：** 定义 `CLI` 类型，负责处理命令行用户输入并调用存储接口记录比赛结果。

**关键代码：**
```go
type CLI struct {
    playerStore PlayerStore      // 存储接口，用于持久化数据
    in          *bufio.Scanner   // 用于读取用户输入
}

// NewCLI 创建一个新的 CLI 实例
func NewCLI(store PlayerStore, in io.Reader) *CLI

// PlayPoker 开始游戏，读取用户输入并记录赢家
func (cli *CLI) PlayPoker()

// extractWinner 从输入中解析赢家名字（如 "Chris wins" -> "Chris"）
func extractWinner(userInput string) string
```

**工作流程：**
1. 通过 `NewCLI()` 创建实例，注入存储接口和输入源
2. 调用 `PlayPoker()` 等待用户输入
3. 读取一行输入（如 "Chris wins"）
4. 提取赢家名字，调用 `playerStore.RecordWin()`
5. 游戏结束

---

### 2. CLI_test.go - CLI 单元测试

**作用：** 使用 TDD（测试驱动开发）方式验证 CLI 的正确性。

**测试场景：**

```go
// 场景1：记录 Chris 的胜利
func TestCLI(t *testing.T) {
    // 输入: "Chris wins\n"
    // 验证: playerStore 中记录了 Chris 的一次胜利
}

// 场景2：记录 Cleo 的胜利
func TestCLI(t *testing.T) {
    // 输入: "Cleo wins\n"
    // 验证: playerStore 中记录了 Cleo 的一次胜利
}

// 场景3：不要超过第一行读取
func TestCLI(t *testing.T) {
    // 输入: "Chris wins\n hello there"
    // 验证: 只读取到第一行，不继续读取
}
```

**测试中的模拟输入 vs 实际交互式输入：**

```go
// ❌ 测试中：使用 strings.NewReader 模拟输入（非交互式）
in := strings.NewReader("Chris wins\n")
cli := poker.NewCLI(store, in)
cli.PlayPoker()  // 直接读取预设的字符串，不等待用户输入

// ✅ 实际应用中：使用 os.Stdin 真实交互（交互式）
cli := poker.NewCLI(store, os.Stdin)
cli.PlayPoker()  // 等待用户在终端输入
```

| 特性 | strings.NewReader | os.Stdin |
|------|------------------|----------|
| **用途** | 测试 | 实际应用 |
| **交互性** | 非交互式，预设输入 | 交互式，等待用户输入 |
| **输入源** | 内存中的字符串 | 用户终端输入 |
| **适用场景** | 单元测试、自动化测试 | CLI 应用、生产环境 |

**测试工具：**
- `StubPlayerStore`：模拟的玩家存储，用于单元测试
- `failOnEndReader`：自定义读取器，确保不会读取超过需要的数据
- `AssertPlayerWin()`：断言函数，验证玩家记录是否正确

**如何运行测试：**
```bash
# 运行所有 CLI 测试
go test -v ./command-line

# 运行特定测试
go test -v -run TestCLI ./command-line

# 在当前目录运行测试
go test CLI_test.go
```

**测试输出示例：**
```
=== RUN   TestCLI
=== RUN   TestCLI/record_chris_win_from_user_input
=== RUN   TestCLI/record_cleo_win_from_user_input
=== RUN   TestCLI/do_not_read_beyond_the_first_newline
--- PASS: TestCLI (0.00s)
    --- PASS: TestCLI/record_chris_win_from_user_input (0.00s)
    --- PASS: TestCLI/record_cleo_win_from_user_input (0.00s)
    --- PASS: TestCLI/do_not_read_beyond_the_first_newline (0.00s)
PASS
```

---

### 3. cli/main.go - 命令行应用

**作用：** 提供命令行界面的应用程序入口，玩家可以在终端中输入比赛结果。

**应用流程：**

```go
func main() {
    // 1. 从 "game.db.json" 文件读取玩家数据
    store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
    defer close()  // 确保文件正确关闭
    
    // 2. 创建 CLI 实例，绑定存储和标准输入（os.Stdin）
    cli := poker.NewCLI(store, os.Stdin)
    
    // 3. 启动游戏
    cli.PlayPoker()  // 等待用户输入
}
```

**使用方式 - 交互式输入：**

```bash
# 1. 编译和运行
go run ./cli/main.go

# 2. 程序启动后会打印：
# Let's play poker
# Type {Name} wins to record a win
# [等待你的输入...]

# 3. 在终端输入玩家名字和 "wins"（然后按 Enter）
Chris wins
# 程序处理此输入，记录 Chris 的胜利，然后退出

# 再次运行，输入另一个玩家
Alice wins
```

**注意事项：**
- 程序运行后会**卡住等待输入** 🔄
- 你需要在终端中**手动输入**玩家名字 ⌨️
- 格式必须是 `{PlayerName} wins`
- 按 Enter 后程序记录结果并退出
- 数据自动保存到 `game.db.json` 文件

**完整交互流程：**

```
$ go run ./cli/main.go
Let's play poker
Type {Name} wins to record a win
Alice wins                           ← 你在这里输入，然后按 Enter
$                                    ← 程序处理完毕，返回命令行

$ go run ./cli/main.go               ← 再次运行
Let's play poker
Type {Name} wins to record a win
Bob wins                             ← 再输入一个玩家
$                                    ← 程序退出
```

**如果想支持多次输入，需要修改 `PlayPoker()` 方法加入循环：**

```go
// 修改前（当前）：只读取一次
func (cli *CLI) PlayPoker() {
    userInput := cli.readLine()
    cli.playerStore.RecordWin(extractWinner(userInput))
}

// 修改后（可选）：循环读取多次
func (cli *CLI) PlayPoker() {
    for {
        userInput := cli.readLine()
        if userInput == "quit" {
            break
        }
        cli.playerStore.RecordWin(extractWinner(userInput))
    }
}
```

---

### 4. webserver/main.go - Web 服务器应用

**作用：** 提供 HTTP Web 界面的应用程序入口，支持通过 HTTP API 进行比赛管理。

**应用流程：**

```go
func main() {
    // 1. 从文件读取玩家数据
    store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
    defer close()
    
    // 2. 创建 HTTP 服务器实例
    server := poker.NewPlayerServer(store)
    
    // 3. 启动 HTTP 服务，监听 5000 端口
    http.ListenAndServe(":5000", server)
}
```

**使用方式：**
```bash
# 编译和运行
go run ./webserver/main.go

# 然后通过浏览器或 curl 访问
# http://localhost:5000/players/{name}        获取玩家信息
# http://localhost:5000/league                 获取排行榜
# POST http://localhost:5000/players/{name}/win 记录胜利
```

**支持的 HTTP 接口：** 由 `server.go` 中的 `PlayerServer` 定义

---

## CLI vs WebServer 对比

| 特性 | CLI 应用 | Web 服务器应用 |
|------|---------|--------------|
| **接口** | 命令行终端 | HTTP API / Web 浏览器 |
| **输入方式** | 键盘输入 | HTTP 请求 |
| **适用场景** | 快速本地测试 | 多用户、远程访问 |
| **交互方式** | 同步、单行输入 | 异步、REST API |
| **启动** | `go run ./cli/main.go` | `go run ./webserver/main.go` |
| **端口** | 无（本地终端） | 5000 |

---

## 核心概念

### PlayerStore 接口
所有存储实现（文件系统、内存等）都必须实现此接口：
```go
type PlayerStore interface {
    GetPlayerScore(name string) int
    RecordWin(name string)
    GetLeague() League
}
```

### 数据流向

```
用户输入
   ↓
CLI / HTTP 请求
   ↓
CLI 结构体 / PlayerServer
   ↓
PlayerStore 接口（文件系统存储）
   ↓
game.db.json（持久化文件）
```

---

## 快速开始

### 1. 运行命令行应用
```bash
cd build-app/command-line
go run ./cli/main.go
```

### 2. 运行 Web 服务器
```bash
cd build-app/command-line
go run ./webserver/main.go
# 访问 http://localhost:5000
```

### 3. 运行所有测试
```bash
cd build-app/command-line
go test -v ./...
```

### 4. 运行特定测试
```bash
go test -v -run TestCLI ./command-line
go test -v -run TestPlayerServer ./command-line
```

---

## 文件说明

| 文件 | 类型 | 说明 |
|------|------|------|
| `CLI.go` | 实现 | CLI 类型定义和方法 |
| `CLI_test.go` | 测试 | CLI 单元测试 |
| `server.go` | 实现 | HTTP 服务器实现 |
| `server_test.go` | 测试 | 服务器单元测试 |
| `server_integration_test.go` | 集成测试 | 完整流程集成测试 |
| `file_system_store.go` | 实现 | 文件系统持久化实现 |
| `tape.go` | 工具 | 文件 I/O 辅助 |
| `league.go` | 实现 | 排行榜逻辑 |
| `testing.go` | 工具 | 测试辅助函数 |
| `cli/main.go` | 应用 | 命令行应用入口 |
| `webserver/main.go` | 应用 | Web 服务器应用入口 |

---

## 总结

这个项目展示了如何：
1. ✅ 使用 **接口** 进行依赖注入，使代码可测试
2. ✅ 用 **TDD** 方式开发，先写测试后实现
3. ✅ 实现 **多个应用入口**（CLI 和 Web）共用同一核心逻辑
4. ✅ 处理 **文件 I/O** 和数据持久化
5. ✅ 编写 **单元测试、集成测试和端到端测试**
