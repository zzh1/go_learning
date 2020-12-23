package series

import (
	"fmt"
)

//可以定义多个init方法
func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}

//小写的声明是不能被访问的
// func square(n int) int {
func Square(n int) int {
	return n * n
}

func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-1]+ret[i-2])
	}
	return ret
}
