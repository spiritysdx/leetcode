// https://leetcode.cn/problems/container-with-most-water/description/
func maxArea(height []int) int {
    // 哪条边短移动哪一边，都一样就随便移动
    res := 0
    left := 0
    right := len(height)-1
    for left < right {
        areas := (right-left) * min(height[left], height[right])
        res = max(res, areas)
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }   
    return res
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
