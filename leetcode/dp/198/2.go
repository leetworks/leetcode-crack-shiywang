func solve(idx int, nums, memory []int) int {
	if idx < 0 {
		return 0
	}
	if memory[idx] != -1 {
		return memory[idx]
	}
	memory[idx] = Max(nums[idx]+solve(idx-2, nums, memory), solve(idx-1, nums, memory))
    return memory[idx]
}
//pick: idx --> (idx-2, idx-3, idx-4)
//not pick: idx --> (idx-1, idx-2, idx-3, idx-4)
func rob(nums []int) int {
    memory := make([]int, len(nums))
    for i := range memory {
        memory[i] = -1
    }
	return solve(len(nums)-1, nums, memory)
}
