func solve(idx int, nums []int) int {
	if idx < 0 {
		return 0
	}
	return Max(nums[idx]+solve(idx-2, nums), solve(idx-1, nums))
}
func rob(nums []int) int {
	return solve(len(nums)-1, nums)
}
