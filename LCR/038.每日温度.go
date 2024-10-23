// 无中间值记录

// func dailyTemperatures(temperatures []int) []int {
//     res := []int{}
//     var status bool
//     for i1, v1 := range temperatures {
//         status = false
//         for i2, v2 := range temperatures[i1:] {
//             if v2 > v1 {
//                 res = append(res, i2)
//                 status = true
//                 break
//             }
//         }
//         if !status {
//             res = append(res, 0)
//         }
//     }
//     return res
// }
