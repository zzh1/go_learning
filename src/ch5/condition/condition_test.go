package condition_test

import (
	"testing"
)

func TestIfMultiSec(t *testing.T) {

	if v, err := someFun(); err == nil {
		t.Log("1==1")
	} else {
		t.Log("1==1")
	}

}
