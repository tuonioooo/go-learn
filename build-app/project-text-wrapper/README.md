# 文本换行挑战指南

在这个项目中，你的目标是模拟文本编辑器中的软换行功能。例如，当一行有100个字符，而软换行设置为40个字符时，编辑器可能会将超过40个字符的部分截断，并在下一行显示剩余内容。

## 示例

将给定文本按每行40个字符进行换行。例如，对于以下输入，程序应输出以下内容。

**输入：**

    Hello world, how is it going? It is ok.. The weather is beautiful.

**输出：**

    Hello world, how is it going? It is ok..
    The weather is beautiful.

## 规则

* 程序应支持Unicode文本。你可以在story.txt文件中找到Unicode文本。

* 程序不应在单词结束前切断单词。相反，它应该将整个单词放到下一行。

例如，以下是不正确的：

    Hello world, how is it goi
    ng? It is o
    k. The weather is beautifu
    l.

## 解决方案

* main.go