package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("hello world, test2-main")
	log.Info("new info")
	log.Warn("new warn")
	log.Fatal("new fatal")
	// the next line won't be executed
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	log.Info("+new info")
}
