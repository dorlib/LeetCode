/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    return rangeSumBSTRec(root, low, high)
   
}

func rangeSumBSTRec(root *TreeNode, low int, high int) int {
    if root == nil {
        return 0
    }

    if root.Val < low || root.Val > high {
        return rangeSumBSTRec(root.Left, low, high) + rangeSumBSTRec(root.Right, low, high)
    }


    return rangeSumBSTRec(root.Left, low, high) + root.Val + rangeSumBSTRec(root.Right, low, high)
}
