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

	s = "中"
	t.Log(len(s)) //是byte数

	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c) //[1]表示都匹配第一个参数
	}
}
