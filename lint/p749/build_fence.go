package main

import "fmt"

func main() {
	fmt.Println(isBuild(4))
	fmt.Println(isBuild(10))
	fmt.Println(isBuild(6))
}

func isBuild(x int) string {
	m := make([]bool, x+1)
	if x >= 3 {
		m[3] = true
	}
	if x >= 7 {
		m[7] = true
	}
	for i := 4; i <= x; i++ {
		if m[i-3] {
			m[i] = true
		}
		if i-7 > 0 && m[i-7] {
			m[i] = true
		}
	}
	if m[x] {
		return "YES"
	} else {
		return "NO"
	}
}
