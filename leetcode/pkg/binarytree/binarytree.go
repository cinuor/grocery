package binarytree

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func PreorderTraverse(root *TreeNode) {
	fmt.Printf("%v ", root.Val)
	PreorderTraverse(root.Left)
	PreorderTraverse(root.Right)
}

// 124.二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {

	ans := math.MinInt32
	var maxGain func(*TreeNode) int

	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(0, maxGain(node.Left))
		right := max(0, maxGain(node.Right))

		// 更新最大值
		ans = max(ans, left+right+node.Val)

		// 找出最大的那一条路径
		return max(left, right) + node.Val

	}

	maxGain(root)
	return ans
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}

	i := 0

	for ; i < len(preorder); i++ {
		// 在中序遍历中找到跟节点的分界
		if inorder[i] == preorder[0] {
			break
		}
	}

	root.Left = buildTree(preorder[1:1+len(inorder[:i])], inorder[0:i])
	root.Right = buildTree(preorder[1+len(inorder[:i]):], inorder[i+1:])
	return root
}

// 100.比较两个二叉树
// 前序遍历
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 98.验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt64, math.MaxInt64)
}
func valid(r *TreeNode, min int, max int) bool {
	if r == nil {
		return true
	}
	if r.Val <= min || r.Val >= max {
		return false
	}
	return valid(r.Left, min, r.Val) && valid(r.Right, r.Val, max)
}

// 验证一个数是否在BST树中
func isInBST(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.Val == target {
		return true
	}

	if root.Val < target {
		return isInBST(root.Right, target)
	}

	return isInBST(root.Left, target)
}
