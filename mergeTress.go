/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    
    newRoot := &TreeNode{}
    
    if root1 == nil {
        return root2
    }
    
    if root2 == nil {
        return root1
    }
    
    newRoot.Val = root1.Val + root2.Val
    newRoot.Left = mergeTrees(root1.Left, root2.Left)
    newRoot.Right = mergeTrees(root1.Right, root2.Right)
    
    return newRoot
}
