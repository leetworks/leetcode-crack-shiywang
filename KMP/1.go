package main

import (
	"fmt"
)

func match(a,b string) bool {
	if len(a) > len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func search(pattern, str string) int {
	for i := 0; i < len(str); i++ {
		if match(pattern, str[i:]) {
			return i
		}
	}
	return -1
}

func main() {
	P := "DAB"
	S := "ABCDABCA"

	if v := search(P,S); v != -1 {
		fmt.Println("pattern matched at: a[", v, "]")
	} else {
		fmt.Println("not pattern matched")
	}

}
