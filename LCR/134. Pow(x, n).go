// 快速幂
func myPow(x float64, n int) float64 {
    // 处理指数为负的情况
    if n < 0 {
        x = 1 / x
        n = -n
    }
    result := 1.0
    base := x
    for n > 0 {
        // 如果当前指数是奇数，将当前的 base 乘入结果
        if n % 2 == 1 {
            result *= base
        }
        // 将 base 平方，指数减半
        base *= base
        n /= 2
    }
    return result
}
