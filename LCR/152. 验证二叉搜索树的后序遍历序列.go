// func verifyTreeOrder(postorder []int) bool {
//     if len(postorder) == 0 {
//         return true
//     }
//     // 辅助函数验证区间
//     var verify func(start, end int) bool
//     verify = func(start, end int) bool {
//         if start >= end { 
//             return true
//         }
//         // 根节点的值
//         rootValue := postorder[end]
//         // 找到分割点
//         splitIndex := start
//         for splitIndex < end && postorder[splitIndex] < rootValue {
//             splitIndex++
//         }
//         // 验证右子树的所有值都大于根节点
//         for i := splitIndex; i < end; i++ {
//             if postorder[i] < rootValue {
//                 return false
//             }
//         }
//         // 递归验证左右子树
//         return verify(start, splitIndex-1) && verify(splitIndex, end-1)
//     }
//     return verify(0, len(postorder)-1)
// }

func verifyTreeOrder(postorder []int) bool {
    if len(postorder) == 0 {
        return true
    }
    stack := []int{}
    root := int(^uint(0) >> 1) // 初始化为正无穷
    // 从后向前遍历数组
    for i := len(postorder) - 1; i >= 0; i-- {
        // 当前值不能大于根节点值，否则不合法
        if postorder[i] > root {
            return false
        }
        // 模拟构建BST，找到当前节点的父节点
        for len(stack) > 0 && stack[len(stack)-1] > postorder[i] {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        // 将当前值入栈
        stack = append(stack, postorder[i])
    }
    return true
}
