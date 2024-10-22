/**
 * Definition for a Node.
 * type Node struct {
 *     Val   int
 *     Prev  *Node
 *     Next  *Node
 *     Child *Node
 * }
 */

// flatten 函数用于将多级双向链表扁平化
func flatten(root *Node) *Node {
	// 如果根节点为空，返回 nil
	if root == nil {
		return nil
	}

	// 深度优先搜索 (DFS) 函数
	var dfs func(*Node) (*Node, *Node)
	dfs = func(head *Node) (*Node, *Node) {
		// 如果当前节点没有下一个节点和子节点，返回当前节点作为头和尾
		if head.Next == nil && head.Child == nil {
			return head, head
		}

		var tail *Node = head // 初始化尾节点为当前节点

		// 处理下一个节点
		if head.Next != nil {
			nextHead, nextTail := dfs(head.Next) // 递归处理下一个节点
			head.Next = nextHead                  // 将当前节点的下一个指向处理后的头
			nextHead.Prev = head                  // 设置下一个节点的前驱为当前节点
			tail = nextTail                       // 更新尾节点
		}

		// 处理子节点
		if head.Child != nil {
			childHead, childTail := dfs(head.Child) // 递归处理子节点
			head.Child = nil                        // 清空当前节点的子节点
			childHead.Prev = head                   // 设置子节点的前驱为当前节点
			if head.Next != nil {
				head.Next.Prev = childTail // 设置下一个节点的前驱为子节点的尾节点
				childTail.Next = head.Next  // 将子节点的尾节点指向当前节点的下一个节点
			}
			head.Next = childHead // 将当前节点的下一个指向子节点的头
			tail = childTail      // 更新尾节点
		}

		return head, tail // 返回处理后的头和尾
	}

	root, _ = dfs(root) // 执行深度优先搜索
	return root          // 返回扁平化后的链表头
}
