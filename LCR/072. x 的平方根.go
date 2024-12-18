func mySqrt(x int) int {
    // 如果输入为 0，则直接返回 0，因为 0 的平方根是 0
    if x == 0 {
        return 0
    }

    // 使用数学方法估算平方根
    // 首先将 x 转为 float64 类型，计算其对数的 0.5 倍，再取指数
    // math.Exp(0.5 * math.Log(float64(x))) 等价于 sqrt(x)
    ans := int(math.Exp(0.5 * math.Log(float64(x))))

    // 检查 ans + 1 是否是更接近的结果
    // 如果 (ans + 1) 的平方仍然小于等于 x，说明 ans + 1 是正确的整数平方根
    if (ans + 1) * (ans + 1) <= x {
        return ans + 1
    }

    // 否则返回 ans
    return ans
}
