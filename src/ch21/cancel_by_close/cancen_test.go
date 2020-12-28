package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel_2(cancenChan chan struct{}) {
	close(cancenChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	// time.Sleep(time.Second * 3)
	//这里会随机向一个发送消息(发送消息可导致isCancelled返回true)，如果未加上面等待时间，一直为第4个。加上时间后，可关闭其他的协程
	cancel_1(cancelChan)
	time.Sleep(time.Second * 5)
}
