type RandomizedSet struct {
    values map[int]int // 值到索引的映射
    nums   []int       // 存储值的切片
}

/** 初始化数据结构 */
func Constructor() RandomizedSet {
    rand.Seed(time.Now().UnixNano()) // 为随机数生成器设置种子
    return RandomizedSet{
        values: make(map[int]int),
        nums:   []int{},
    }
}

/** 向集合中插入一个值。如果集合中没有该元素，则返回 true。 */
func (this *RandomizedSet) Insert(val int) bool {
    if _, exists := this.values[val]; exists {
        return false // 值已存在
    }
    // 将值添加到切片和映射中
    this.values[val] = len(this.nums)
    this.nums = append(this.nums, val)
    return true
}

/** 从集合中移除一个值。如果集合中包含该元素，则返回 true。 */
func (this *RandomizedSet) Remove(val int) bool {
    index, exists := this.values[val]
    if !exists {
        return false // 值不存在
    }
    // 用最后一个元素替代被移除元素的索引
    lastElement := this.nums[len(this.nums)-1]
    this.nums[index] = lastElement
    this.values[lastElement] = index // 更新最后一个元素的索引
    // 从切片中删除最后一个元素
    this.nums = this.nums[:len(this.nums)-1]
    delete(this.values, val) // 从映射中移除该值
    return true
}

/** 从集合中获取一个随机元素。 */
func (this *RandomizedSet) GetRandom() int {
    return this.nums[rand.Intn(len(this.nums))] // 返回一个随机元素
}

/**
 * 你的 RandomizedSet 对象将被实例化并按如下方式调用：
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
