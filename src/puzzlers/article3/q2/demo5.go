package main

import (
	"flag"
	// "puzzlers/article3/q2/lib" // 代码包相对src的路径

	in "puzzlers/article3/q4/lib/internal"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	// lib5.Hello(name) // 这里是代码包中package定义的lib5
	// 这是一个bad case的示例
	// 实际情况中应尽量避免包名与目录名不一致的情况

	in.Hello(name)
}
