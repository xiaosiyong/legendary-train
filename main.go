package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {

		for {

			fmt.Println("1 seconds past")
			time.Sleep(20 * time.Millisecond)

		}

	}()
	t := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Println("2 seconds past")
		}
	}

}
