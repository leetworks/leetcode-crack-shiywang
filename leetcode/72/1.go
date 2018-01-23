package main

import (
	"fmt"
)

func Min(a,b,c,d int) int {
	if c < a && c < b && c < d {
		return c
	}
	if a < b && a < c && a < d {
		return a
	}
	if b < a && b < c && b < d {
		return b
	}
	if d < a && d < b && d < c {
		return d
	}
	return a
}

var max = 9999999999999

func Dist(m, n int, A, B string) int {
	var a,b,c,d int

	if m == len(A) && n == len(B) {
		return 0
	}

	if m >= len(A) {
		return max
	}

	if n >= len(B) {
		return max
	}

	// A[m] == B[n]
	a = Dist(m+1, n+1, A, B)
	//insert at the front
	b = Dist(m, n+1, A, B) + 1
	//delete
	c = Dist(m+1, n, A, B) + 1
	//replace
	d = Dist(m+1, n+1, A, B) + 1

	return Min(a,b,c,d)
}

func main() {
	A := "ABCCA"
	B := "ABCDABCA"
	fmt.Println(Dist(0,0, A,B))

}
