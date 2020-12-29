package obj_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	//尝试放置超出池大小的对象，现在是满的，再放一个
	if err := pool.ReusableObj(&ReusableObj{}); err != nil {
		t.Error(err)
	}
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			//拿出后就放回
			if err := pool.ReusableObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}
