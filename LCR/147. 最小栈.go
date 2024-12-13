type MinStack struct {
    minNum int
    size int
    p []int   
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{size: 0}
}

func (this *MinStack) Push(x int)  {
    if this.size == 0 {
        this.minNum = x
    } else {
        this.minNum = min(this.minNum, x)
    }
    this.p = append(this.p, x)
    this.size += 1
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func (this *MinStack) Pop()  {
    if this.minNum == this.p[this.size-1] {
        this.minNum = this.p[0]
        if this.size >= 2 {
            for _, k := range this.p[1:this.size-1] {
                this.minNum = min(this.minNum, k)
            }
        }
    }
    this.p = this.p[:this.size-1]
    this.size -= 1
}

func (this *MinStack) Top() int {
    return this.p[this.size-1]
}

func (this *MinStack) GetMin() int {
    return this.minNum
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
