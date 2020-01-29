package extends_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("Wang !")
}
func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Chao") // 不支持重载， 不能当做继承来用
	dog.Speak()

}
