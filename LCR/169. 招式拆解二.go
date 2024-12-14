func dismantlingAction(arr string) byte {
    m := map[byte]int{} // 使用 byte 作为键，int 作为值
    arrs := []byte(arr)
    temp := []byte{}
    // 遍历字符串，记录字符出现次数并维护顺序
    for _, val := range arrs {
        if _, exist := m[val]; !exist {
            temp = append(temp, val) // 保留首次出现的顺序
        }
        m[val] += 1 // 更新字符计数
    }
    // 查找第一个只出现一次的字符
    for _, v := range temp {
        if m[v] == 1 {
            return v
        }
    }
    return ' '
}
