package main

import (
	"log"
	"time"
)

func main() {
	c := 0
	for {
		c++
		log.Printf("log-test:%d", c)
		time.Sleep(1 * time.Second)
	}
}
