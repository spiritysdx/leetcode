func maxSales(sales []int) int {
    // 初始化最大值为销售数据的第一个值
    max := sales[0]
    // 从第二个元素开始遍历销售数据
    for i := 1; i < len(sales); i++ {
        // 如果当前销售额与前一个销售额之和大于当前销售额
        // 则将当前销售额更新为两者之和，累积连续的最大销售额
        if sales[i] + sales[i-1] > sales[i] {
            sales[i] += sales[i-1]
        }
        // 如果当前销售额大于当前记录的最大值，则更新最大值
        if sales[i] > max {
            max = sales[i]
        }
    }
    // 返回最大的销售额
    return max
}
