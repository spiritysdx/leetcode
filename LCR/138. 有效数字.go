func validNumber(s string) bool {
    bytes := []byte(s)
    // 去掉前后的空格
    i, n := 0, len(bytes)
    for i < n && bytes[i] == ' ' {
        i++
    }
    for n > i && bytes[n-1] == ' ' {
        n--
    }
    // 如果去掉空格后为空字符串，直接返回 false
    if i == n {
        return false
    }
    // 检查是否有小数或整数部分
    isNumber := func(start, end int) bool {
        if start >= end {
            return false
        }
        // 检查符号 '+' 或 '-'
        if bytes[start] == '+' || bytes[start] == '-' {
            start++
        }
        hasDigits := false
        for start < end {
            if bytes[start] < '0' || bytes[start] > '9' {
                return false
            }
            hasDigits = true
            start++
        }
        return hasDigits
    }
    // 检查是否是小数
    isDecimal := func(start, end int) bool {
        if start >= end {
            return false
        }
        // 检查符号 '+' 或 '-'
        if bytes[start] == '+' || bytes[start] == '-' {
            start++
        }
        hasDigits, hasDot := false, false
        for start < end {
            if bytes[start] == '.' {
                if hasDot { // 多个点 '.' 不合法
                    return false
                }
                hasDot = true
            } else if bytes[start] >= '0' && bytes[start] <= '9' {
                hasDigits = true
            } else {
                return false
            }
            start++
        }
        // 小数至少需要有数字部分
        return hasDigits
    }
    // 检查是否含有 'e' 或 'E'
    for j := i; j < n; j++ {
        if bytes[j] == 'e' || bytes[j] == 'E' {
            // 分成两部分，分别检查小数和整数部分
            return (isDecimal(i, j) || isNumber(i, j)) && isNumber(j+1, n)
        }
    }
    // 没有 'e' 或 'E'，直接检查是否是小数或整数
    return isDecimal(i, n) || isNumber(i, n)
}
