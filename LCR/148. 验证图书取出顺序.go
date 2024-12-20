func validateBookSequences(putIn []int, takeOut []int) bool {
    stack := []int{}
    j := 0 // 指向 takeOut 的索引
    for _, book := range putIn {
        stack = append(stack, book) // 模拟入栈
        // 当栈顶元素和 takeOut 的当前元素相等时，进行出栈
        for len(stack) > 0 && j < len(takeOut) && stack[len(stack)-1] == takeOut[j] {
            stack = stack[:len(stack)-1] // 弹出栈顶元素
            j++                          // 移动 takeOut 的指针
        }
    }

    // 如果最终栈为空，说明是有效的入栈和出栈序列
    return len(stack) == 0
}
