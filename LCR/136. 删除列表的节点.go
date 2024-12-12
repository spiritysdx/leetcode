/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
    dummy := &ListNode{}
    current := dummy
    current.Next = head
    for current != nil && current.Next != nil {
        if current.Next.Val == val {
            current.Next = current.Next.Next
            break
        }
        current = current.Next
    }
    return dummy.Next
}
