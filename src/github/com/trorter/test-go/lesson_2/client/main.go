package main

import "fmt"

type Client interface {
	Do(r string) (string, error)
}

type ClientImpl struct {
	result string
}

func (c ClientImpl) Do(r string) (string, error) {
	return c.result + "\n" + r, nil
}

type ClientFunc func(r string) (string, error)

func (t ClientFunc) Do(r string) (string, error) {
	return t(r)
}

type Decorator func(Client) Client

func main() {

	client := Decorate(
		ClientImpl{"init"},
		Authorization("token1"),
		Authorization("token2"))

	result, _ := client.Do("request")
	fmt.Println(result)
}
