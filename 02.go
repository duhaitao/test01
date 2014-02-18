package main

import (
		"fmt"
)

func main () {
	fmt.Println ("hello")
	//var sum int

	aaa := make ([]int, 10)
	for i := 1; i <= 10; i++ {
		aaa[i - 1] = i
	}

	ch := make (chan int)
	go func (aaa []int, ch chan int) {
		var sum int
		for _, val := range aaa {
			sum += val
		}
		ch<- sum
	} (aaa, ch)
	sum := <-ch
	fmt.Println (sum)
}

// test ssh-key 1
// test ssh-key 2
