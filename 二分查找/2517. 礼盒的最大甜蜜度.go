// https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/description/
func maximumTastiness(price []int, k int) int {
    // 对价格进行排序
    slices.Sort(price)
    size := len(price)
    // 二分查找范围：最小差值为 0，最大差值为数组最大值-最小值
    left, right := 0, price[size-1]-price[0]
    for left < right {
        mid := (left + right + 1) / 2 // 注意这里是向上取中值，避免死循环
        if check(price, mid, k) {
            left = mid
        } else {
            right = mid - 1
        }
    }
    return left
}

func check(price []int, mid, k int) bool {
    // 贪心选择：从前往后找相隔至少为 mid 的数
    count, prev := 1, price[0] // 第一个数总是选中
    for _, p := range price[1:] {
        if p-prev >= mid {
            count++
            prev = p
        }
        if count >= k {
            return true
        }
    }
    return false
}
