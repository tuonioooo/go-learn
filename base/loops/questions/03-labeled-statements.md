
# Go 语言中的标签（Label）、break、continue、goto — 中文解析文档

## 📌 1. 标签属于哪个作用域？

1. 属于其所在语句的作用域
2. **属于整个函数体的作用域（正确）**
3. 属于包级作用域

**解释：**

* **标签不属于语句作用域**，不同于变量、常量等名称
* **标签在整个函数内都可被引用**，甚至在声明之前就能被使用
* 这也是为什么 `goto` 能跳到函数内任意标签（只要遵守一些跳转限制规则）

---

## 📌 2. 下面两层 for 循环中，标签 `words` 标记的是哪一个语句？

```go
for range words {
words:
	for range letters {
		// ...
	}
}
```

1. 外层循环
2. **内层循环（正确）**
3. 所有循环

**解释：**

* 标签只会标记紧挨着它的语句
* 因此 `words:` 只标记第二个（内层）循环

---

## 📌 3. 下面这段代码中的 break 会终止整个循环吗？

```go
package main

func main() {
main:
	for {
		switch "A" {
		case "A":
			break // <- 这里！
		case "B":
			continue main
		}
	}
}
```

1. **不会，break 只会终止 switch，而循环继续（正确）**
2. 会终止循环
3. 会终止 switch

**解释：**

* `break` 若未指定标签，则只会终止 **最内层语句**
* 这里 break 的是 `switch`，不是 `for`
* 因此循环不会终止，会继续执行下一轮

---

## 📌 4. 下面这段循环会终止吗？

```go
package main

func main() {
	flag := "A"

main:
	for {
		switch flag {
		case "A":
			flag = "B"
			break
		case "B":
			break main
		}
	}
}
```

1. 不会，循环将永远继续
2. 会，第一次 break 会终止循环
3. **会，第二次 break 会终止循环（正确）**

**解释：**

* 第一次循环里 `flag == "A"`，修改 flag 为 `"B"`，但只是普通 break → 退出 switch，不退出 for
* 下一轮进入 `"B"` 分支，执行 `break main` → 跳出标记的循环
* 所以程序能正常结束

---

## 📌 5. 下列程序中的第一个 break 做了什么？

```go
package main

func main() {
	for {
	switcher:
		switch 1 {
		case 1:
			switch 2 {
			case 2:
				break switcher
			}
		}
		break
	}
}
```

1. 它跳出第二个 switch，使程序无限循环
2. 它跳出第二个 switch，然后第二个 break 结束循环
3. **它跳出第一个 switch，然后第二个 break 结束循环（正确）**

**解释：**

* `break switcher` 作用于 **被 switcher 标签标记的 switch（第一个 switch）**
* 跳出第一个 switch 后，下一个语句就是 `break`，终止 for 循环
* 因此程序立即正常退出，不会无限循环

---

## 📌 6. 下列程序哪里有问题？

```go
package main

func main() {
	for {
	switcher:
		switch {
		case true:
			switch {
			case false:
				continue switcher
			}
		}
	}
}
```

1. **continue 只能继续循环，而不能继续 switch（正确）**
2. continue 不能用于 switch
3. 程序会无限循环

**解释：**

* `continue switcher` 期望跳转到一个循环标签，但 `switcher` 用在了 switch 上
* continue 后面必须是一个循环标签，否则编译报错

---

## 📌 7. 下列哪个程序会正常终止？

### 代码 1

```go
func main() {
start: goto exit
exit : fmt.Println("exiting")
goto start
}
```

### 代码 2

```go
func main() {
exit: fmt.Println("exiting")
goto exit
}
```

### 代码 3（正确）

```go
func main() {
    goto exit
start : goto getout
exit  : goto start
getout: fmt.Println("exiting")
}
```

**解释：**

* 代码 1：`start → exit → start` 无限循环
* 代码 2：`exit → exit` 无限循环
* **代码 3：最终跳到 getout 执行打印并结束 main，因此程序终止**

---
