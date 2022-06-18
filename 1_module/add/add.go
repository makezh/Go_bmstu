package main

import "fmt"

func addition_length(a []int32, length int) (answer []int32) { // дополним нулями
	if len(a) >= length {
		answer = a
		return
	}

	addition := make([]int32, length-len(a))
	answer = append(a, addition...)
	return
}

func add(a, b []int32, p int) (answer []int32) {
	if len(a) > len(b) {
		b = addition_length(b, len(a))
	} else {
		if len(b) > len(a) {
			a = addition_length(a, len(b))
		}
	}

	plusOne := int32(0)
	for i := 0; i < len(a); i++ {
		sum := (a[i] + b[i] + plusOne) % int32(p)
		if (a[i] + b[i] + plusOne) >= int32(p) {
			plusOne = 1
		} else {
			plusOne = 0
		}
		answer = append(answer, sum)
	}
	if plusOne == 1 {
		answer = append(answer, plusOne)
	}
	return
}

func main() {
	a := []int32{5, 2}
	b := []int32{0, 0, 5}

	fmt.Println(add(a, b, 10))
}
