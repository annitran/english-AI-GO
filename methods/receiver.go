package methods

// vd1
func (p Person) ChangeHeight(height int) {
	p.Height = height
}

func (p *Person) ChangeHeight2(height int) {
	p.Height = height
}

// vd2
func (p *Person) SetName(name string) {
	p.Name = name
}

// Pointer receiver
func (p *Person) SetHeight(height int) {
	p.Height = height
}
