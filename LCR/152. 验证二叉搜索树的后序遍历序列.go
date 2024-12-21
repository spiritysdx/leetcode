func verifyTreeOrder(postorder []int) bool {
    if len(postorder) == 0 {
        return true
    }
    // 辅助函数验证区间
    var verify func(start, end int) bool
    verify = func(start, end int) bool {
        if start >= end { 
            return true
        }
        // 根节点的值
        rootValue := postorder[end]
        // 找到分割点
        splitIndex := start
        for splitIndex < end && postorder[splitIndex] < rootValue {
            splitIndex++
        }
        // 验证右子树的所有值都大于根节点
        for i := splitIndex; i < end; i++ {
            if postorder[i] < rootValue {
                return false
            }
        }
        // 递归验证左右子树
        return verify(start, splitIndex-1) && verify(splitIndex, end-1)
    }
    return verify(0, len(postorder)-1)
}
