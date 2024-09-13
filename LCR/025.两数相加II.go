/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 两数相加主函数
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // 反转链表 l1 和 l2
    l1 = ReverseChain(l1)
    l2 = ReverseChain(l2)

    // 初始化一个虚拟头节点
    dummy := &ListNode{}
    cur := dummy
    carry := 0 // 进位

    // 遍历两个链表，逐位相加
    for l1 != nil || l2 != nil || carry != 0 {
        sum := carry // 加上进位

        // 如果 l1 非空，加上 l1 当前节点的值
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }

        // 如果 l2 非空，加上 l2 当前节点的值
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        // 计算当前位的值和新的进位
        cur.Next = &ListNode{Val: sum % 10}
        carry = sum / 10 // 更新进位

        // 移动到下一个节点
        cur = cur.Next
    }

    // 最后将结果链表再反转回来
    return ReverseChain(dummy.Next)
}

// 反转链表的辅助函数
func ReverseChain(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head

    // 迭代反转链表
    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }

    // 返回反转后的链表头
    return prev
}
