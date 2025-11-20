```markdown
## 以下哪个是 Go 中有效的循环语句?
1. while
2. forever
3. until
4. for *正确答案*

> **4:** 正确。Go 中只有一种循环语句。


## 这段代码会打印什么?
```go
for i := 3; i > 0; i-- {
    fmt.Println(i)
}
```
1. 3 2 1 *正确答案*
2. 1 2 3
3. 0 1 2
4. 2 1 0


## 这段代码会打印什么?
```go
for i := 3; i > 0; {
    i--
    fmt.Println(i)
}
```
1. 3 2 1
2. 1 2 3
3. 0 1 2
4. 2 1 0 *正确答案*


## 这段代码会打印什么?
```go
for i := 3; ; {
    if i <= 0 {
        break
    }

    i--
    fmt.Println(i)
}
```
1. 3 2 1
2. 1 2 3
3. 0 1 2
4. 2 1 0 *正确答案*


## 这段代码会打印什么?
```go
for i := 2; i <= 9; i++ {
    if i % 3 != 0 {
        continue
    }

    fmt.Println(i)
}
```
1. 3 6 9 *正确答案*
2. 9 6 3
3. 2 3 6 9
4. 2 3 4 5 6 7 8 9


## 如何简化这段代码?
```go
for ; true ; {
    // ...
}
```
1. ```go
   for true {
   }
   ```
2. ```go
   for true; {
   }
   ```
3. ```go
   for {
   }
   ```
   *正确答案*
4. ```go
   for ; true {
   }
   ```


## 这段代码会打印什么?
假设你这样运行程序:
```bash
go run main.go go is awesome
```

```go
for i, v := range os.Args[1:] {
    fmt.Println(i+1, v)
}
```
1. ```
   1 go
   2 is
   3 awesome
   ```
   *正确答案*
2. ```
   go
   is
   awesome
   ```
3. ```
   0 go
   1 is
   2 awesome
   ```
4. ```
   1
   2
   3
   ```


## 这段代码会打印什么?
假设你这样运行程序:
```bash
go run main.go go is awesome
```

```go
for i := range os.Args[1:] {
    fmt.Println(i+1)
}
```
1. ```
   1 go
   2 is
   3 awesome
   ```
2. ```
   go
   is
   awesome
   ```
3. ```
   0 go
   1 is
   2 awesome
   ```
4. ```
   1
   2
   3
   ```
   *正确答案*


## 这段代码会打印什么?
假设你这样运行程序:
```bash
go run main.go go is awesome
```

```go
for _, v := range os.Args[1:] {
    fmt.Println(v)
}
```
1. ```
   1 go
   2 is
   3 awesome
   ```
2. ```
   go
   is
   awesome
   ```
   *正确答案*
3. ```
   0 go
   1 is
   2 awesome
   ```
4. ```
   1
   2
   3
   ```


## 这段代码会打印什么?
假设你这样运行程序:
```bash
go run main.go go is awesome
```

```go
var i int

for range os.Args {
    i++
}

fmt.Println(i)
```
1. go is awesome
2. 1 2 3
3. 2
4. 4 *正确答案*

> **4:** 如你所见,你也可以使用 for range 语句来计数。
```