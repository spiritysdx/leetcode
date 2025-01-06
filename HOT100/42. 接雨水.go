// https://leetcode.cn/problems/trapping-rain-water/?envType=study-plan-v2&envId=top-100-liked
func trap(height []int) int {
    firsts := []int{height[0]}
    for i, v := range height {
        if i == 0 {
            continue
        }
        if v > firsts[i-1] {
            firsts = append(firsts, v)
        } else {
            firsts = append(firsts, firsts[i-1])
        }
    }
    size := len(height)
    ends := []int{height[size-1]}
    for i := size-2; i>0;i--{
        if height[i] > ends[0] {
            ends = append([]int{height[i]}, ends...)
        } else {
            ends = append([]int{ends[0]}, ends...)
        }
    }
    ans := 0
    for i := 0; i < size-1; i++ {
        minHeight := min(firsts[i], ends[i])
        if minHeight > height[i] {
            ans += minHeight - height[i]
        }
    }
    return ans
}
