package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

type Logger struct {
	Handler http.Handler
}

// method of Logger
func (l Logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, req) // in this case Handler is mux
	log.Printf("Server http middleware: %s %s %s %s", req.RemoteAddr, req.Method, req.URL, time.Since(start))
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		u, p, ok := req.BasicAuth()
		log.Println("auth: ", u, p, ok)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`can't parse the basic auth.`))
			return
		}

		if u != "apidemo" || p != "34567" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`Username/Password is Incorrect.`))
			return
		}

		fmt.Println("Auth passed.")
		next(w, req)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", AuthMiddleware(usersHandler))
	mux.HandleFunc("/health", healthHandler)

	logMux := Logger{Handler: mux}
	srv := http.Server{
		Addr:    ":2565",
		Handler: logMux,
	}

	log.Println("Server started at :2565")
	// log.Fatal(http.ListenAndServe(":2565", mux))
	log.Fatal(srv.ListenAndServe())
	log.Println("bye")
}
