package main

import "fmt"

func main() {
	fmt.Println("start")

	go process(0)

	go func() {
		fmt.Println("Anonymous function")
	}()

	for i := 0; i < 1_000; i++ {
		go process(i)
	}

	fmt.Scanln()
}

func process(i int) {
	fmt.Println("Process", i)
}
