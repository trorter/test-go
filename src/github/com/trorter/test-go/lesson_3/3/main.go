package main

import (
	"fmt"
	"strconv"
)

var c chan int

func main() {

	c = make(chan int, 100)

	go greet(c)

	for i := 0; i < 5; i++ {
		first := <-c
		fmt.Println("red first")

		second := <-c
		fmt.Println("red second")

		fmt.Println(first, ",", second)
	}

	stuff := make(chan int, 7)
	for i := 0; i < 19; i += 3 {
		stuff <- i
	}
	close(stuff)

	fmt.Println("res", processStuff(stuff))

}

func processStuff(c <-chan int) (res string) {
	for i := range c {
		res += strconv.Itoa(i)
	}
	return res
}

func greet(c chan<- int) {
	i := 0
	for {
		i++
		c <- i
		fmt.Println("wrote first")
		i++
		c <- i
		fmt.Println("wrote second")
	}
}
