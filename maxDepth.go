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

func maxDepth(root *TreeNode) int {
    var max int 
    if root == nil {
        return 0
    }
    
    max = maxDepthReq(root, max)
    return max
}

func maxDepthReq(root *TreeNode, max int) int {
    if root == nil {
       return max 
    }
    
    if root.Left == nil && root.Right != nil {
        right := maxDepthReq(root.Right, max + 1)
        return right
    }
    
    if root.Right == nil && root.Left != nil {
        left := maxDepthReq(root.Left, max + 1)
        return left
    }
    
    left := maxDepthReq(root.Right, max + 1)
    right := maxDepthReq(root.Left, max + 1)
    
    return int(math.Max(float64(right), float64(left)))
} 
