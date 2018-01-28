
//这题的brute force本来是可以用一个三层循环写完的大概最多时间复杂度也就是O(n3)
//但是非要递归写的话，有很多分支条件是没有意义的，然后时间复杂度大概是O(3^n * n)
//但是记忆化搜索就很明确，只要dp[b][e]算过，就不算了
var maxVal int
func Max(a, b int) int {
    if a > b {
        return a
    }
    
    return b
}

func Sum(i, j int, num []int) int {
    sum := 0
    for a := i; a <= j; a++ {
        sum = sum + num[a]
    }
    return sum
}

func slove(b, e int, nums []int) {
    if e >= len(nums) {
        return 
    }
    if b > e {
        return
    }
    if b <= e {
        if t := Sum(b,e,nums); t > maxVal {
            maxVal = t
        }
    }
    slove(b+1, e+1, nums)
    slove(b, e+1, nums)
    slove(b+1, e, nums)
}

func maxSubArray(nums []int) int {
    maxVal = -99999999999
    slove(0,0,nums);
    return maxVal
}
