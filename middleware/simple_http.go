package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func usersHandler(w http.ResponseWriter, req *http.Request) {
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
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("Server http middleware: %s %s %s %s", req.RemoteAddr, req.Method, req.URL, time.Since(start))
	}
}

func main() {
	http.HandleFunc("/users", logMiddleware(usersHandler))
	http.HandleFunc("/health", logMiddleware(healthHandler))

	log.Println("Server started at :2565")
	log.Fatal(http.ListenAndServe(":2565", nil))
	log.Println("bye")
}
