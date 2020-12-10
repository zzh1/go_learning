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
