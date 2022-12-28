/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if nums == nil || len(nums) < 1 {
        return nil
    }

    maxIndex := findMax(nums)
    var root TreeNode
    root.Val = nums[maxIndex]

    left := nums[0:maxIndex]
    right := nums[maxIndex + 1:]

    root.Left = constructMaximumBinaryTree(left)
    root.Right = constructMaximumBinaryTree(right)

    return &root
}

func findMax(nums []int) int {
    if len(nums) < 1 {
        return 0
    }

    max := nums[0]
    maxIndex := 0

    for i := 1; i < len(nums); i++ {
        if nums[i] > max {
            maxIndex = i;
            max = nums[i]
        }
    }

    return maxIndex
}
