func maxSubArray(nums []int) int {
    max := nums[0]
    sum := 0
    for i := 0; i < len(nums); i++ {
        if sum < 0 {
            sum = 0 + nums[i]
        } else {
            sum = sum + nums[i]
        }
        if sum >= max {
            max = sum
        }
    }
    return max
}
