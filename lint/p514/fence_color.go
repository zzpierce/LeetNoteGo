package main

import "fmt"

func main() {
	fmt.Println(numWays(3, 2))
}

func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}
	m := make([]*P, n)
	m[0] = &P{N: k}
	for i := 1; i < n; i++ {
		m[i] = &P{
			E: m[i-1].N,
			N: (m[i-1].E + m[i-1].N) * (k - 1),
		}
	}
	return m[n-1].E + m[n-1].N
}

type P struct {
	E int
	N int
}
