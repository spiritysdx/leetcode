// // 哈希表法
// func inventoryManagement(stock []int) int {
//     size := len(stock)
//     if size == 1 {
//         return stock[0]
//     }
//     m := map[int]int{}
//     for _, key := range stock {
//         if v, exist := m[key]; exist {
//             v += 1
//             m[key] = v
//         } else {
//             m[key] = 1
//         }
//     }
//     for key, value := range m {
//         if value > size/2 {
//             return key
//         }
//     }
//     return -1
// }

// // 数组排序法
// func inventoryManagement(stock []int) int {
//     size := len(stock)
//     if size == 1 {
//         return stock[0]
//     }
//     sort.Slice(stock, func(i,j int) bool {
//         return stock[i] > stock[j]
//     })
//     return stock[int(size/2)]
// }


// 摩尔投票法的核心思想是通过抵消策略寻找可能的多数元素。
// 假设存在一个元素出现的次数超过数组长度的一半，则该元素一定可以在摩尔投票法中被筛选出来。
func inventoryManagement(stock []int) int {
    // 初始化候选元素和计数器
    candidate := 0
    count := 0
    // 第一阶段：寻找可能的候选多数元素
    for _, num := range stock {
        if count == 0 {
            // 当前计数器为0时，将当前元素设为候选元素
            candidate = num
            count = 1
        } else if num == candidate {
            // 如果当前元素等于候选元素，计数器加1
            count++
        } else {
            // 如果当前元素不等于候选元素，计数器减1
            count--
        }
    }
    // 第二阶段：验证候选元素是否为多数元素
    count = 0 // 重置计数器
    for _, num := range stock {
        if num == candidate {
            count++
        }
    }
    // 判断候选元素出现的次数是否超过数组长度的一半
    if count > len(stock)/2 {
        return candidate
    }
    // 如果不存在多数元素，返回-1表示未找到
    return -1
}
