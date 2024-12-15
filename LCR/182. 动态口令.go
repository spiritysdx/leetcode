func dynamicPassword(password string, target int) string {
    passwordBytes := []byte(password)
    res1 := []byte{}
    res2 := []byte{}
    for index, v := range passwordBytes {
        if index + 1 <= target {
            res1 = append(res1, v)
        } else {
            res2 = append(res2, v)
        }
    }
    res2 = append(res2, res1...)
    return string(res2)
}

// func dynamicPassword(s string, n int) string {
//     length := len(s)
//     b := make([]byte, length) // 辅助数组
//     idx := 0
//     // 写入尾部字符串
//     for i := n; i < length; i++ {
//         b[idx] = s[i]
//         idx++
//     }
//     // 写入头部字符串
//     for i := 0; i < n; i++ {
//         b[idx] = s[i]
//         idx++
//     }
//     return string(b)
// }
