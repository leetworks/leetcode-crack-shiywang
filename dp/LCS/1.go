//最长公共子序列。
//题目：如果字符串一的所有字符按其在字符串中的顺序出现在另外一个字符串二中，
//则字符串一称之为字符串二的子串。
//注意，并不要求子串（字符串一）的字符必须连续出现在字符串二中。
//请编写一个函数，输入两个字符串，求它们的最长公共子串，并打印出最长公共子串。
//例如：输入两个字符串BDCABA和ABCBDAB，字符串BCBA和BDAB都是是它们的最长公共子序列，则输出它们的长度4，并打印任意一个子序列。


package main
import (
	"fmt"
)

func Max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Max3(a, b, c int) int {
	if a > b && a > c {
		return a
	}
	if b > a && b > c {
		return b
	}
	return c
}


func solve(a, b int, A, B string) int {
	if a >= len(A) || b >= len(B) {
		return 0
	}

	if A[a] == B[b] {
		return solve(a+1, b+1, A, B) + 1
	}

	return Max2(solve(a, b+1, A, B),solve(a+1, b, A, B))
}

func main() {
	//str1 := "BDCABA"
	//str2 := "ABCBDAB"

	str3 := "aaaaacbadfsd"
	str4 := "aaaaacafd"
	fmt.Println(solve(0,0,str3,str4))
}

