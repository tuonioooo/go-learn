# 弹跳球挑战提示

请在遇到困难时参考以下提示。本文档内容无特定顺序，请勿严格按顺序遵循。

## 计算速度

你可以使用速度来改变球的速度和位置。在我的示例中，速度是恒定的，所以我总是使用单位值：1。

*   在每次循环步骤中：将速度值加到球的位置上。这将使球移动。

*   **速度的含义：速率和方向**
    *   X 速度 = 1 -> _球向右移动_
    *   X 速度 = -1 -> _球向左移动_
    *   Y 速度 = 1 -> _球向下移动_
    *   Y 速度 = -1 -> _球向上移动_

*   **关于图形和速度的更多信息：**
    *   https://www.youtube.com/watch?v=7Jr0SFMQ4Rs&t=529
    *   https://www.youtube.com/watch?v=ZM8ECpBuQYE

## 创建面板

我使用 `[][]bool` 来表示面板，但你可以使用任何你喜欢的数据类型。例如，你可以直接使用 `[][]rune` 或 `[]rune`。尝试一下它们，然后决定哪种最适合你。

## 清空屏幕

*   在循环开始之前，使用我的 https://github.com/inancgumus/screen 清空屏幕一次，点击链接查看。你可以在https://godoc.org/github.com/inancgumus/screen。
*   在每次循环步骤之后，使用 screen 包将光标移动到左上角位置。这样你就可以在同一位置重新绘制动画帧。

*   你可以在 https://github.com/inancgumus/learngo/tree/master/15-project-retro-led-clock 中找到关于 screen 包和清屏的更多信息。

## 绘制面板

不要每次都直接将面板和球绘制到屏幕上，你可以先填充一个缓冲区，完成后，通过打印缓冲区来一次性绘制面板和球。我使用一个 `[]rune` 切片作为缓冲区，因为 `rune` 可以存储一个表情符号字符。

*   使用 `make` 函数创建一个足够大的 rune 切片，命名为 `buf`。
    *   **提示：** `宽度 * 高度` 将为你提供一个足够大的缓冲区。
    *   **技巧：** 你也可以使用 `string` 拼接来绘制到 `string` 缓冲区，但那样效率较低。
    *   你可以在字符串章节中找到关于字节和 rune 的更多信息。

```go
// 转换缓冲区的技巧
var buffer []rune

// 为了打印，你可以像这样将 rune 切片转换为字符串：
str := string(buffer)
```

## 降低速度

调用 `time.Sleep` 函数来稍微降低循环的速度，这样你就能看清球了 :)

```go
time.Sleep(time.Second / 20)
```