package csp_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// chan string 表示返回通道类型，通道的类型是string
func AsyncService() chan string {
	//普通channal
	// retCh := make(chan string)
	//buffer channel,后面的1表示缓冲容量
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		//往channel中放数据
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	// fmt.Println(<-AsyncService())
	// time.Sleep(time.Second * 1)
}
