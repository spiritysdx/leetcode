/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func trainningPlan(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var current = head
    var pre *ListNode
    for current != nil {
        next := current.Next
        current.Next = pre
        pre = current
        current = next
    }
    return pre
}
