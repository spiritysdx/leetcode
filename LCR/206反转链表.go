/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head
    // 迭代反转链表 - 双指针
    // 一个变量存当前指针，一个变量存实际未反转前的下一个指针
    for cur != nil {
        next := cur.Next       // 保存当前节点的下一个节点，防止断链
        cur.Next = prev        // 将当前节点的 Next 指向前一个节点，反转指针方向

        prev = cur             // 将 prev 移动到当前节点
        cur = next             // cur 移动到下一个节点，继续迭代
    }
    // 返回反转后的链表头
    return prev
}
