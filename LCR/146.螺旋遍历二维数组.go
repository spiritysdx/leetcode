func spiralArray(array [][]int) []int {
    res := []int{}
    if len(array) == 0 || len(array[0]) == 0 {
		return res // 如果二维数组为空，直接返回空数组
	}
    // 定义 上 下 左 右 的边界
    head, end, left, right := 0, len(array)-1, 0, len(array[0])-1
    // 当边界仍然有效时进行遍历
    for head <= end && left <= right {
        // 从左到右
        for i := left; i <= right; i++ {
            res = append(res, array[head][i])
        }
        head++
        // 从上到下
        for i := head; i <= end; i++ {
            res = append(res, array[i][right])
        }
        right--
        if head <= end {
            // 从右到左
            for i := right; i >= left; i-- {
                res = append(res, array[end][i])
            }
            end--
        }
        if left<=right {
            // 从下到上
            for i := end; i >= head; i-- {
                res = append(res, array[i][left])
            }
            left++
        }
    }
    return res
}
