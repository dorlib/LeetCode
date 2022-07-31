/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var lst = make([]int, 0)
    if root == nil {
        return lst
    }
    
    lst = append(lst, inorderTraversal(root.Left)...)
    lst = append(lst, root.Val)
    lst = append(lst, inorderTraversal(root.Right)...)
    
    return lst
}
