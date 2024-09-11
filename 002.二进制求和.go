func addBinary(a string, b string) string {
    // 补齐零位
    // 如果 a 比 b 长，则在 b 前面补足相应的 '0'
    if len(a) > len(b) {
        b = strings.Repeat("0", len(a)-len(b)) + b
    // 如果 b 比 a 长，则在 a 前面补足相应的 '0'
    } else if len(a) < len(b) {
        a = strings.Repeat("0", len(b)-len(a)) + a
    }
    // 声明等长的 byte 数组用于存储结果
    res := make([]byte, len(a))
    // 判断进位值（逢二进一）
    var carry int
    // 定义是否要在结果前插入前导'1'的标志
    flag := false
    // 从末位开始逆序遍历
    for i := len(a)-1; i >= 0; i-- {
        // 计算当前位置的和，包括进位
        carry = int(a[i]-'0') + int(b[i]-'0')
        if carry == 2 {
            // 如果当前位置已经进位变为'1'，不用管，直接置为'0'
            if res[i] != '1' {
                res[i] = '0'
            }
            if i-1 >= 0 {  // 如果不是最前面的位，进位
                res[i-1] = '1'
            } else {  // 如果是最前面的位，需要插入前导'1'
                flag = true
            }
        } else if carry == 1 {
            // 如果当前位置已经是'1'，则表示需要再次进位
            if res[i] == '1' {
                res[i] = '0'
                if i-1 >= 0 {  // 如果不是最前面的位，进位
                    res[i-1] = '1'
                } else {  // 如果是最前面的位，需要插入前导'1'
                    flag = true
                }
            } else {
                res[i] = '1'
            }
        } else {
            // 如果和为 0，且当前位置还未被进位影响，则置为'0'
            if res[i] != '1' {
                res[i] = '0'
            }
        }
    }
    // 根据是否需要前导'1'来返回结果
    if flag {
        return "1" + string(res)
    } else {
        return string(res)
    }
}
