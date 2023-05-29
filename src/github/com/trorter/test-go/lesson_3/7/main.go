package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	table := make(chan *Ball)

	go player("ping", table)
	go player("pong", table)

	newBall := new(Ball)
	table <- newBall
	go func() {
		for {
			fmt.Println("bad")
			ball := <-table
			ball.hits++
			time.Sleep(2 * time.Millisecond)
			table <- ball
		}
	}()
	time.Sleep(1 * time.Second)
	<-table
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
