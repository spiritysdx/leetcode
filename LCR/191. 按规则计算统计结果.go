func statisticalResult(arrayA []int) []int {
    // 不能使用除法，因为大数
    n := len(arrayA)
    // 创建结果数组
    arrayB := make([]int, n)
    // 第一步：计算左边的积
    temp := 1
    for i := 0; i < n; i++ {
        arrayB[i] = temp
        temp *= arrayA[i]  // 更新temp为当前位置左边的所有元素的乘积
    }
    // 第二步：计算右边的积并与左边的积相乘
    temp = 1
    for i := n - 1; i >= 0; i-- {
        arrayB[i] *= temp  // 将右边的积乘到结果数组
        temp *= arrayA[i]  // 更新temp为当前位置右边的所有元素的乘积
    }
    return arrayB
}
