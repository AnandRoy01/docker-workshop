package main

import (
	"fmt"
	"time"
)

func main () {
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Hello World!")
	}
}
