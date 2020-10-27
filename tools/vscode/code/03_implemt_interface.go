package code

type Speaker interface {
	// Speak speak action
	Speak()
}

type Student struct {
}

// Speak speak action
func (s *Student) Speak() {
	panic("not implemented") // TODO: Implement
}

type Teacher struct {
}

// Speak speak action
func (t *Teacher) Speak() {
	panic("not implemented") // TODO: Implement
}
