package greetings

import "fmt"

// 实现一个函数来返回问候语
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
