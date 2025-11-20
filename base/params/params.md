## Go 语言的参数传递方式

Go 语言支持**值传递**和**指针传递**两种方式。这里需要澄清一个重要概念。

### 1. Go 语言只支持值传递

这是 Go 语言的一个重要特性：**Go 只有值传递，没有引用传递**。

- 当你传递一个变量时，Go 会复制该变量的值
- 即使看起来像"引用传递"，实际上也是值传递

### 2. 数组 vs 切片的传递区别

#### 数组传递（值传递 - 整个数组被复制）
```go
func ProcessArray(arr [5]int) {
    arr[0] = 999  // 只修改复制的副本
}

func main() {
    original := [5]int{1, 2, 3, 4, 5}
    ProcessArray(original)
    fmt.Println(original[0])  // 输出: 1 (未改变)
}
```

**特点**：
- 整个数组被复制
- 函数内的修改不会影响原数组
- 如果数组很大，会消耗大量内存和 CPU

#### 切片传递（看似"引用" - 但实际上是值传递）
```go
func ProcessSlice(slice []int) {
    slice[0] = 999  // 修改底层数组的数据
}

func main() {
    original := []int{1, 2, 3, 4, 5}
    ProcessSlice(original)
    fmt.Println(original[0])  // 输出: 999 (改变了!)
}
```

**特点**：
- 切片本身被复制（但切片很小，只是指针、长度、容量）
- 切片指向的底层数组没有被复制
- 函数内的修改会影响原数据
- 更高效，因为只复制了指针信息

### 3. 切片的内部结构

```go
// 切片内部结构（约24字节）
type SliceHeader struct {
    Data uintptr  // 指向底层数组的指针
    Len  int      // 长度
    Cap  int      // 容量
}
```

当传递切片时，只复制这个小的结构体，而不是整个数据。

### 4. 指针传递

如果你想修改接收者本身（比如重新分配指针指向的对象），需要使用指针：

```go
func ReallocateSlice(slicePtr *[]int) {
    *slicePtr = append(*slicePtr, 100)  // 修改切片本身
}

func main() {
    original := []int{1, 2, 3}
    ReallocateSlice(&original)
    fmt.Println(original)  // 输出: [1 2 3 100]
}
```

### 5. Go vs Java 的比较

#### Java
- 只有值传递
- 传递对象时，传递的是对象引用的副本（引用地址值）
- 原始类型（int、double等）传递的是值本身

```java
void modifyArray(int[] arr) {
    arr[0] = 999;  // 修改原数组（因为 arr 是引用副本）
}
```

#### Go
- 只有值传递
- 传递数组时，整个数组被复制
- 传递切片时，切片结构体被复制（但数据共享）
- 传递指针时，指针值被复制

```go
// 传递数组 - 整个数组复制
func modifyArray(arr [5]int) {
    arr[0] = 999  // 只修改副本
}

// 传递切片 - 切片结构体复制，但指向同一数据
func modifySlice(s []int) {
    s[0] = 999  // 修改原数据
}

// 传递指针 - 指针值复制
func modifyViaPointer(p *[5]int) {
    p[0] = 999  // 修改原数组
}
```

#### 相似点
- 两者都只支持值传递
- 传递的是某种形式的"引用"（指向数据的指针）
- 函数内的修改会影响原数据

#### 不同点
| 特性 | Java | Go |
|------|------|-----|
| 传递对象 | 传递引用副本 | 传递对象本身（复制） |
| 传递数组 | 传递引用副本 | 传递数组副本（大量复制！） |
| 传递切片 | 不适用 | 传递切片结构体副本（24字节） |
| 修改对象 | 指向同一对象 | 指向同一数据 |


### 6. 为什么切片比数组更高效

```
数组传递 [5]int:
┌─────────────────────────┐
│ 复制 5 个整数            │ 较大的复制操作
└─────────────────────────┘

切片传递 []int:
┌──────┬────┬────┐
│ Ptr  │ Len│Cap │ 仅复制 24 字节
└──────┴────┴────┘
  ↓
 指向原始数据（不复制）
```

### 7. 最佳实践

1. **优先使用切片而不是数组**
   ```go
   func Sum(numbers []int) int { ... }  // ✅ 推荐
   func Sum(numbers [5]int) int { ... } // ❌ 不推荐
   ```

2. **大结构体使用指针**
   ```go
   type LargeStruct struct {
       data [1000]int
   }
   func Process(s *LargeStruct) { ... }  // ✅ 使用指针避免复制
   ```

3. **小结构体可以值传递**
   ```go
   type Point struct {
       X, Y int
   }
   func Distance(p Point) { ... }  // ✅ 结构体小，值传递可以接受
   ```

4. **需要修改接收者时使用指针**
   ```go
   type Counter struct {
       count int
   }
   func (c *Counter) Increment() {  // 使用指针接收器
       c.count++
   }
   ```
