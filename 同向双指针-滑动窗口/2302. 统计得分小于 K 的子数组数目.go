// https://leetcode.cn/problems/count-subarrays-with-score-less-than-k/description/
func countSubarrays(nums []int, k int64) int64 {
    ans := 0
    left := 0
    s1 := 0 // 数组之和
    s2 := 0 // 数组长度
    for right, x := range nums {
        s1 += x
        s2 += 1
        for int64(s1 * s2) >= k {
            s1 -= nums[left]
            s2 -= 1
            left+=1
        }
        ans += right-left+1
    }
    return int64(ans)
}
