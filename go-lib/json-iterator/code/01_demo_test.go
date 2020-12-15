package code

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

const jsonByte = `{"id":1,"Name":"Reds","Colors":["Crimson", "Red", "Ruby", "Maroon"]}`

func TestDemo(t *testing.T) {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := jsoniter.Marshal(group)
	t.Logf("bytes %v", b)

	if err != nil {
		t.Error(err)
	}
	t.Log("------------split-----------")
	if err = jsoniter.Unmarshal([]byte(jsonByte), &group); err != nil {
		t.Error(err)
	}
	t.Logf("group %v", group)
}
