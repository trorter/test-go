package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"test-go/src/github/com/trorter/test-go/test"
)

const constInt int32 = 1

const (
	zero = iota

	one
	_
	three = 3
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func listAsString(list []int) (result string, er error) {
	if list == nil || len(list) == 0 {
		return "", fmt.Errorf("list is empty")
	}

	for _, v := range list {
		result += strconv.Itoa(v)
	}
	return result, nil
}

func (p Person) Print() {
	fmt.Println("name=", p.Name, ", age=", p.Age, ", sex=", p.Sex)
}

func (p *Person) grow(years int) {
	p.Age += years
}

func main() {

	v1 := 1
	var v2 = 2.1
	var v3 string = "Hello"
	var v4 bool
	println(v1, v2, v3, v4)

	println(v2 + float64(constInt))

	arr1 := [3]int{1, 2, 3}

	var arr2 [10]float32

	arr3 := [...]int{1, 2, 3, 4, 5}

	var matrix [3][3]int

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(matrix)

	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)
	fmt.Println(slice)
	fmt.Println(slice[1:3])
	fmt.Println(slice[5])

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println("====================")

	slice = append(slice, arr3[:]...)
	fmt.Println(slice)

	//newSlice := make([]string, 0, 10)
	//fmt.Println(newSlice, len(newSlice), cap(newSlice))

	sliceToCopy := make([]int, len(slice))
	copy(sliceToCopy, slice)

	fmt.Println(sliceToCopy[2:len(sliceToCopy)])

	mm1 := map[string]string{}
	mm2 := map[string]map[string]string{}

	mm1["mm1Key"] = "mm1Value"
	mm2["mm2Key"] = mm1

	fmt.Println(mm2)

	value, flag := mm1["test"]
	fmt.Println(value, flag)

	b := rand.Intn(11)
	println("b=", b)

	if b == 1 {
		fmt.Println("b is 1")
	} else {
		fmt.Println("b is not 1")
	}

	if name, exist := mm1["mm1Key1"]; exist {

		fmt.Println(
			"name is",
			name)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Println(test.A)

	//test panic and defer and recover
	panicResult := panicAndDefer()
	fmt.Println("panicResult=", panicResult)

	//for i := 0; i < 10; i++ {
	//	defer println(i)
	//}

	println("=====")

	person1 := Person{}
	person1.Print()
	person1.grow(11)
	person1.Print()
}

func panicAndDefer() (result string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	result = "successful result"

	panic("test panic")

	result = "unsuccessful result"

	return result
}
