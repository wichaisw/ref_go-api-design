package main

import (
	"fmt"
	"time"
)

func slow(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s, ":", i)
	}
}

func main() {
	//channel
	done := make(chan bool)
	go func() {
		slow("gopher")
		done <- true
	}()
	slow("outside")
	<-done

	fmt.Println("all task done.")
}
