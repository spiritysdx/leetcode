type Checkout struct {
	maxPrice int
	size     int
	pMap     map[int]int
	p        []int
	maxDeque []int // 双端队列维护最大值
}

func Constructor() Checkout {
	return Checkout{
		maxPrice: -1,
		size:     0,
		pMap:     make(map[int]int),
		p:        []int{},
		maxDeque: []int{},
	}
}

func (this *Checkout) Get_max() int {
	if this.size == 0 {
		return -1
	}
	return this.maxPrice
}

func (this *Checkout) Add(value int) {
	this.p = append(this.p, value)
	this.size++
	// 更新最大值队列
	for len(this.maxDeque) > 0 && this.maxDeque[len(this.maxDeque)-1] < value {
		this.maxDeque = this.maxDeque[:len(this.maxDeque)-1]
	}
	this.maxDeque = append(this.maxDeque, value)
	this.maxPrice = this.maxDeque[0] // 更新当前最大值
}

func (this *Checkout) Remove() int {
	if this.size == 0 {
		return -1
	}
	// 移除队列首部元素
	removed := this.p[0]
	this.p = this.p[1:]
	this.size--
	// 更新最大值队列
	if len(this.maxDeque) > 0 && this.maxDeque[0] == removed {
		this.maxDeque = this.maxDeque[1:]
	}
	if this.size == 0 {
		this.maxPrice = -1 // 队列为空，最大值重置
	} else {
		this.maxPrice = this.maxDeque[0]
	}
	return removed
}

/**
 * Your Checkout object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get_max();
 * obj.Add(value);
 * param_3 := obj.Remove();
 */
