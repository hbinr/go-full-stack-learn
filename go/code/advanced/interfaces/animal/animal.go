package animal

type Dog struct{}

func (a *Dog) Speaks() string {
	return "woof"
}
