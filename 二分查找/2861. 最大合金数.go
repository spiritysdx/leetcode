// https://leetcode.cn/problems/maximum-number-of-alloys/
func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
    // 定义二分查找的左边界为0
    left := 0
    // 定义二分查找的右边界为预算允许的最大可能值
    right := budget + stock[0] // +stock[0]是为了这个上界比实际最大值大，需要检索的范围包含了实际答案
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

/**
这个解法的主要思路是：

使用二分查找来寻找最大可能的合金产量。
对于每个猜测的产量(mid)：

遍历每台机器，计算使用该机器生产mid个合金所需的成本
考虑已有的库存和需要额外购买的原材料
如果有任何一台机器可以在预算内完成生产，则这个产量是可行的


注意事项：

使用int64进行计算防止溢出
提前结束计算如果成本已经超过预算
正确处理二分查找的边界条件



时间复杂度：O(k * n * log(budget))，其中：

k 是机器数量
n 是金属类型数量
log(budget) 是二分查找的次数

空间复杂度：O(1)，只使用了常量额外空间。
**/
