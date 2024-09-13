// func divide(a int, b int) int { // 实现两个整数的除法运算
//     if a == math.MinInt32 { // 如果被除数是最小值
//         if b == 1 { // 除数是1
//             return math.MinInt32 // 结果仍然是最小值
//         }
//         if b == -1 { // 除数是-1
//             return math.MaxInt32 // 结果会导致溢出，返回最大值
//         }
//     }
//     if b == math.MinInt32 { // 如果除数是最小值
//         if a == math.MinInt32 { // 被除数也是最小值
//             return 1 // 结果是1
//         }
//         return 0 // 任何其他情况下结果是0
//     }
//     if a == 0 { 
//         // 如果被除数是0
//         return 0 // 结果是0
//     }

//     // 一般情况，使用二分查找
//     // 将所有的正数取相反数，这样就只需要考虑一种情况
//     rev := false // 用于记录结果的符号
//     if a > 0 { // 如果被除数是正数
//         a = -a // 转为负数
//         rev = !rev // 记录符号变化
//     }
//     if b > 0 { // 如果除数是正数
//         b = -b // 转为负数
//         rev = !rev // 记录符号变化
//     }

//     candidates := []int{b} // 初始化候选除数数组
//     for y := b; y >= a-y; { // 计算除数的倍数，防止溢出，因为都是负数了，所以是 y > x ( b > a)
//         y += y // 将当前倍数加倍
//         candidates = append(candidates, y) // 添加到候选数组
//     }

//     ans := 0 // 存储结果的商
//     for i := len(candidates) - 1; i >= 0; i-- { // 从最大的倍数开始
//         if candidates[i] >= a { // 如果当前倍数不超过被除数
//             ans |= 1 << i // 将对应的位设为1
//             a -= candidates[i] // 减去当前倍数
//         }
//     }
//     if rev { // 如果结果应该是负数
//         return -ans // 返回负值
//     }
//     return ans // 返回结果
// }


func divide(a int, b int) int {
    // 处理除数为0的情况
    if b == 0 {
        return math.MaxInt32 // 除数为0时返回一个极大的值表示错误
    }

    // 处理溢出情况
    if a == math.MinInt32 && b == -1 {
        return math.MaxInt32 // 被除数为最小值且除数为-1时，结果超出int范围
    }
    
    // 记录结果的符号
    negative := (a < 0) != (b < 0) // 如果被除数和除数符号不同，结果应为负数

    // 使用绝对值来简化计算
    a = abs(a)
    b = abs(b)

    // 用于存储结果的变量
    quotient := 0

    // 使用减法来计算商
    // 通过不断左移b来找到最大可能的倍数
    for a >= b {
        // temp是当前倍数的除数
        temp := b
        multiple := 1
        // 找到最大的temp，使得a - temp >= 0
        for a >= (temp << 1) {
            temp <<= 1
            multiple <<= 1
        }
        // 减去最大的temp，并将对应的倍数加到商中
        a -= temp
        quotient += multiple
    }

    // 根据结果符号调整商的符号
    if negative {
        return -quotient
    }
    return quotient
}

// 计算绝对值的辅助函数
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
