package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//断言的方式， p.()   即p如果是int类型
	if i, ok := p.(int); ok {
		fmt.Println("Interger", i)
		return
	}

	if s, ok := p.(string); ok {
		fmt.Println("string", s)
		return
	}

	fmt.Println("Unknow Type")
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
