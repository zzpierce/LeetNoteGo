package main

import "fmt"

func main() {
	a := []int{3, 1}
	sortIntegers(&a)
	fmt.Println(a)
}

func sortIntegers(A *[]int) {
	a := *A
	l := len(a)
	// select
	x, y := 0, 0
	for x < l {
		s := a[x]
		y = x
		for i := x + 1; i < l; i++ {
			if a[i] < s {
				s = a[i]
				y = i
			}
		}
		t := a[x]
		a[x] = a[y]
		a[y] = t
		x++
	}
}
