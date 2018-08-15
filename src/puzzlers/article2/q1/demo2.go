package main

import (
	"flag"
	"fmt"
)

func main() {
	// 方式1： 读取参数直接赋值给变量
	var foo string
	flag.StringVar(&foo, "foo", "everyone", "The greeting object.")

	// 方式2：读取参数返回指针
	var bar *string
	bar = flag.String("bar", "everyone", "The greeting object.")

	flag.Parse()
	fmt.Printf("Hello, %s!\n", foo)
	fmt.Printf("Hello, %s!\n", *bar)
}
