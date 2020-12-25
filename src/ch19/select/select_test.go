package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
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

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
