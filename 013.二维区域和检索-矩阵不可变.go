// NumMatrix 结构体定义，包含一个二维整数切片用于存储累积和
type NumMatrix struct {
    sums [][]int
}

// Constructor 函数用于初始化 NumMatrix 对象
func Constructor(matrix [][]int) NumMatrix {
    // 创建一个与输入矩阵行数相同的二维切片
    sums := make([][]int, len(matrix))
    for i, row := range matrix {
        // 对每一行，创建一个比原行多一列的切片（用于方便计算区域和）
        sums[i] = make([]int, len(row)+1)
        for j, v := range row {
            // 计算每个位置的累积和
            sums[i][j+1] = sums[i][j] + v
        }
    }
    // 返回初始化后的 NumMatrix 对象
    return NumMatrix{sums}
}

// SumRegion 方法用于计算指定矩形区域的和
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    var sum int
    // 遍历指定的行
    for i := row1; i <= row2; i++ {
        // 使用累积和数组计算每一行在指定列范围内的和，并累加到总和中
        sum += this.sums[i][col2+1] - this.sums[i][col1]
    }
    // 返回计算得到的区域和
    return sum
}

/**
 * NumMatrix 对象的实例化和调用方式如下：
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
