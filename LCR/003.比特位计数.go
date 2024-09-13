func countBits(n int) []int {
    // 创建一个大小为 n+1 的整型切片，并初始化为 0
    bits := make([]int, n+1)
    // 从 1 遍历到 n
    for i := 1; i <= n; i++ {
        // bits[i>>1] 表示 i 右移一位后的值的 bit 数量
        // i&1 表示 i 的最低位（如果是 1，则 +1；如果是 0，则 +0）
        // 两者相加即为 bits[i] 的值
        bits[i] = bits[i>>1] + i&1
    }
    // 返回包含每个数的二进制表示中 1 的数量的切片
    return bits
}
