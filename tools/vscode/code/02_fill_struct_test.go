package code

import (
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestFillStruct(t *testing.T) {
	u := &User{
		Name: "",
		Age:  0,
	}
	t.Log(u)
}
