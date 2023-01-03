package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Start Server")
	mux := http.NewServeMux()

	mux.HandleFunc("/muxlai", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`mux hang lai`))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`HELLO`))
	})

	srv := http.Server{
		Addr:    ":2565",
		Handler: mux,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	fmt.Println("server starting at :2565")

	// channel with 1 buffer
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	fmt.Println("shutting down...")
	//The Shutdown function is called with a Context argument, which can be used to specify a timeout for the shutdown process. If the timeout expires before the shutdown is complete, the function returns an error.
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Println("shutdown err:", err)
	}
	fmt.Println("bye bye")
}
