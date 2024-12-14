func twoSum(price []int, target int) []int {
    // 创建一个哈希表，用于存储数组中的值及其索引
    hashMap := make(map[int]int)
    // 遍历数组中的每个元素，查找与目标的差值是否已检索过，若已检索过就可以组成组合返回
    for i, num := range price {
        // 计算目标值与当前数的差值
        complement := target - num
        // 检查差值是否存在于哈希表中
        if idx, found := hashMap[complement]; found {
            // 如果找到，返回对应的两个值
            return []int{price[idx], price[i]}
        }
        // 将当前值及其索引存入哈希表
        hashMap[num] = i
    }
    // 如果没有找到符合条件的两数之和，返回空数组
    return []int{}
}
