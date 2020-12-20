package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

func TestDog(t *testing.T) {
	// dog := new(Dog)
	// dog.SpeakTo("chao")
	// dog.Speak()

	//不能按照继承的来用
	// var dog Pet = new(Dog)

	/*
		//也无法支持（强制转换）LSP？
		var dog *Dog = new(Dog)
		var p = (*Pet)(dog)
		p.SpeakTo("Chao")
		dog.SpeakTo("Chao")
	*/

	dog := new(Dog)
	dog.SpeakTo("Chao")

}
