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
