package main

import (
	"fmt"
	"math"
)

func main() {
	var values []int = []int{17, 28, 43, 67, 88, 92, 100}
	target := 42
	min := 0
	max := len(values)

	loc := binSearch(target, values...)
	fmt.Println(loc)

	locR := binSearchRec(min, max, target, values...)
	fmt.Println(locR)
}

func binSearch(target int, i ...int) int {
	min := 0
	max := len(i)
	mid := 0

	for min <= max {
		mid = int(math.Round(float64((min + max) / 2)))
		if i[mid] == target {
			return mid
		} else if i[mid] < target {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}

func binSearchRec(min, max, target int, i ...int) int {
	mid := int(math.Round(float64((min + max) / 2)))
	if min > max {
		return -1
	}
	if i[mid] == target {
		return mid
	} else if i[mid] < target {
		return binSearchRec(mid+1, max, target, i...)
	} else {
		return binSearchRec(min, mid-1, target, i...)
	}

}
