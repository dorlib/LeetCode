/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {
    result := 0
    averageOfSubtreeRec(root, &result)
    return result
}

func averageOfSubtreeRec(root *TreeNode, val *int) []int {
    if root == nil {
        return []int{0, 0}
    }

    left := averageOfSubtreeRec(root.Left, val)
    right := averageOfSubtreeRec(root.Right, val)

    sum := right[0] + root.Val + left[0]
    numOfNodes := right[1] + 1 + left[1]

    if sum / numOfNodes == root.Val {
        *val++
    }

    return []int{sum, numOfNodes}
}

