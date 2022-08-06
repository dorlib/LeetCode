/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    
    if root != nil && root.Right == nil && root.Left == nil && (targetSum > root.Val || targetSum < root.Val) {
        return false
    }
    
    if root != nil && root.Right == nil && root.Left == nil && targetSum == root.Val {
        return true
    }
    
    if root == nil  {
        return false
    }
    
    var left bool 
    var right bool 
    
    if root != nil {
        left = hasPathSum(root.Left, targetSum - root.Val)
        right = hasPathSum(root.Right, targetSum - root.Val)
    }
    
    return left || right
    
}
