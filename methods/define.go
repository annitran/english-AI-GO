package methods

import "fmt"

// định nghĩa struct
type Person struct {
	Name   string
	Height int
}

// định nghĩa methods
func (p Person) ShowDetail() {
	fmt.Println("\nTên:", p.Name, "Chiều cao:", p.Height)
}
