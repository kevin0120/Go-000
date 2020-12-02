package main

import (
	"fmt"
	"github.com/kevin0120/Go-000/Week02/sync"
	"github.com/pkg/errors"
	"time"
)

type errorString struct {
	s string
}

var slice = []int{1, 2, 3}

//go 中结构体的==判断的是结构提中的每个值
func main() {
	a := &errorString{"hello"}
	b := &errorString{"hello"}
	c := errorString{"hello"}
	d := errorString{"hello"}

	if a != b {
		fmt.Println("a!=b")
	}

	if *a == *b {
		fmt.Println("*a==*b")
	}

	if c == d {
		fmt.Println("c==d")
	}

	if &c != &d {
		fmt.Println("&c!=&d")
	}

	//fmt.Println("hello world !")
	//
	var e error
	e = errors.New("hello world, i am sorry")
	fmt.Println(e.Error())
	//
	//
	//fmt.Println(slice[5])
	//defer func() {
	//	if err:=recover(); err !=nil{
	//		//
	//		fmt.Println("hello panic")
	//	}
	//}()
	//fmt.Println(slice[5])

	//fmt.Println(len(slice),cap(slice))
	go sync.Go(func() {
		fmt.Println(slice[5])
	})

	for {
		time.Sleep(1 * time.Second)
	}

}
