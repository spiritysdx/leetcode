// https://leetcode.cn/problems/add-two-numbers-ii/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // 翻转链表，并获取链表长度
    l1, n1 := reverse(l1)
    l2, n2 := reverse(l2)
    // 确定最长链表的长度
    nMax := max(n1, n2)
    dummy := &ListNode{} // 哑节点，用于构造结果链表
    cur := dummy         // 当前节点
    carry := 0           // 进位值
    // 遍历两个链表，逐位相加
    for i := 0; i < nMax || carry > 0; i++ {
        sum := carry
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        carry = sum / 10 // 计算进位
        cur.Next = &ListNode{Val: sum % 10}
        cur = cur.Next
    }
    // 再次翻转链表，恢复正确的顺序
    pre, _ := reverse(dummy.Next)
    return pre
}

// 翻转链表，并返回翻转后的链表和长度
func reverse(L *ListNode) (*ListNode, int) {
    var pre *ListNode
    cur := L
    n := 0 // 记录链表长度
    for cur != nil {
        nxt := cur.Next
        cur.Next = pre
        pre = cur
        cur = nxt
        n++
    }
    return pre, n
}

// 返回两个整数中的较大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
