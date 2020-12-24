package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		//在协程中执行
		go func(i int) {
			fmt.Println(i)
		}(i) //i是值传递

		/*
			//这种形式，参数i在test协程和创建的协程中被共享了
			go func() {
				fmt.Println(i)
			}()
		*/

	}
	time.Sleep(time.Millisecond * 50)
}
