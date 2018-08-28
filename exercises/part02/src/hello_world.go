package main

import (
	"fmt"
	"time"
)

func main () {
	for {
		fmt.Println("Hello World!")
		time.Sleep(1000*time.Millisecond)
	}
}
