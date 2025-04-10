package methods

import "fmt"

// private
type person struct {
	Name string
	Age  int
}

// hàm khởi tạo
func NewPerson(name string, age int) person {
	return person{Name: name, Age: age}
}

// methods
func (p person) ChangeValue(newName string) {
	p.Name = newName
}
func (p *person) ChangePointer(newName string) {
	p.Name = newName
}

// Show
func (p person) ShowValue() {
	fmt.Println("Value:", p.Name, p.Age)
}

func (p *person) ShowPointer() {
	fmt.Println("Pointer:", p.Name, p.Age)
}
