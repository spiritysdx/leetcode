# 最近公共祖先

node == nil || p == node.Val || q == node.Val 返回本身

left和right递归得到结果

都不空返回node，否则返回不空的那一边，都空返回nil

二叉搜索树就不需要left和right都递归，判断值在哪边递归哪边
