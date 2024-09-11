// func reverse(nums []int, size int) []int {
//     for i, j := 0, size-1; i < j; {
//         nums[i], nums[j] = nums[j], nums[i]
//         i++
//         j--
//     }
//     return nums
// }

// func checknums(nums []int) int {
//     c1, c2 := 0, 0
//     ans := 0
//     for index, n := range nums {
//         if n == 0 {
//             c1 += 1
//         } else {
//             c2 += 1
//         }
//         if c1 == c2 {
//             ans = index + 1
//         }
//     }
//     return ans
// }

// func findMaxLength(nums []int) int {
//     size := len(nums)
//     if size == 0 {
//         return 0
//     }
//     ans0 := 0
//     for i,j := 0, size; i < j; {
//         ans1 := checknums(nums[i:j])
//         ans2 := checknums(reverse(nums[i:j], size))
//         if ans1 > ans0 && ans1 >= ans2 {
//             ans0 = ans1
//             if ans0 == size {
//                 break
//             }
//         } else if ans2 > ans0 {
//             ans0 = ans2
//             if ans0 == size {
//                 break
//             }
//         }
//         i++
//         j--
//         size -= 2
//     }
//     return ans0
// }

// 暴力方法会超时

func max(a, b int) int {
    if a > b { // 如果 a 大于 b
        return a // 返回 a
    }
    return b // 否则返回 b
}

func findMaxLength(nums []int) int {
    ans := 0 // 初始化答案变量为 0
    mp := map[int]int{0: -1} // 初始化哈希表，记录 counter 值对应的索引，初始值为 0:-1 
    counter := 0 // 初始化计数器为 0
    for i, num := range nums { // 遍历数组，i 是索引，num 是当前值
        if num == 1 { // 如果当前值是 1
            counter++ // 计数器加 1
        } else { // 如果当前值是 0
            counter-- // 计数器减 1
        }
        prevIndex, has := mp[counter] // 获取记录过的数组长度的map的值
        if has { // 如果当前计数器值之前已经出现过
            ans = max(ans, i-prevIndex) // 更新答案为当前索引减去之前出现该计数器值的索引与当前答案的最大值
        } else { // 如果当前计数器值之前没有出现过
            mp[counter] = i // 将当前计数器值和索引存入哈希表
        }
    }
    return ans // 返回答案
}
