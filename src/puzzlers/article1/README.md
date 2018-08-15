# 01 | 工作区和GOPATH

概念：

- 工作区(workspace)：
    - bin目录：编译安装后的可执行程序
    - pkg目录：库文件安装后的归档目录
    - src目录：源码文件
- GOPATH：
    - 一个或多个目录路径
    - 每个目录路径代表一个workspace工作区
- GOROOT：
    - Go安装程序的目录，即下载官方go后解压的目录
- GOBIN：
    - bo install时将可执行程序安装的目录

新发现之前遗漏知识点：
- 哪些安装到pkg，哪些安装到bin？
    - 如果定义了main()函数，则为主程序入口，`go install`后将可执行程序复制到bin目录
    - 如果未定义main()函数，则安装后复制至pkg目录

## 思考题

1. Go 语言在多个工作区中查找依赖包的时候是以怎样的顺序进行的？
    - 个人认为：从GOPATH中首个工作区开始查找依赖包，逐个向后查找
2. 如果在多个工作区中都存在导入路径相同的代码包会产生冲突吗？
    - 个人认为：按GOPATH顺序逐个匹配，首个找到的

貌似从未用过多工作区的情况，对这个理解有偏差  
`go help gopath`获取gopath的文档


## 如何获取输入参数？

- 使用`flag`包进行输入参数获取
- flag推荐的参数方式`-foo fooValue`，同时也兼容`-foo=fooValue`或`--foo=fooValue`形式

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	var foo string
	flag.StringVar(&foo, "foo", "everyone", "The greeting object.")
	var bar *string
	bar = flag.String("bar", "everyone", "The greeting object.")
    flag.Parse()
	fmt.Printf("Hello, %s!\n", foo)
	fmt.Printf("Hello, %s!\n", *bar)
}

// 输出
$ go run demo2.go -name2
flag provided but not defined: -name2
Usage of Temp\go-build355183790\b001\exe\demo2.exe:
  -bar string
        The greeting object. (default "everyone")
  -foo string
        The greeting object. (default "everyone")
exit status 2

$ go run demo2.go -foo hello -bar world
Hello, hello!
Hello, world!
```

## 参考

- [go get学习笔记](https://github.com/hyper0x/go_command_tutorial/blob/master/0.3.md)