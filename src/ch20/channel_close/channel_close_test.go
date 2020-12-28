package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//关闭channel
		close(ch)
		//给已关闭的通道中继续输入，会报错 panic
		// ch <- 11
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {

		for {
			// ok会返回channel状态，false的话，即channel关闭。无须再和生产者规定10
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}

		/*
			// i为11的时候，通道已经关闭，得到该类型的0值
			for i := 0; i < 11; i++ {
				data := <-ch
				fmt.Println(data)
			}
		*/

		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	//可以采用多个receiver
	// wg.Add(1)
	// dataReceiver(ch, &wg)
	wg.Wait()
}
