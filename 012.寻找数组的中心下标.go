// func pivotIndex(nums []int) int {
//     size := len(nums)
//     if size == 0 {
//         return -1
//     }
//     var sum, left, right int
//     for i := 0; i <= size-1; {
//         sum, left, right = 0, 0, 0
//         if i == 0 {
//             for _, j := range nums[1:] {
//                 sum += j
//             }
//             if sum == 0 {
//                 return 0
//             }
//         } else if i == size -1 {
//             for _, j := range nums[:size-1] {
//                 sum += j
//             }
//             if sum == 0 {
//                 return size-1
//             }
//         } else {
//             for _, j := range nums[:i] {
//                 left += j
//             }
//             for _, j := range nums[i+1:] {
//                 right += j
//             }
//             if left == right {
//                 return i
//             }
//         }
//         i++
//     }
//     return -1
// }

func pivotIndex(nums []int) int {
//利用前缀和
sum := 0
pre := 0
//将数组的值全部加起来
for _,v := range nums {
    sum += v
}
for i, v := range nums {
//就是说明左边的和右边的加上中间值等于数组的总和
    if (2*pre + v) == sum {
        return i
    }
    pre += v
}
return -1
}
