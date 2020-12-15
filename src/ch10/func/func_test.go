package fn_test

import (
	"math/rand"
	"testing"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	// a, b := returnMultiValues()
	//不需要某个参数，可以使用占位符
	a, _ := returnMultiValues()
	t.Log(a)
}
