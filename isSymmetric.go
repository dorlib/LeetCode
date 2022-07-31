/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return false
    }
    
    return isSameTree(root.Left, root.Right)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    // cannot be true and true because we allready checked this 
    if p == nil || q == nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }    
    return isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
}
