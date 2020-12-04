package type_test

import (
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {

	var a int = 1
	var b int64
	//不能隐式转换
	// b = a

	b = int64(a)
	var c MyInt
	//不能隐式转换
	// c = b
	c = MyInt(b)
	t.Log(a, b, c)

}
