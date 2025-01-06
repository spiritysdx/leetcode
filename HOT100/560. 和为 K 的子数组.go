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

func subarraySum(nums []int, k int) int {
    // 前缀和生成
    temp := []int{0}
    s := 0
    for _, v := range nums {
        s += v
        temp = append(temp, s)
    }
    // 遍历前缀和，如果存在 当前前缀和-目标和值 的记录存在，证明去掉这个 差值的和 的时候剩下的连续数组就是目标和值的数组
    ans := 0
    tempMap := map[int]int{}
    for _, v := range temp {
        ans += tempMap[v-k]
        tempMap[v]++
    }
    return ans
}
