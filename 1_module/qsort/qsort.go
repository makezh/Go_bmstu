package main

import (
	"fmt"
)

var to_sort = []int{5, 4, 3, 5, 2, 3, 2, 0, 2, -1, -342, 213, 341, 3, 5}

func less(i, j int) (answer bool) {
	answer = to_sort[i] < to_sort[j]
	return
}

func swap(i, j int) {
	to_sort[i], to_sort[j] = to_sort[j], to_sort[i]
}

func partition(left int, right int, less func(i, j int) bool, swap func(i, j int)) (i int) {
	i = right
	j := right

	for j > left {
		if less(j, left) {
			j--
		} else {
			swap(i, j)
			i--
			j--
		}
	}
	if i != left {
		swap(i, left)
	}
	return
}

func QuickSort(left int, right int, less func(i, j int) bool, swap func(i, j int)) {
	if left < right {
		part := partition(left, right, less, swap)
		QuickSort(left, part-1, less, swap)
		QuickSort(part+1, right, less, swap)
	}
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	QuickSort(0, n-1, less, swap)
}

func main() {
	qsort(len(to_sort), less, swap)
	fmt.Println(to_sort)
}
