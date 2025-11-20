# Go 指针和方法说明

## 1. 方法接收器（Method Receiver）

在 Go 语言中，方法是一种特殊的函数，它与特定类型关联。方法有两种接收器类型：值接收器和指针接收器。

### 1.1 基本语法
```go
// 方法的一般形式
func (接收器 类型) 方法名(参数列表) 返回值 {
    // 方法体
}

// 具体示例
func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}
```

### 1.2 指针接收器 vs 值接收器

1. **指针接收器 `(w *Wallet)`**
   - 可以修改接收器的状态
   - 避免复制整个结构体
   - 适用于需要修改数据的方法
   ```go
   func (w *Wallet) Deposit(amount Bitcoin) {
       w.balance += amount  // 可以修改数据
   }
   ```

2. **值接收器 `(w Wallet)`**
   - 操作的是数据的副本
   - 不会修改原始数据
   - 适用于只读操作
   ```go
   func (w Wallet) Display() {
       fmt.Println(w.balance)  // 只读操作
   }
   ```

## 2. 方法的返回值

方法可以有返回值，也可以没有返回值，这与接收器类型无关：

### 2.1 无返回值方法
```go
func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}
```

### 2.2 有返回值方法
```go
func (w *Wallet) Balance() Bitcoin {
    return w.balance
}
```

### 2.3 返回错误的方法
```go
func (w *Wallet) Withdraw(amount Bitcoin) error {
    if amount > w.balance {
        return ErrInsufficientFunds
    }
    w.balance -= amount
    return nil
}
```

## 3. 方法调用

```go
wallet := Wallet{}           // 创建实例
wallet.Deposit(10)          // 调用方法
balance := wallet.Balance() // 获取余额
```

## 4. 使用指针接收器的原因

1. **修改数据**
   - 需要在方法中修改接收器的状态
   - 例如：存款、取款操作

2. **性能考虑**
   - 避免复制大型结构体
   - 节省内存使用

3. **一致性**
   - 如果类型的某些方法需要指针接收器，最好所有方法都使用指针接收器
   - 保持接口的一致性

## 5. 最佳实践

1. **选择接收器类型**
   - 需要修改状态时使用指针接收器
   - 只读操作可以使用值接收器
   - 大型结构体优先使用指针接收器

2. **方法命名**
   - 使用清晰、描述性的名称
   - 遵循 Go 的命名约定

3. **错误处理**
   - 需要报告错误的方法应返回 error 类型
   - 使用有意义的错误信息

## 6. 示例：完整的类型定义

```go
type Wallet struct {
    balance Bitcoin
}

// 存款方法：修改状态，使用指针接收器
func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}

// 余额方法：只读操作，但为了一致性也使用指针接收器
func (w *Wallet) Balance() Bitcoin {
    return w.balance
}

// 取款方法：修改状态并返回错误，使用指针接收器
func (w *Wallet) Withdraw(amount Bitcoin) error {
    if amount > w.balance {
        return ErrInsufficientFunds
    }
    w.balance -= amount
    return nil
}
```
