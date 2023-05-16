package main

import "fmt"

const constInt int32 = 1

const (
	zero = iota

	one
	_
	three = 3
)

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
}
