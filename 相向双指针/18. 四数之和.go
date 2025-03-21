// https://leetcode.cn/problems/4sum/
func fourSum(nums []int, target int) [][]int {
    // 首先对数组进行排序
    sort.Ints(nums)
    res := [][]int{}
    size := len(nums)
    // 遍历数组，外层两层循环固定前两个数
    for i := 0; i < size-3; i++ {
        if i > 0 && nums[i] == nums[i-1] { // 跳过重复元素
            continue
        }
        for j := i + 1; j < size-2; j++ {
            if j > i+1 && nums[j] == nums[j-1] { // 跳过重复元素
                continue
            }
            // 双指针解决后两个数
            left, right := j + 1, size - 1
            for left < right {
                sum := nums[i] + nums[j] + nums[left] + nums[right]
                if sum == target {
                    res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
                    // 跳过重复的第三个数和第四个数
                    for left < right && nums[left] == nums[left+1] {
                        left++
                    }
                    for left < right && nums[right] == nums[right-1] {
                        right--
                    }
                    left++
                    right--
                } else if sum < target {
                    left++
                } else {
                    right--
                }
            }
        }
    }
    return res
}

// func fourSum(nums []int, target int) [][]int {
//     // 首先对数组进行排序
//     sort.Ints(nums)
//     res := [][]int{}
//     size := len(nums)
//     // 遍历数组，外层两层循环固定前两个数
//     for a := 0; a < size-3; a++ {
//         if a > 0 && nums[a] == nums[a-1] { // 跳过重复元素
//             continue
//         }
//         // 提前剪枝优化
//         if nums[a]+nums[a+1]+nums[a+2]+nums[a+3] > target {
//             break
//         }
//         if nums[a]+nums[size-3]+nums[size-2]+nums[size-1] < target {
//             continue
//         }
//         for b := a + 1; b < size-2; b++ {
//             if b > a+1 && nums[b] == nums[b-1] { // 跳过重复元素
//                 continue
//             }
//             // 提前剪枝优化
//             if nums[a]+nums[b]+nums[b+1]+nums[b+2] > target {
//                 break
//             }
//             if nums[a]+nums[b]+nums[size-2]+nums[size-1] < target {
//                 continue
//             }
//             // 双指针解决后两个数
//             left, right := b + 1, size - 1
//             for left < right {
//                 sum := nums[a] + nums[b] + nums[left] + nums[right]
//                 if sum == target {
//                     res = append(res, []int{nums[a], nums[b], nums[left], nums[right]})
//                     // 跳过重复的第三个数和第四个数
//                     for left < right && nums[left] == nums[left+1] {
//                         left++
//                     }
//                     for left < right && nums[right] == nums[right-1] {
//                         right--
//                     }
//                     left++
//                     right--
//                 } else if sum < target {
//                     left++
//                 } else {
//                     right--
//                 }
//             }
//         }
//     }
//     return res
// }
