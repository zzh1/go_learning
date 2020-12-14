package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	//第2各参数意思是cap，但map无法获取cap
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
	// t.Logf("len m3=%d", cap(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	//会返回两个变量，一个为值，一个为布尔。布尔为true，值存在；
	//在访问的key不存在是，仍然返回零值，不能通过返回nil来判断元素是否存在
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
		t.Log(ok)
	} else {
		t.Log("Key 3 is not existing.")
	}

}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
