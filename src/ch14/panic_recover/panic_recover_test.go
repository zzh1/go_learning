package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally!")
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	// os.Exit(-1)
}
