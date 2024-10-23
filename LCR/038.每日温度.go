// 无中间值记录 双指针滑动窗口
// func dailyTemperatures(temperatures []int) []int {
//     res := []int{}
//     var status bool
//     for i1, v1 := range temperatures {
//         status = false
//         for i2, v2 := range temperatures[i1:] {
//             if v2 > v1 {
//                 res = append(res, i2)
//                 status = true
//                 break
//             }
//         }
//         if !status {
//             res = append(res, 0)
//         }
//     }
//     return res
// }

// 从后往前遍历，构造温度单调递减的栈，栈中存储索引方便计算
func dailyTemperatures(temperatures []int) []int {
    stack := make([]int, 0) // 用来存储索引的栈
    res := make([]int, len(temperatures)) // 结果数组，初始为0，长度与温度数组相同
    for i := len(temperatures) - 1; i >= 0; i-- { // 从右向左遍历温度数组
        // 当栈不为空且当前温度大于等于栈顶所指的温度时，弹出栈顶
        for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
            stack = stack[:len(stack)-1]
        }
        // 如果栈为空，说明后面没有比当前温度高的天
        if len(stack) == 0 {
            res[i] = 0
        } else {
            // 否则栈顶的索引对应的温度就是下一天比当前温度高的天
            res[i] = stack[len(stack)-1] - i
        }
        // 将当前索引压入栈中
        stack = append(stack, i)
    }
    return res
}
