// https://leetcode.cn/problems/length-of-longest-subarray-with-at-most-k-frequency/description/
func maxSubarrayLength(nums []int, k int) int {
    ans := 0
    start := 0
    count := map[int]int{}
    for index, x := range nums {
        if _, ok := count[x]; ok {
            count[x] += 1
            for count[x] > k {
                count[nums[start]] -= 1
                start += 1
            }
        } else {
            count[x] = 1
        }
        // 到这里保证map中的值都符合好数组
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
