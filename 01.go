package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main () {
	fmt.Println (time.Now ().UnixNano ())
	/// ticker := time.NewTicker (time.Second)
	retChan := make (chan int)
	timeChan := make (chan int)

	go func (retChan *chan int) {
		for {
			time.Sleep (time.Millisecond * 100)
			*retChan<- rand.Int ()
		}
	}(&retChan)

	go func (timeChan *chan int) {
		for {
			time.Sleep (time.Millisecond * 800)
			*timeChan<- 0
		}
	}(&timeChan)

	var ret int
	for {
		select {
			case <-timeChan:
				fmt.Println ("---------------------------------")
			case ret = <-retChan:
				fmt.Println ("random: ", ret)
		}
	}
}
