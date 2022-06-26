package main

// 为代码启用依赖项跟踪 go mod init example/hello
import (
	"fmt"

	"rsc.io/quote"
)

// 在系统环境中设置 - 配置go在vsocode中的开发环境
// go env -w GOPROXY=https://goproxy.io,direct
func main() {
	fmt.Println("Hello World")
	// 外部包的调用代码 - 添加新的模块 go mod tidy

	fmt.Println(quote.Go())
}

// 您的代码调用该函数，并打印有关通信的智能消息
