/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
    var lst []string
    
    if root == nil {
        return lst
    }
    
    return binaryTreePathsReq(root, lst , "")
}

func binaryTreePathsReq(root *TreeNode, lst []string, path string ) []string {
    var res1 []string
    var res2 []string
    
    if root != nil && root.Right == nil && root.Left == nil {
        path = path + strconv.Itoa(root.Val)
        var result []string = append(lst, path)
        return result
    }
    
    if root != nil && root.Left != nil {
        res1 = binaryTreePathsReq(root.Left, lst, path + strconv.Itoa(root.Val) + "->")
    }
    
    if root != nil && root.Right != nil {
        res2 = binaryTreePathsReq(root.Right, lst, path + strconv.Itoa(root.Val) + "->")
    }
        
    
    return append(res1,res2...)
    
}
