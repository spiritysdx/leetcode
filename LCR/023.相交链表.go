/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    tempMap := map[*ListNode]struct{}{}
    for headA != nil {
        tempMap[headA] = struct{}{}
        headA = headA.Next
    }
    for headB != nil {
        if _, has := tempMap[headB]; has {
            return headB
        }
        headB = headB.Next
    }
    return nil
}
