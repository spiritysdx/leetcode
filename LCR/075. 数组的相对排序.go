
func relativeSortArray(arr1 []int, arr2 []int) []int {
    temp1 := map[int]int{}          // 用于存储 arr2 中每个元素的顺序
    temp2 := map[int][]int{}        // 用于存储按照 arr2 顺序分组的 arr1 中的元素
    temp3 := []int{}                // 用于存储 arr1 中不在 arr2 中的元素

    // 遍历 arr2，将元素的顺序存储到 temp1 中
    for index, value := range arr2 {
        temp1[value] = index        // 将 arr2 的值作为键，索引作为值
    }

    // 遍历 arr1，将元素分为在 arr2 中的和不在 arr2 中的两部分
    for _, key := range arr1 {
        if value, ok := temp1[key]; ok { // 如果 arr1 的值在 arr2 中
            if _, exists := temp2[value]; exists { // 如果 temp2 已经有该分组
                temp2[value] = append(temp2[value], key) // 添加到对应分组
            } else {
                temp2[value] = []int{key} // 初始化对应分组
            }
        } else {
            temp3 = append(temp3, key) // 不在 arr2 中的元素加入 temp3
        }
    }

    sort.Ints(temp3) // 对 temp3 中不在 arr2 的元素进行升序排序

    res := []int{}   // 存储最终结果

    // 按照 arr2 的顺序合并结果
    for i := 0; i < len(arr2); i++ {
        if v, ok := temp2[i]; ok { // 如果 temp2 中存在分组
            res = append(res, v...) // 添加该分组的所有元素
        }
    }

    res = append(res, temp3...) // 将排序后的 temp3 中元素追加到结果中
    return res                  // 返回最终结果
}
