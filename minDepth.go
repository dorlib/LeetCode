/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
    "math"
)

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0 
    }
    return minDepthReq(root, 1)
}

func minDepthReq(root *TreeNode, sum int) int {
    
    var left int = sum
    var right int = sum 
    
    if root != nil && root.Left == nil && root.Right == nil {
        return sum
    }
    
    if root != nil && root.Left != nil {
        left = minDepthReq(root.Left, sum + 1)
    }
    
    if root != nil && root.Right != nil {
        right = minDepthReq(root.Right, sum + 1)
    }
    
    if root != nil && root.Left == nil {
        return right
    }
    
    if root != nil && root.Right == nil {
        return left
    }
    
    return int(math.Min(float64(right), float64(left)))
    
}
