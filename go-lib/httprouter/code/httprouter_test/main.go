package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := httprouter.New()
	r.Handle(http.MethodGet, "/get/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "id: %v\n", params.ByName("id"))
	})

	r.Handle(http.MethodPost, "/post", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// 法一:
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			return
		}
		/*
			// 法二:
			b, err := ioutil.ReadAll(r.Body)
			fmt.Println("ioutil.ReadAll(r.Body)", r.Body)

			if err != nil {
				fmt.Println("ioutil.ReadAll(r.Body)", err)
				return
			}

			if err = json.Unmarshal(b, &app); err != nil {
				fmt.Println("json.Unmarshal(b, &uReq)", err)
				return
			}
		*/
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "post : %v\n", user)
	})
	http.ListenAndServe(":8081", r) //5040311
}
