func cuttingBamboo(bamboo_len int) int {
    const MOD = 1000000007 // 定义模数
    // 快速幂函数，用于计算 (base^exp) % mod
    fastPow := func(base, exp, mod int) int {
        res := 1 // 初始化结果为1，因为任何数的0次幂都是1
        for exp > 0 { // 当指数大于0时，持续循环
            if exp%2 == 1 { 
                // 如果指数是奇数，将当前的base乘到结果中，并取模
                res = (res * base) % mod 
            }
            // 将base平方，计算下一步的幂，且取模，防止溢出
            base = (base * base) % mod 
            // 指数除以2（相当于指数右移一位）
            exp /= 2 
        }
        // 返回最终结果
        return res
    }
    // 如果竹子长度小于等于3，直接返回长度减1
    if bamboo_len <= 3 {
        return bamboo_len - 1
    }
    // 计算商和余数
    q1 := bamboo_len / 3
    q2 := bamboo_len % 3
    var res int
    if q2 == 0 {
        // 如果余数为0，结果为3^q1 % MOD
        res = fastPow(3, q1, MOD)
    } else if q2 == 1 {
        // 如果余数为1，将一个3变为2和2的组合，结果为 (3^(q1-1) * 4) % MOD
        res = (fastPow(3, q1-1, MOD) * 4) % MOD
    } else if q2 == 2 {
        // 如果余数为2，结果为 (3^q1 * 2) % MOD
        res = (fastPow(3, q1, MOD) * 2) % MOD
    }
    // 返回最终结果
    return res
}
