// https://leetcode.cn/problems/container-with-most-water/?envType=study-plan-v2&envId=top-100-liked
func maxArea(height []int) int {
    ans := 0
    left := 0
    right := len(height)-1
    for left < right {
        ans = max(ans, min(height[right],height[left])*(right-left))
        if height[left] <= height[right] {
            left++
        } else {
            right--
        }
    }
    return ans
}

func min(a,b int) int {
    if a > b {
        return b
    }
    return a
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
