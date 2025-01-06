// https://leetcode.cn/problems/subarray-sum-equals-k/?envType=study-plan-v2&envId=top-100-liked
func subarraySum(nums []int, k int) int {
    // 生成前缀和，首元素必须为0，表示存在前缀和为0的数组
    sums := make([]int, len(nums)+1)
    for i, x := range nums {
        sums[i+1] = sums[i]+x
    }
    // 根据前缀和查找符合条件的和
    ans := 0
    cnt := map[int]int{}
    for _, sj := range sums {
        ans += cnt[sj-k]
        cnt[sj]++
    }
    return ans
}
