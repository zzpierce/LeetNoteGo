package main

import "fmt"

func main() {
	fmt.Println(MedianSort([]int{4, 4, 1, 11, 1, 1, 1, 1, 1, 1}))
}

// 中位数O(n)
func MedianSort(n []int) ([]int, int) {
	t := find(n, 0, len(n)-1)
	return n, t
}

func find(n []int, x, y int) int {
	if x == y {
		return x
	}
	swap(n, x, (x+y)/2)
	t := x
	s := x
	e := y
	for s < e {
		for n[e] >= n[t] && s < e {
			e--
		}
		if n[e] < n[t] && s < e {
			c := n[t]
			n[t] = n[e]
			n[e] = c
			t = e
			e--
		}
		for n[s] <= n[t] && s < e {
			s++
		}
		if n[s] > n[t] && s < e {
			c := n[t]
			n[t] = n[s]
			n[s] = c
			t = s
			s++
		}
	}
	if t > (len(n)-1)/2 {
		t = find(n, x, t-1)
	} else if t < (len(n)-1)/2 {
		t = find(n, t+1, y)
	}
	return t
}

func swap(n []int, x, y int) {
	c := n[x]
	n[x] = n[y]
	n[y] = c
}
