package string

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) //初始化为默认零值 “”
	s = "hello"
	t.Log(len(s))
	// s[1] = "3" //string 是不可变的 byte slice
	s = "\xE4\xBB\xA5"
	// s = "\xE4\xBA\xBB\xFF" //可以存储任何二进制数据
	t.Log(s)
	t.Log(len(s))
}
