func cuttingBamboo(bamboo_len int) int {
    // 当竹子的长度拆分为3和2的乘积时可以获得最大的结果
    if bamboo_len <= 3 {
        // 如果竹子长度小于等于3，直接返回长度减1，因为不能再进行有效拆分
        return bamboo_len - 1
    }
    // 计算竹子长度被3整除的商
    quotient := bamboo_len / 3
    // 计算竹子长度被3整除的余数
    remainder := bamboo_len % 3
    if remainder == 0 {
        // 如果余数为0，直接返回3的商次方
        return int(math.Pow(3, float64(quotient)))
    } else if remainder == 1 {
        // 如果余数为1，将一个3变成2和2的组合，因此商减1，然后乘以4
        return int(math.Pow(3, float64(quotient - 1))) * 4
    }
    // 如果余数为2，直接在3的商次方结果上乘以2
    return int(math.Pow(3, float64(quotient))) * 2
}
