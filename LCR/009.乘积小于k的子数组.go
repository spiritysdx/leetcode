func numSubarrayProductLessThanK(nums []int, k int) int {
    ans := 0 // 初始化答案为0
	prod, i := 1, 0 // 初始化乘积为1，起始指针i为0
    for j, num := range nums { // 遍历数组，j为当前指针，num为当前值
        prod *= num // 更新当前乘积
        for ; i <= j && prod >= k; i++ { // 当乘积大于等于k时，移动左指针i
            prod /= nums[i] // 将乘积除以左指针i指向的值
        }
        ans += j - i + 1 // 计算以当前j为结束的子数组数量并累加到答案中
    }
	return ans // 返回最终答案
}
