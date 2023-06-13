package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(2 * time.Second)

	<-timer.C

	fmt.Println("Timer fired!")

	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(2000 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
