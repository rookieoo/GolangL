package main

// import (
// 	"fmt"

// 添加对 golang.org/x/example 模块的依赖关系

// 	"golang.org/x/example/stringutil"
// )

// func main() {
// 	fmt.Println(stringutil.Reverse("Hello"))
// }
import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.ToUpper("Hello"))
}

// 创建工作区  go work init ./hello
// 将新模块添加到工作区 go work use ./example
