func singleNumber(nums []int) int {
    temp := make(map[int]int, len(nums))
    for _, v := range nums {
        temp[v]++
    }
    for i, v := range temp {
        if v == 1 {
            return i
        }
    }
    return -1
}
