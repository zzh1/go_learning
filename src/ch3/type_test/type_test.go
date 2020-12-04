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

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)

}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {

	}
}
