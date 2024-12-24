// https://leetcode.cn/problems/maximum-number-of-alloys/
func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
    // 定义二分查找的左边界为0
    left := 0
    // 定义二分查找的右边界为预算允许的最大可能值
    right := budget + stock[0]
    // 进行二分查找，直到左右边界相遇
    for left < right {
        // 计算中间值
        mid := (left + right + 1) / 2
        // found 标记是否找到可行解
        found := false
        // 遍历每台机器
        for i := 0; i < k; i++ {
            // 当前机器所需的总成本
            totalCost := 0
            // 计算使用当前机器生产 mid 个合金所需的成本
            for j := 0; j < n; j++ {
                // 计算需要的第j种金属总量
                need := int64(composition[i][j]) * int64(mid)
                // 如果库存不足，需要购买的量
                if need > int64(stock[j]) {
                    buy := need - int64(stock[j])
                    // 累加购买成本，注意类型转换防止溢出
                    totalCost += int(buy) * cost[j]
                    // 如果成本已经超过预算，提前结束
                    if totalCost > budget {
                        break
                    }
                }
            }
            // 如果当前机器的成本在预算内，标记找到可行解
            if totalCost <= budget {
                found = true
                break
            }
        }
        // 根据是否找到可行解调整二分查找的边界
        if found {
            // 如果找到可行解，尝试更大的产量
            left = mid
        } else {
            // 如果没找到可行解，需要减小产量
            right = mid - 1
        }
    }
    // 返回最大可能的合金产量
    return left
}
