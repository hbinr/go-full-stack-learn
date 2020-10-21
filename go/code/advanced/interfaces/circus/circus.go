package circus

type Speaker interface {
	Speaks() string
}

func Perform(a Speaker) string {
	return a.Speaks()
}
