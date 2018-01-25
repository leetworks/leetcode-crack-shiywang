package main

import (
	"fmt"
)

var max = 10000000000000

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var dp[1000][100000]int

func slove(idx, remain int, coins []int) int {
	if idx >= len(coins) {
		return max
	}

	if remain == 0 {
		return 0
	}

	if remain < 0 {
		return max
	}

	if dp[idx][remain] != 0 {
		return dp[idx][remain]
	}

	dp[idx][remain] = slove(idx, remain-coins[idx], coins) + 1
	dp[idx+1][remain] = slove(idx+1, remain, coins)

	return Min(dp[idx][remain],dp[idx+1][remain])
}

func coinChange(coins []int, amount int) int {
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100000; j++ {
			dp[i][j] = -1
		}
	}
	if v := slove(0, amount, coins); v == max {
		return -1
	} else {
		return v
	}
}

func main() {
	//[270,373,487,5,62]
	//8121
	fmt.Println(coinChange([]int{2}, 1))
}
