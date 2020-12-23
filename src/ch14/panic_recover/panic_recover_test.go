package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	//panic后会调用defer函数
	// defer func() {
	// 	fmt.Println("Finally!")
	// }()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	// os.Exit(-1)
}
