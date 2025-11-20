# Go 语言学习笔记

## 目录

### 基础

- [hello 示例补充说明](base/hello/readme.md)
- 整数
  - [示例代码](base/arrays/sum_test.go)
  - [官网教程](https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/integers)
- 数字和字符串
  - [示例代码](/base/numbers-and-strings/)
  - [常见问题](/base/numbers-and-strings/readme.md)
  - 相关库教程文档
    - [strconv](/base/numbers-and-strings/strings/strconv.md)
    - [utf8](/base/numbers-and-strings/strings/utf8.md)
  - 字符串、Rune 和字节测验
    - [示例代码](base/strings-runes-bytes/)
    - [常见问题](base/strings-runes-bytes/questions/README.md)
- Go 类型系统
  - [类型系统基础](base/go-type-system/types.md)
  - [位与字节](base/go-type-system/bites.md)
  - [整数溢出](base/go-type-system/overflow.md)
    - 示例代码：[问题演示](base/go-type-system/04-overflow/01-problem/main.go)、[原理解释](base/go-type-system/04-overflow/02-explain/main.go)、[数据丢失](base/go-type-system/04-overflow/03-destructive/main.go)
    - 包含：有符号/无符号整数溢出、浮点数溢出、编译时/运行时检查、避免方法、最佳实践
  - [自定义类型（Defined Types）](base/go-type-system/defined-types.md)
    - 示例代码：[Duration 示例](base/go-type-system/05-defined-types/01-duration-example/main.go)、[类型定义](base/go-type-system/05-defined-types/02-type-definition-create-your-own-type/main.go)、[底层类型](base/go-type-system/05-defined-types/03-underlying-types/main.go)
    - 包含：类型定义语法、底层类型概念、类型转换规则、跨包类型、为类型添加方法、类型安全、常见陷阱、最佳实践
  - [类型别名（Aliased Types）](base/go-type-system/aliased.md)
  - [Go vs Java 编码方式](base/go-type-system/go-vs-java-encoding.md)
  - 常见问题
    - [常见问题1](base/go-type-system/problem1.md)
    - [常见问题2](base/go-type-system/problem2.md)
- 常量
  - [示例代码](base/constants/01-declarations/main.go)
  - [常量教程](base/constants/readme.md)
  - [常见问题](base/constants/problem.md)
- if
  - [示例代码](base/if/01-boolean-operators/)
  - [常见问题](base/if/problem.md)
- switch
  - [示例代码](base/switch/01-one-case/)
  - [常见问题](base/switch/questions/questions.md)
- loops
  - [示例代码](base/loops/)
  - [常见问题](base/loops/questions/)
- 包和作用域（Package）
  - [示例代码](base/package-scopes)
  - [完整教程](base/package-scopes/package.md)
  - 包含：包的基本概念、作用、命名规则、组织结构、导入、导出与非导出、常见问题
- 声明-表达式-注释
  - [示例代码](base/expressions-comments/main.go)
  - [教程](base/expressions-comments/readme.md)
- 循环
  - [示例代码](base/for/reapeat_test.go) 
  - [官网教程](https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/iteration)
- 数组
  - [示例代码](base/integers/adder_test.go)
  - [示例补充说明](base/arrays/readme.md)
  - [常见问题](base/arrays/questions/README.md)
- 切片
  - [示例代码](base/slices/)
  - [常见问题](base/slices/questions/)
- 参数传递的方式
  - [示例教程](base/params/params.md)
- fmt-打印格式化
  - [示例教程](/base/fmt-guide.md)
  - [官方文档](https://pkg.go.dev/fmt)
- 结构体-struts，方法和接口
  - [示例代码](base/structs/shapes_test.go)
  - [官网教程](https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/structs-methods-and-interfaces)
  - [常见问题](base/structs/questions/README.md)
- func
  - [示例代码](base/functions/01-basics/)
  - [常见问题](base/functions/questions/README.md)
  - [方法声明的方式详细说明](base/functions/readme.md)
- [接口](https://go.dev/ref/spec#Interface_types)
- 指针(pointer)和错误
  - [示例代码](base/pointers/wallet_test.go)
  - [常见问题](base/pointers/questions/README.md)
  - [教程](https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/pointers-and-errors)
- Maps
  - [示例代码](base/maps/dictionary.go)
  - [示例补充](base/maps/readme.md)
  - [常见问题](base/maps/questions/README.md)
  - [教程](https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/maps)
- 输入流
  - [示例代码](base/input-scanning/)
  - [常见问题](base/input-scanning/questions/README.md)
- 初始化（Initialization）
  - [完整教程](base/init/init.md)
  - 包含：`{}` 统一初始化语法、复合类型初始化、bytes.Buffer、最佳实践
- 依赖注入
  - [示例代码](base/di/di.go)
  - [教程](https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/dependency-injection)
- Mocking
  - [示例代码](base/mocking/)
  - [教程](https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/mocking)
- 并发（Concurrency）
  - [示例代码](base/concurrency/check_websites.go)
  - [完整教程](base/concurrency/concurrency.md)
  - 包含：Goroutine、Channel、并发模式、代码详解、最佳实践
  - [官网教程](https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/concurrency)
- select
  - [示例代码](base/select/racer.go)
  - [示例教程补充](base/select/select.md)
  - [完整教程](https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/select)
- 反射
  - [示例代码](base/reflection/)
  - [示例教程补充](base/reflection/reflection.md)
  - [完整教程](https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/reflection)
- 导入远程库包
  - [示例代码](base/library-package/cmd/main.go)
  - [示例教程](base/library-package/readme.md)

### 构建应用程序

- HTTP服务器
  - [示例代码](/build-app/http-server)
  - [示例代码补充](/build-app/http-server/http-server.md)
  - [教程](https://studygolang.gitbook.io/learn-go-with-tests/gou-jian-ying-yong-cheng-xu/http-server)

- JSON，路由和嵌入
  - [示例代码](/build-app/json-route/main.go)
  - [示例讲解](/build-app/json-route/json.md)
  - [官网教程](https://studygolang.gitbook.io/learn-go-with-tests/gou-jian-ying-yong-cheng-xu/json)

- IO 和排序
  - [示例代码](/build-app/io/)
  - [示例代码补充](/build-app/io/io.md)
  - [官网教程](https://studygolang.gitbook.io/learn-go-with-tests/gou-jian-ying-yong-cheng-xu/io)

- 复古led时钟项目
  - [示例项目](/build-app/project-retro-led-clock/)
  - [常见问题](/build-app/project-retro-led-clock/README.md)

- 项目空文件查找
  - [示例项目](build-app/project-empty-file-finder/)

- 弹力球
  - [示例项目](build-app/project-bouncing-ball/)

- 垃圾信息屏蔽挑战技巧
  - [示例项目](build-app/project-spam-masker)

- 文本换行挑战指南
  - [示例项目](build-app/project-text-wrapper/)