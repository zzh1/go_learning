package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}

	//退出返回值
	// os.Exit(-1)
}
