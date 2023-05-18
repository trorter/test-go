package main

import (
	"sort"
	"strconv"
)

func main() {

}

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(slice []int) string {
	var result string

	for _, v := range slice {
		result += strconv.Itoa(v)
	}

	return result
}

func MergeSlices(floatSlice []float32, intSlice []int32) []int {
	result := make([]int, 0, len(intSlice)+len(floatSlice))

	for _, v := range intSlice {
		result = append(result, int(v))
	}

	for _, v := range floatSlice {
		result = append(result, int(v))
	}

	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

	return result
}

func GetMapValuesSortedByKey(initialMap map[int]string) []string {
	var keys []int

	for key := range initialMap {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var result []string

	for _, value := range keys {
		result = append(result, initialMap[value])
	}

	return result
}
