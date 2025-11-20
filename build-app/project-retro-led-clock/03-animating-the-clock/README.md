

# **目标 3：让时钟动起来（Animate the Clock）**

## **说明（Notes）**

* `main.go` 文件中包含了上一步的解决方案。
* `solution` 文件夹中包含此步骤的完整参考实现。

---

## **挑战步骤（Challenge Steps）**

1. **创建一个无限循环，用来更新时钟**

2. **让时钟每秒更新一次**

   使用如下代码让程序暂停 1 秒钟：

   `time.Sleep(time.Second)`

3. **在无限循环开始前清屏**

   1. 使用我编写的清屏库：

      ```bash
      go get -u github.com/inancgumus/screen
      ```

   2. 在代码中导入并调用它：

      ```go
      screen.Clear()
      ```

   3. 如果你使用的是 **Go Playground**，请改用：

      ```go
      fmt.Println("\f")
      ```

4. **在每次循环开始前，将光标移动到屏幕左上角**

   * 在代码中调用：

     ```go
     screen.MoveTopLeft()
     ```

   * 如果你使用的是 **Go Playground**，同样使用：

     ```go
     fmt.Println("\f")
     ```

---

# **额外说明（SIDE NOTE FOR THE CURIOUS）**

如果你对清屏库的实现感兴趣，可以继续阅读下面的内容。

### **在 bash（类 Unix 系统）中**

该库使用特殊的终端控制指令。如果你打开源码，你会看到：

* `\033` 是一个特殊的控制码
* `[2J` 表示清空屏幕并重置光标位置
* `[H` 表示把光标移动到屏幕 `(0,0)` 位置
* 更多信息可以参考文档：[https://bluesock.org/~willkg/dev/ansi.html](https://bluesock.org/~willkg/dev/ansi.html)

---

### **在 Windows 中**

我直接调用了 Windows 的 API 函数。这部分对目前的课程内容来说太高级了，不过以后我可能会解释。

因此，该清屏包会根据编译所在的平台自动选择使用哪一种实现：

* 在 **Windows** 上 → 使用 Windows API
* 在 **其他系统**（Linux / macOS）上 → 使用 bash 的特殊控制代码
