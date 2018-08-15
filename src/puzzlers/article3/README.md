# 03 | 库源码文件

摘要笔记:

`go build`相关：

- `go build xxxx`时，依次从`GOROOT`、`GOPATH`找此代码包
- 如果是`go build`，则缺省当前目录下的`main`包
- 通过`go build -x`看到完整构建过程日志
- 将任意目录的代码包加入`GOPATH`后，则可以再任意目录就行此代码包的构建
- `go run a.go b.go`：如果a->b有依赖时，及时同目录下run时也需要完整文件列表（`go run`不自动关联依赖文件，构建过程考虑）
  - 同级目录时`go run`不会自动查找依赖定义：`q1`中执行`go run demo4.go`时报错
  - 但如果在子包时，`go run`会自动查找依赖：`q2或q4`中执行`go run demox.go`均正常

重点：

- 同目录下的源码文件的代码包声明语句要一致
- 生成目标文件缺省与父目录名称一致
- **代码包导入路径**：源码文件所在的目录相对于 src 目录的相对路径
- 实际使用代码包时给定的限定符要与其声明所属的代码包名称对应
        - `import`需与文件目录名一致
        - 限定符(`xxx.`形式包名前缀)应与package声明保持一致
- 为了不让该代码包的使用者产生困惑，我们总是应该让声明的包名与其父目录的名称一致
- 访问权限：包级私有、公共、模块级私有
- `internal`模块级私有，更严格限制仅限父包、子包引用

## q1解析

### 将`Golang_Puzzlers`目录未加入到GOPATH时

- 进入`src\puzzlers\article3\q1`目录
- 执行`go build demo4.go demo4_lib.go`，正常运行没问题
- 因为当前目录下都可以找到这两个文件，并不依赖GOPATH

但，如果在任意目录执行`go build puzzlers/article3/q1`，则必然找不到这个包，因为`go build`是从GOPATH目录找到这个包进行编译的。
错误信息如下：

```go
$ cd ~
$ go build puzzlers/article3/q1
can't load package: package puzzlers/article3/q1: cannot find package "puzzlers/article3/q1" in any of:
        C:\Go\src\puzzlers\article3\q1 (from $GOROOT)
        C:\Users\fivezh\go\src\puzzlers\article3\q1 (from $GOPATH)
```

### 将`Golang_Puzzlers`目录加入到GOPATH后

- 进入`src\puzzlers\article3\q1`目录，执行`go build`都能成功编译
- 任意目录执行`go build puzzlers/article3/q1`，成功编译，并将可执行程序拷贝至当前目录

> 该项目的打包文件下载到本地的任意目录下，然后经解压缩后把“Golang_Puzzlers”目录加入到环境变量GOPATH

```go
$ echo $GOPATH
C:\Users\fivezh\go;C:\cygwin64\home\fivezh\github.com\Golang_Puzzlers;

// 当前目录go build
~/github.com/Golang_Puzzlers/src/puzzlers/article3/q1
$ go build

~/github.com/Golang_Puzzlers/src/puzzlers/article3/q1
$ ls
demo4.go  demo4_lib.go  q1.exe

// 任意目录go build，当前目录生成可执行程序
$ cd ~
$ go build puzzlers/article3/q1
$ ls
fivezh.github.io  github.com  go  q1.exe
```

## q2解析

- 包名与目录名不一致时，增加理解成本
  - 引入包名是目录相对src的路径名
  - 限定符是源码定义的`package name`名称
  - 如果二者不一致，就会出现q2中说的情况，应避免不一致

## q4解析

- internal模块级私有的访问权限实例
- 在其他包中`import internal`包会报错：`demo5.go:7:2: use of internal package not allowed`
- `internal`可以看成是更严格的访问权限限制