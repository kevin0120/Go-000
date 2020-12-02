package sync

import "fmt"

func Go(x func()) {

	defer func() {
		if err := recover(); err != nil {
			//
			fmt.Println("hello panic")
		}
	}()

	x()
}
