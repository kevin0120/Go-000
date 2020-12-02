package main

import (
	"fmt"
	"time"
)

var slice1 = []int{1, 2, 3}

func Go(x func()) {

	defer func() {
		if err := recover(); err != nil {
			//
			fmt.Println("hello panic")
		}
	}()

	x()
}

func main() {
	go Go(func() {
		fmt.Println(slice1[500])
	})
	for {
		time.Sleep(1 * time.Second)
	}
}
