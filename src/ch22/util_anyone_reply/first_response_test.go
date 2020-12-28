package util_anyone_reply

import (
	"fmt"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	//channel类型的ch会阻塞在这里，当从ch中取到值后，直接返回。实现了 一获取到就返回的效果
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log(FirstResponse())
}
