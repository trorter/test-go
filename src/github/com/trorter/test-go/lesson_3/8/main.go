package main

import "sync"

func boring(wg *sync.WaitGroup, die chan bool) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case <-die:

				die <- true

				wg.Done()
				return
			case c <- "boring":
			}
		}
	}()

	return c
}

func main() {
	die := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)

	res1 := <-boring(&wg, die)
	res2 := <-boring(&wg, die)

	for i := 0; i < 10; i++ {
		println(res1, res2)
	}

	die <- true

	<-die

	wg.Wait()
}
