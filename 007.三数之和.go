// func threeSum(nums []int) [][]int {
//     if len(nums) < 3 {
//         return [][]int{}
//     }
//     sort.Ints(nums)
//     res := [][]int{}
//     var tp, otp, left, right, a, b int
//     for index, v := range nums {
//         if index > 0 && v == nums[index-1] {
//             continue
//         }
//         otp = 0 - v
//         left, right = index+1, len(nums)-1
//         for ; left < right; {
//             a = nums[left]
//             b = nums[right]
//             tp = a + b
//             if tp == otp {
//                 res = append(res, []int{v,a,b})
//                 left++
//             } else if tp > otp {
//                 right--
//             } else if tp < otp {
//                 left++
//             }
//         }
//     }
//     if len(res) == 1 {
//         return res
//     }
//     tres := [][]int{}
//     var status bool
//     for i, v := range res {
//         status = false
//         for _, s := range res[i+1:] {
//             if v[0] == s[0] && v[1] == s[1] && v[2] == s[2] {
//                 status = true
//                 break
//             }
//         }
//         if !status {
//             tres = append(tres, v)
//         } 
//     }
//     return tres
// }

func threeSum(nums []int) [][]int {
    n := len(nums)
    sort.Ints(nums)
    ans := make([][]int, 0)
 
    // 枚举 a
    for first := 0; first < n; first++ {
        // 需要和上一次枚举的数不相同
        if first > 0 && nums[first] == nums[first - 1] {
            continue
        }
        // c 对应的指针初始指向数组的最右端
        third := n - 1
        target := -1 * nums[first]
        // 枚举 b
        for second := first + 1; second < n; second++ {
            // 需要和上一次枚举的数不相同
            if second > first + 1 && nums[second] == nums[second - 1] {
                continue
            }
            // 需要保证 b 的指针在 c 的指针的左侧
            for second < third && nums[second] + nums[third] > target {
                third--
            }
            // 如果指针重合，随着 b 后续的增加
            // 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
            if second == third {
                break
            }
            if nums[second] + nums[third] == target {
                ans = append(ans, []int{nums[first], nums[second], nums[third]})
            }
        }
    }
    return ans
}
