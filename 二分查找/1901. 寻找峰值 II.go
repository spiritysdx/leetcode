// https://leetcode.cn/problems/find-a-peak-element-ii/description/
// func findPeakGrid(mat [][]int) []int {
//     m := len(mat)
//     n := len(mat[0])
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             // 检查当前元素是否大于上下左右的相邻元素
//             isPeak := true
//             // 检查上方元素
//             if i > 0 && mat[i][j] <= mat[i-1][j] {
//                 isPeak = false
//             }
//             // 检查下方元素
//             if i < m-1 && mat[i][j] <= mat[i+1][j] {
//                 isPeak = false
//             }
//             // 检查左侧元素
//             if j > 0 && mat[i][j] <= mat[i][j-1] {
//                 isPeak = false
//             }
//             // 检查右侧元素
//             if j < n-1 && mat[i][j] <= mat[i][j+1] {
//                 isPeak = false
//             }
//             // 如果是峰值，返回坐标
//             if isPeak {
//                 return []int{i, j}
//             }
//         }
//     }
//     // 如果没有找到峰值（实际上根据题意一定会有峰值）
//     return []int{0, 0}
// }

func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    n := len(mat[0])
    // 对列进行二分搜索
    left, right := 0, n-1
    for left <= right {
        mid := left + (right-left)/2
        // 找到当前列的最大值所在行
        maxRow := 0
        for i := 0; i < m; i++ {
            if mat[i][mid] > mat[maxRow][mid] {
                maxRow = i
            }
        }
        // 检查左边的值（如果存在）
        leftIsBigger := mid > 0 && mat[maxRow][mid-1] > mat[maxRow][mid]
        // 检查右边的值（如果存在）
        rightIsBigger := mid < n-1 && mat[maxRow][mid+1] > mat[maxRow][mid]
        if !leftIsBigger && !rightIsBigger {
            // 找到峰值
            return []int{maxRow, mid}
        } else if leftIsBigger {
            // 左边有更大的值，往左边搜索
            right = mid - 1
        } else {
            // 右边有更大的值，往右边搜索
            left = mid + 1
        }
    }
    // 根据题意一定存在峰值，不会达到这里
    return []int{0, 0}
}
