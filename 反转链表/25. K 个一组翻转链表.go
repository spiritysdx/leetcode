// https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
func reverseKGroup(head *ListNode, k int) *ListNode {
    n := 0
    current := head
    // 计算链表的长度
    for current != nil {
        n += 1
        current = current.Next
    }
    // 创建一个虚拟头节点，用于简化边界处理
    dummyNode := &ListNode{Next: head}
    p0 := dummyNode
    var pre *ListNode
    current = p0.Next
    // 按组进行反转，直到剩余节点不足k个
    for n >= k {
        n -= k
        // 执行k个节点的反转
        for i := 0; i < k; i++ {
            next := current.Next // 暂存下一个节点
            current.Next = pre   // 当前节点指向前一个节点，完成反转
            pre = current        // 更新前一个节点为当前节点
            current = next       // 移动到下一个节点
        }
        // 连接已反转部分和未反转部分
        next := p0.Next         // 暂存当前组的起点
        p0.Next.Next = current  // 反转后的最后一个节点指向下一部分
        p0.Next = pre           // p0指向反转后的头节点
        p0 = next               // 移动到下一组的起点
    }
    return dummyNode.Next // 返回新的头节点
}

// p0和dummy起初都指向同一个对象，也就是我们创建的哨兵节点。
// 第一次对p0.next的更新使得哨兵节点的next指向了第一组链表反转后的头节点，这时通过dummy.next就可以返回正确的链表头节点了
// 接着把p0更新为下一组反转链表的前一个节点，这时p0就不指向哨兵节点了，而dummy.next不受影响
