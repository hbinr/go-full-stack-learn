package circus

import (
	"testing"

	"hb.study/go/code/advanced/interfaces/animal"
)

type Speaker interface {
	Speaks() string
}

func Perform(a Speaker) string {
	return a.Speaks()
}

func TestPerform(t *testing.T) {
	t.Logf("res:%s", Perform(new(animal.Dog)))
}
