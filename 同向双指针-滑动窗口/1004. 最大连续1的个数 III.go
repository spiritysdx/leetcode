// https://leetcode.cn/problems/max-consecutive-ones-iii/
func longestOnes(nums []int, k int) int {
    // 找0的数量
    ans := 0
    start := 0
    temp := 0
    for index, x := range nums {
        if x == 0 {
            temp += 1
        }
        for temp > k {
            if nums[start] == 0 {
                temp -= 1
            }
            start += 1
        }
        // 在这之前需要保证规则是满足的
        ans = max(ans, index-start+1)
    }
    return ans
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
