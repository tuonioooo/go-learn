# Go 语言输入流和 bufio 包相关知识测验

## 什么是输入流？
1. 任何生成连续数据的数据源，如标准输入 *正确*
2. 来自用户的输入
3. 文件的内容
4. 网站的内容

> **1:** 正确。输入流可以来自任何数据源。标准输入是其中之一，您可以将其重定向到几乎任何数据源，如用户输入、文件、网站内容等。
>
> **2-4:** 是的，这些可能是输入流，但它们没有解释什么是输入流。

## 这个程序会打印什么？
```go
in := bufio.Scanner(os.Stdin)
in.Scan() // 用户输入: "hi!"
in.Scan() // 用户输入: "how are you?"
fmt.Println(in.Text())
```
1. "hi" 和 "how are you?"
2. "hi"
3. "how are you?" *正确*
4. 什么都不打印

> **3:** Text() 方法只返回最后扫描的令牌。令牌可以是一行、一个单词等。

## 使用 bufio.Scanner 时，如何检测输入流中的错误？
1. 使用 Err() 方法 *正确*
2. 使用 Error() 方法
3. 通过检查 Scanner 是否为 nil

## 如何配置 bufio.Scanner 只扫描单词？
```go
in := bufio.Scanner(os.Stdin)
// ...
```
1. `in = bufio.NewScanner(in, bufio.ScanWords)`
2. `in.Split(bufio.ScanWords)` *正确*
3. `in.ScanWords()`

> **2:** 正确。bufio 有一些分割器，如 ScanWords，还有 ScanLines（默认）、ScanRunes 等。

## 这个函数使用 "Must" 前缀，为什么？
```go
regexp.MustCompile("...")
```
1. 只是为了提高可读性
2. 保证函数无论如何都会工作
3. 函数可能会使程序崩溃 *正确*

> **3:** "Must" 前缀是一个约定。如果一个函数或方法可能会 panic（= 使程序崩溃），那么通常会在前面加上 "must" 前缀。