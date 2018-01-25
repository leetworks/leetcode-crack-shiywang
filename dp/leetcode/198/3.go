func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
// func solve(idx int, nums, memory []int) int {
// 	if idx < 0 {
// 		return 0
// 	}
// 	if memory[idx] != -1 {
// 		return memory[idx]
// 	}
    //ONLY KEEP THIS ONE !!
	// memory[idx] = Max(nums[idx]+solve(idx-2, nums, memory), solve(idx-1, nums, memory))
//     return memory[idx]
// }
func rob(nums []int) int {
    if len(nums) == 0 {
        return 0;
    }
    if len(nums) == 1 {
        return nums[0];
    }
    if len(nums) == 2 {
        return Max(nums[0], nums[1])
    }
    memory := make([]int, len(nums))
    memory[0] = nums[0]
    memory[1] = Max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        memory[i] = Max(nums[i]+memory[i-2], memory[i-1])
    }
    return memory[len(nums)-1]
}
