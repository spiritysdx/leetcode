// https://leetcode.cn/problems/koko-eating-bananas/description/
func minEatingSpeed(piles []int, h int) int {
    // 检查珂珂是否能以速度 k 在 h 小时内吃完所有香蕉
    canEatAll := func(k int) bool {
        hours := 0
        for _, pile := range piles {
            // 计算吃完当前堆香蕉所需的小时数
            hours += (pile + k - 1) / k // 等价于向上取整 pile / k
            if hours > h {
                return false
            }
        }
        return true
    }
    // 设置二分查找的左右边界，找piles中的最大值赋值到right
    left, right := 1, 0
    for _, pile := range piles {
        if pile > right {
            right = pile
        }
    }
    // 二分查找最小速度 k
    for left < right {
        mid := left + (right-left)/2
        if canEatAll(mid) {
            right = mid // 如果能在 h 小时内吃完，尝试更小的速度
        } else {
            left = mid + 1 // 如果不能吃完，增加速度
        }
    }
    return left // 此时 left 等于最小速度
}
