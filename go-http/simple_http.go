package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// func handle(w http.ResponseWriter, req *http.Request) {
// 	w.Write(([]byte(`{"name": "KradasA4`)))
// }

type User struct {
	Id   int    `json:"id"`
	Name string `json:"nickname"`
	Age  int    `json:"age"`
}

var users = []User{
	{Id: 1, Name: "Kradas", Age: 12},
	{Id: 2, Name: "Noom", Age: 132},
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			log.Println("GET")
			b, err := json.Marshal(users)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(([]byte(err.Error())))
				return
			}

			w.Write(b)
			return
		}

		if req.Method == "POST" {
			log.Println("POST")
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			var u User
			err = json.Unmarshal(body, &u)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				fmt.Fprintf(w, "error: %v", err)
				return
			}

			users = append(users, u)
			// Fprintf call w.Write since we pass w as a writer
			fmt.Fprintf(w, "hello %s created users", "POST") // hello POST created users
			// res, err := json.Marshal("hello POST created users")
			// w.Write(res) // "hello POST created users"
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	log.Println("Server started at :2565")
	log.Fatal(http.ListenAndServe(":2565", nil))
	log.Println("bye")
}
