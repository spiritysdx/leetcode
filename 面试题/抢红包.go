package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 10个人抢100元，每人最少抢1元
    count, amount := 10, 100
    remain := amount
    rand.Seed(time.Now().UnixNano()) // 初始化随机种子

    for i := 0; i < count; i++ {
        var x int
        if i == count-1 {
            // 最后一个人直接拿走剩余金额
            x = remain
        } else {
            // 其他人随机抢，保证剩余金额足够其他人抢到1元
            x = randomInt(1, remain-(count-i-1))
        }
        remain -= x
        fmt.Printf("第%d个人抢到了%d元\n", i+1, x)
    }
    fmt.Printf("剩余金额：%d元\n", remain)
}

// 生成一个在[min, max]范围内的随机整数
func randomInt(min, max int) int {
    return rand.Intn(max-min+1) + min
}

// package main

// import (
//     "fmt"
//     "math/rand"
//     "time"
// )

// func main() {
//     // 10个人抢100元，每人最少抢0.01元
//     count, amount := 10, 10000 // 金额以分为单位，100元 = 10000分
//     remain := amount
//     rand.Seed(time.Now().UnixNano()) // 初始化随机种子

//     for i := 0; i < count; i++ {
//         var x int
//         if i == count-1 {
//             // 最后一个人直接拿走剩余金额
//             x = remain
//         } else {
//             // 其他人随机抢，保证剩余金额足够其他人抢到1分（0.01元）
//             x = randomInt(1, remain-(count-i-1))
//         }
//         remain -= x
//         fmt.Printf("第%d个人抢到了%.2f元\n", i+1, float64(x)/100)
//     }
//     fmt.Printf("剩余金额：%.2f元\n", float64(remain)/100)
// }

// // 生成一个在[min, max]范围内的随机整数
// func randomInt(min, max int) int {
//     return rand.Intn(max-min+1) + min
// }
