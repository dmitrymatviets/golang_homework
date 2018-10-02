package main

import (
	"fmt"
)

func main() {
	s0, s1, s2, s3, s4, s5, s6 :=
		[]int{0, 1, 2, 3, 4, 5},
		[]int{0, 1, 2, 3, 4, 5},
		[]int{0, 1, 2, 3, 4, 5},
		[]int{0, 1, 2, 3, 4, 5},
		[]int{0, 1, 2, 3, 4, 5},
		[]int{0, 1, 2, 3, 4, 5},
		[]int{0, 1, 2, 3, 4, 5}

	rotate(s0, 0)

	rotate(s1, 1)

	rotate(s2, 2)

	rotate(s3, 3)

	rotate(s4, 4)

	rotate(s5, 5)

	rotate(s6, 6)

}

func rotate(list []int, shift int) {

	var from = 0
	var val = list[from]
	var nextGroup = 1
	var listLen = cap(list)

	var i int

	fmt.Println(list)
	fmt.Println("Length", listLen)
	fmt.Println("Сдвиг", shift)

	for i = 0; i < listLen; i++ {

		var to = (from - shift + listLen) % listLen
		if to == from {
			break
		}

		var temp = list[to]
		list[to] = val
		from = to
		val = temp

		if from < nextGroup {
			from = nextGroup
			nextGroup++
			val = list[from]
		}

	}

	fmt.Println(list)
	fmt.Println("****")

}
