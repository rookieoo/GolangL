package main

import "fmt"

type Number interface {
	int64 | float64
}

// 声明两个函数以将映射的值相加并返回总和
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// 1. 添加泛型函数以处理多个类型
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// 声明一个泛型函数，其逻辑与之前声明的泛型函数相同，但使用新的接口类型而不是联合作为类型约束
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// Initialize a map for the integer values
	// 添加非泛型函数
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		// 2. 调用泛型函数时删除类型参数
		// 当 Go 编译器可以推断出要使用的类型时，可以在调用代码中省略类型参数。编译器从函数参数的类型推断类型参数
		// SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}
