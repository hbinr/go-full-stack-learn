package code

import (
	"context"
	"net/http"
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

type ColorRequest struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Colors []string `json:"colors"`
}
type ColorResponse struct {
	Data   interface{}
	ErrMsg error
}

func decodeProducerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req ColorRequest
	if err := jsoniter.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return jsoniter.NewEncoder(w).Encode(response)
}
