var visit []bool
var g [][]int
var gIndex int

func pow(x, n int) int {
    ret := 1 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
    for n != 0 {
        if n%2 != 0 {
            ret = ret * x
        }
        n /= 2
        x = x * x
    }
    return ret
}

func slove(idx int, nums []int) {
    if idx == len(nums) {
        for i := 0; i < len(visit); i++ {
            if visit[i] {
                g[gIndex] = append(g[gIndex], nums[i])
            }
        }
        gIndex = gIndex + 1
        return
    }
    visit[idx] = true
    slove(idx+1, nums)
    visit[idx] = false
    slove(idx+1, nums)
}

func subsets(nums []int) [][]int {
    visit = make([]bool, len(nums))
    g = make([][]int, pow(2,len(nums)))
    gIndex=0
    slove(0, nums)
    return g
}
