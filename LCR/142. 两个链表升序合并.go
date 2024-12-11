/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val  int
 *     Next *ListNode
 * }
 */
func trainningPlan(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{} // 使用虚拟节点做头节点避免处理头节点情况
    current := dummy
    for l1 != nil && l2 != nil { // 循环主体是目前已拼接好的链表
        if l1.Val < l2.Val {
            current.Next = l1
            l1 = l1.Next
        } else {
            current.Next = l2
            l2 = l2.Next
        }
        current = current.Next
    }
    if l1 != nil {
        current.Next = l1
    } else if l2 != nil {
        current.Next = l2
    }
    return dummy.Next
}
