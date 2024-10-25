package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// multiply 字符串形式的乘法
func multiply(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	// 用来存放结果的数组
	res := make([]int, len(num1)+len(num2))
	// 从后向前逐位相乘
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			digit1 := int(num1[i] - '0')
			digit2 := int(num2[j] - '0')
			product := digit1 * digit2 // 当前位的运算结果
			// 位置索引
			p1 := i + j
			p2 := i + j + 1
			// 累加到结果数组
			sum := product + res[p2] // 处理进位
			res[p2] = sum % 10  // 进位
			res[p1] += sum / 10 // 当前位
		}
	}
	// 构建结果字符串
	var sb strings.Builder
	for _, num := range res {
		if !(sb.Len() == 0 && num == 0) { // 跳过前导零
			sb.WriteByte(byte(num + '0')) // 将有效数字转换为字符并添加到 sb 中
		}
	}
	return sb.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入第一个数字: ")
	num1, _ := reader.ReadString('\n')
	num1 = strings.TrimSpace(num1) // 去除换行符和空格
	fmt.Print("请输入第二个数字: ")
	num2, _ := reader.ReadString('\n')
	num2 = strings.TrimSpace(num2) // 去除换行符和空格
	result := multiply(num1, num2)
	fmt.Printf("%s * %s = %s\n", num1, num2, result)
}
