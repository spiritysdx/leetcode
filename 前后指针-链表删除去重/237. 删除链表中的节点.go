// https://leetcode.cn/problems/delete-node-in-a-linked-list/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 将当前节点的下一个节点的值拷贝到当前节点，之后删掉当前节点的下一个节点
func deleteNode(node *ListNode) {
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}
