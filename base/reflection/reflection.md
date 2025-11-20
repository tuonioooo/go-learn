# Go 反射完整教程：从入门到实践

## 什么是反射？

反射是程序检查自身结构的能力，特别是通过类型来检查。它是元编程的一种形式，允许程序在运行时动态地操作对象。

在 Go 中，反射主要通过 `reflect` 包来实现，它提供了在运行时检查变量类型和值的能力。

## 为什么需要反射？

### 实际场景

假设我们需要编写一个函数 `walk(x interface{}, fn func(string))`，它能够：
- 接收任意类型的结构体 `x`
- 遍历结构体中的所有字符串字段
- 对每个字符串字段调用函数 `fn`

这种情况下，我们在编译时不知道具体的类型，就需要使用反射。

### interface{} 类型

Go 使用 `interface{}` 类型来处理未知类型的参数，它可以接收任何类型的值。但这也带来了一些问题：

**缺点：**
- 失去了类型安全检查
- 编译器无法提前发现类型错误
- 代码可读性降低
- 性能开销较大

**什么时候使用反射？**
- 只在真正需要时使用反射
- 能用接口（interface）解决的问题不要用反射
- 优先考虑类型安全的设计

## walk 函数详解

### walk 函数的核心作用

`walk(x interface{}, fn func(string))` 函数是一个**通用的数据遍历工具**，它的作用是：

1. **接收任意类型的数据** (`x interface{}`)
2. **递归遍历该数据的所有结构** （使用反射）
3. **收集所有字符串类型的值** （只提取 string 字段）
4. **对每个字符串值执行回调函数** （调用 `fn(string)` 进行处理）

### 完整执行流程

```
walk(test.Input, func(input string) {
    got = append(got, input)
})

                    ↓

walk 函数【内部】使用反射处理 test.Input：
    ├─ 遍历 test.Input 的结构
    ├─ 找到所有 string 类型的值
    ├─ 调用回调函数
    └─ fn("Chris") → got = append(got, "Chris")
    └─ fn("London") → got = append(got, "London")
    
最后 got = ["Chris", "London"]
```

**关键点**：
- `test.Input` 是**被处理的数据源**，被传入 `walk` 函数
- `walk` 函数**内部**使用反射来遍历和提取数据（你看不见）
- `func(input string) { got = append(got, input) }` 是**回调函数**，接收提取的数据
- 虽然测试代码里没有显式使用 `test.Input`，但它已经被 `walk` 函数消费了

### 类比理解

```go
// 简单的类比
myList := []string{"Alice", "Bob", "Charlie"}
forEach(myList, func(name string) {
    fmt.Println(name)
})

// forEach 函数内部：
func forEach(items []string, fn func(string)) {
    for _, item := range items {  // ← 这里使用了 myList
        fn(item)
    }
}

// 同理，walk 函数内部也是用反射遍历 test.Input
```

### walk 与普通函数的区别

| 方面 | 普通函数 | walk 函数 |
|------|---------|---------|
| **输入类型** | 确定的（如 `[]string`) | 任意的（`interface{}`） |
| **类型检查** | 编译时检查 | 运行时反射检查 |
| **适应性** | 只能处理一种类型 | 能处理多种类型结构 |
| **代码复杂度** | 简单 | 使用反射，较复杂 |

## 实战：实现 walk 函数

让我们通过 TDD（测试驱动开发）的方式，逐步实现一个完整的 `walk` 函数。

### 第一步：处理简单结构体

#### 测试代码

```go
func TestWalk(t *testing.T) {
    expected := "Chris"
    var got []string
    
    x := struct {
        Name string
    }{expected}
    
    walk(x, func(input string) {
        got = append(got, input)
    })
    
    if len(got) != 1 {
        t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
    }
    
    if got[0] != expected {
        t.Errorf("got '%s', want '%s'", got[0], expected)
    }
}
```

#### 实现代码

```go
func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)
    field := val.Field(0)
    fn(field.String())
}
```

**关键概念：**
- `reflect.ValueOf(x)`：获取变量的反射值
- `val.Field(0)`：获取第一个字段
- `field.String()`：将字段值转换为字符串

### 第二步：处理多个字段

```go
func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)
    
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fn(field.String())
    }
}
```

**新增方法：**
- `val.NumField()`：获取结构体字段数量
- 使用 for 循环遍历所有字段

### 第三步：类型检查

并非所有字段都是字符串，我们需要类型检查：

```go
func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)
    
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        
        if field.Kind() == reflect.String {
            fn(field.String())
        }
    }
}
```

**关键概念：**
- `field.Kind()`：获取字段的类型种类
- `reflect.String`：字符串类型常量

### 第四步：处理嵌套结构体

使用递归处理嵌套的结构体：

```go
func walk(x interface{}, fn func(input string)) {
    val := reflect.ValueOf(x)
    
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        
        switch field.Kind() {
        case reflect.String:
            fn(field.String())
        case reflect.Struct:
            walk(field.Interface(), fn)  // 递归调用
        }
    }
}
```

**示例结构体：**

```go
type Person struct {
    Name    string
    Profile Profile
}

type Profile struct {
    Age  int
    City string
}
```

### 第五步：处理指针

指针类型需要特殊处理，使用 `Elem()` 获取底层值：

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)
    
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        
        switch field.Kind() {
        case reflect.String:
            fn(field.String())
        case reflect.Struct:
            walk(field.Interface(), fn)
        }
    }
}

func getValue(x interface{}) reflect.Value {
    val := reflect.ValueOf(x)
    
    if val.Kind() == reflect.Ptr {
        val = val.Elem()  // 获取指针指向的值
    }
    
    return val
}
```

### 第六步：处理切片和数组

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)
    
    walkValue := func(value reflect.Value) {
        walk(value.Interface(), fn)
    }
    
    switch val.Kind() {
    case reflect.String:
        fn(val.String())
    case reflect.Struct:
        for i := 0; i < val.NumField(); i++ {
            walkValue(val.Field(i))
        }
    case reflect.Slice, reflect.Array:
        for i := 0; i < val.Len(); i++ {
            walkValue(val.Index(i))
        }
    }
}
```

**关键方法：**
- `val.Len()`：获取切片/数组长度
- `val.Index(i)`：获取索引位置的元素

### 第七步：处理 Map

```go
func walk(x interface{}, fn func(input string)) {
    val := getValue(x)
    
    walkValue := func(value reflect.Value) {
        walk(value.Interface(), fn)
    }
    
    switch val.Kind() {
    case reflect.String:
        fn(val.String())
    case reflect.Struct:
        for i := 0; i < val.NumField(); i++ {
            walkValue(val.Field(i))
        }
    case reflect.Slice, reflect.Array:
        for i := 0; i < val.Len(); i++ {
            walkValue(val.Index(i))
        }
    case reflect.Map:
        for _, key := range val.MapKeys() {
            walkValue(val.MapIndex(key))
        }
    }
}
```

**Map 相关方法：**
- `val.MapKeys()`：获取 map 的所有键
- `val.MapIndex(key)`：通过键获取值

## 完整测试用例

```go
func TestWalk(t *testing.T) {
    cases := []struct {
        Name          string
        Input         interface{}
        ExpectedCalls []string
    }{
        {
            "Struct with one string field",
            struct {
                Name string
            }{"Chris"},
            []string{"Chris"},
        },
        {
            "Struct with two string fields",
            struct {
                Name string
                City string
            }{"Chris", "London"},
            []string{"Chris", "London"},
        },
        {
            "Struct with non string field",
            struct {
                Name string
                Age  int
            }{"Chris", 33},
            []string{"Chris"},
        },
        {
            "Nested fields",
            Person{
                "Chris",
                Profile{33, "London"},
            },
            []string{"Chris", "London"},
        },
        {
            "Pointers to things",
            &Person{
                "Chris",
                Profile{33, "London"},
            },
            []string{"Chris", "London"},
        },
        {
            "Slices",
            []Profile{
                {33, "London"},
                {34, "Reykjavík"},
            },
            []string{"London", "Reykjavík"},
        },
        {
            "Arrays",
            [2]Profile{
                {33, "London"},
                {34, "Reykjavík"},
            },
            []string{"London", "Reykjavík"},
        },
    }
    
    for _, test := range cases {
        t.Run(test.Name, func(t *testing.T) {
            var got []string
            walk(test.Input, func(input string) {
                got = append(got, input)
            })
            
            if !reflect.DeepEqual(got, test.ExpectedCalls) {
                t.Errorf("got %v, want %v", got, test.ExpectedCalls)
            }
        })
    }
}
```

## reflect 包核心 API

### reflect.ValueOf

获取变量的反射值对象：

```go
val := reflect.ValueOf(x)
```

### reflect.Value 常用方法

| 方法 | 说明 | 适用类型 |
|------|------|----------|
| `Kind()` | 获取值的种类 | 所有类型 |
| `NumField()` | 获取字段数量 | Struct |
| `Field(i)` | 获取第 i 个字段 | Struct |
| `Len()` | 获取长度 | Slice, Array, Map, String |
| `Index(i)` | 获取索引元素 | Slice, Array |
| `MapKeys()` | 获取 map 所有键 | Map |
| `MapIndex(key)` | 通过键获取值 | Map |
| `Elem()` | 获取指针指向的值 | Ptr, Interface |
| `Interface()` | 转回 interface{} | 所有类型 |
| `String()` | 转为字符串 | String |

### reflect.Kind 常见类型

```go
reflect.String
reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64
reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64
reflect.Float32, reflect.Float64
reflect.Bool
reflect.Struct
reflect.Slice
reflect.Array
reflect.Map
reflect.Ptr
reflect.Interface
```

## 使用建议

### ✅ 适合使用反射的场景

1. **序列化/反序列化库**（如 JSON、XML）
2. **ORM 框架**（数据库映射）
3. **依赖注入框架**
4. **通用数据处理工具**
5. **RPC 框架**

### ❌ 应该避免的场景

1. 可以用接口解决的问题
2. 已知类型的操作
3. 性能敏感的代码
4. 简单的类型转换

### 最佳实践

1. **优先使用类型安全的设计**
   ```go
   // 好的设计
   type Stringer interface {
       String() string
   }
   
   func process(s Stringer) {
       fmt.Println(s.String())
   }
   ```

2. **添加充分的错误检查**
   ```go
   val := reflect.ValueOf(x)
   if val.Kind() != reflect.Struct {
       return fmt.Errorf("expected struct, got %v", val.Kind())
   }
   ```

3. **封装反射逻辑**
   - 将反射代码封装在独立的函数中
   - 提供清晰的 API 给使用者

4. **性能考虑**
   - 反射操作比直接操作慢 10-100 倍
   - 考虑缓存反射信息
   - 在非热点路径使用

## 总结

反射是 Go 语言中强大但需要谨慎使用的特性：

- **优点**：灵活性、通用性、元编程能力
- **缺点**：性能损失、类型安全丧失、代码复杂度增加

**核心原则**：只在真正需要时使用反射，能用接口解决的问题就不要用反射。

## 延伸阅读

- [Go 官方博客：反射三定律](https://blog.golang.org/laws-of-reflection)
- [reflect 包官方文档](https://pkg.go.dev/reflect)
- 通过 TDD 实践学习更多 Go 知识

---

*本教程基于 TDD 方法论，通过逐步迭代的方式深入理解 Go 反射机制。*