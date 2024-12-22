// https://leetcode.cn/problems/valid-triangle-number/description/
func triangleNumber(nums []int) int {
    // 排序数组
    sort.Ints(nums) // 使用标准库排序
    res := 0
    size := len(nums)
    // 枚举三角形的最长边 nums[k]
    for k := 2; k < size; k++ {
        i, j := 0, k-1 // 初始化两指针
        for i < j {
            if nums[i] + nums[j] > nums[k] {
                // 如果满足条件，说明 [i, j-1] 范围内的所有组合都满足条件
                res += j - i
                j-- // 收缩右边界
            } else {
                i++ // 增加左边界
            }
        }
    }
    return res
}


// func triangleNumber(nums []int) (ans int) {
//     slices.Sort(nums)
//     for k := 2; k < len(nums); k++ {
//         c := nums[k]
//         i := 0     // a=nums[i]
//         j := k - 1 // b=nums[j]
//         for i < j {
//             if nums[i]+nums[j] > c {
//                 // 由于 nums 已经从小到大排序
//                 // nums[i]+nums[j] > c 同时意味着：
//                 // nums[i+1]+nums[j] > c
//                 // nums[i+2]+nums[j] > c
//                 // ...
//                 // nums[j-1]+nums[j] > c
//                 // 从 i 到 j-1 一共 j-i 个
//                 ans += j - i
//                 j--
//             } else {
//                 // 由于 nums 已经从小到大排序
//                 // nums[i]+nums[j] <= c 同时意味着
//                 // nums[i]+nums[j-1] <= c
//                 // ...
//                 // nums[i]+nums[i+1] <= c
//                 // 所以在后续的内层循环中，nums[i] 不可能作为三角形的边长，没有用了
//                 i++
//             }
//         }
//     }
//     return ans
// }
