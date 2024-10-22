/** 
list.List 提供的常用方法：

    创建链表：
        list.New()：创建一个新的链表。

    添加元素：
        PushFront(v interface{})：在链表前面插入元素 v。
        PushBack(v interface{})：在链表后面插入元素 v。

    删除元素：
        Remove(e *Element)：删除指定的节点 e。

    移动元素：
        MoveToFront(e *Element)：将节点 e 移动到链表的前面。
        MoveToBack(e *Element)：将节点 e 移动到链表的后面。

    访问链表的首尾节点：
        Front()：返回链表的第一个节点（首节点）。
        Back()：返回链表的最后一个节点（尾节点）。

    遍历链表：
        你可以使用 e := l.Front() 获取链表的首节点，然后通过 e = e.Next() 遍历链表。

哨兵节点的作用

    简化边界条件的处理：
        在没有哨兵节点的双向链表中，处理链表的首节点和尾节点时需要特别考虑，例如在插入、删除操作时，要检查是否是空链表或只有一个节点的情况。这些边界情况的处理会导致额外的条件判断和复杂性。
        使用哨兵节点后，首节点和尾节点的插入、删除操作都可以统一对待，因为哨兵节点不存储实际数据，但始终存在于链表中。这意味着我们不需要在每次操作时判断链表是否为空，或者是否只有一个元素。

    减少空指针（Nil Pointer）的判断：
        如果没有哨兵节点，链表在一些操作中需要频繁判断是否为空节点（如 nil 检查），否则可能会导致空指针异常。
        哨兵节点始终存在，因此链表的 next 和 prev 指针总是指向一个有效的节点（即使是哨兵节点本身），避免了空指针错误。

    提高代码一致性：
        使用哨兵节点可以让代码更加一致，不需要针对不同节点（如头部、尾部）的操作写不同的逻辑，避免了复杂的条件分支。

什么时候需要实验哨兵节点？

    频繁操作链表头尾：
        当数据结构需要频繁进行首部或尾部的插入、删除操作时，使用哨兵节点可以简化操作逻辑。例如在双向链表中，每次插入时不必区分是否是空链表或者头部、尾部插入，可以一视同仁地处理所有节点。

    需要保持数据结构不变：
        在某些算法中，不希望链表为空（即使没有实际数据）。哨兵节点的存在保证链表始终有一个节点，减少了边界条件检查。

    提升代码的可维护性：
        当代码需要维护多个复杂条件（如头节点和尾节点的特殊处理）时，使用哨兵节点可以让代码变得简洁，易于理解和维护。

具体场景

以双向链表为例，哨兵节点通常用于链表头和尾的两端。例如，如果我们使用一个 head 哨兵节点，表示链表的开始部分，而 tail 哨兵节点表示链表的结束部分。在这种情况下：

    head.next 永远指向链表中的第一个有效节点，tail.prev 永远指向链表中的最后一个有效节点。
    在插入、删除时，始终只需要更新 next 和 prev，不需要额外处理空链表或单个节点的特殊情况。
*/

// 定义一个存储键值对的结构体 entry
type entry struct {
    key, value int
}

// 定义 LRUCache 结构体
type LRUCache struct {
    capacity  int               // 缓存容量
    list      *list.List        // 用于存储缓存内容的双向链表，按照访问顺序存储
    keyToNode map[int]*list.Element // 哈希表，存储键对应链表节点的映射
}

// 构造函数，初始化缓存
func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity:  capacity,             // 初始化缓存容量
        list:      list.New(),           // 创建一个新的双向链表
        keyToNode: map[int]*list.Element{}, // 初始化键到链表节点的映射
    }
}

// Get 方法，获取缓存中的值
func (c *LRUCache) Get(key int) int {
    node := c.keyToNode[key] // 从哈希表中找到 key 对应的链表节点
    if node == nil {         // 如果没有找到节点，表示缓存中没有这个 key
        return -1            // 返回 -1 表示未命中
    }
    c.list.MoveToFront(node) // 将该节点移到链表的最前面，表示它最近被访问过
    return node.Value.(entry).value // 返回节点存储的值
}

// Put 方法，往缓存中存储键值对
func (c *LRUCache) Put(key, value int) {
    if node := c.keyToNode[key]; node != nil { // 如果缓存中已经存在这个 key
        node.Value = entry{key, value} // 更新该节点的值
        c.list.MoveToFront(node)       // 将该节点移到链表的最前面，表示它最近被访问
        return                         // 直接返回
    }
    
    // 如果缓存中没有这个 key，则需要新增
    c.keyToNode[key] = c.list.PushFront(entry{key, value}) // 将新键值对插入链表最前面，并更新哈希表
    if len(c.keyToNode) > c.capacity { // 如果缓存已满
        // 移除链表最后面的节点（即最久未被访问的节点）
        lastNode := c.list.Back() // 获取链表最后一个节点
        lastEntry := c.list.Remove(lastNode).(entry) // 删除该节点并返回其存储的键值对
        delete(c.keyToNode, lastEntry.key) // 从哈希表中删除该键
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
