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

/**
既然是数组中，下一个更大的元素，即用单调栈实现。

既然是 栈，那么一定是 待入栈的元素 和 栈顶 比较。因为需找更大的元素，所以选择单调递减栈。因为只有此结构才可保证 栈顶 > 待入栈的元素。

因为需找数组右侧下一个更大的元素，所以逆序遍历数组。因为逆序遍历才可保证 数组右侧的数都在栈里。而栈又是递减栈，才可通过栈顶 > 待入栈的元素，来保证 待入栈的元素 的 右侧第一个最大的元素 是 栈顶。

注意 corner case：

    求数组下一个更大的元素，则新push的元素 <= 栈顶时出栈，这里的条件是包含等号的。
    求数组下一个更大于或等于的元素，则新push的元素 < 栈顶时出栈，这里的条件是不包含等号的。
**/

func dailyTemperatures(temperatures []int) []int {
    stack := []int{} // 用于存储索引的栈
    ans := make([]int, len(temperatures)) // 用于存储结果，长度与温度数组相同
    for i := len(temperatures) - 1; i >= 0; i-- { // 从右向左遍历温度数组
        // 当栈不为空且当前温度大于或等于栈顶所指的温度时，弹出栈顶
        for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
            stack = stack[:len(stack)-1] // pop栈顶
        }
        // 如果栈不为空，栈顶索引对应的温度就是下一个比当前温度高的天数
        if len(stack) > 0 {
            ans[i] = stack[len(stack)-1] - i // 计算天数差
        }
        // 将当前索引压入栈
        stack = append(stack, i) // push当前索引
    }
    return ans
}
