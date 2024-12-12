func inventoryManagement(stock []int, cnt int) []int {
    if len(stock) == 0 || cnt == 0 {
        return nil
    }
    res := []int{}
    for cnt > 0 {
        minNum := 10001
        minIndex := 10001
        for index, key := range stock {
            minNum = min(minNum, key)
            if minNum == key {
                minIndex = index
            }
        }
        res = append(res, minNum)
        stock[minIndex] = 10001
        cnt--
    }
    return res
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}
