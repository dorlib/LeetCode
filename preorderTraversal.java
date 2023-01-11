/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public List<Integer> preorderTraversal(TreeNode root) {
        ArrayList<Integer> preorderList = new ArrayList<>();
        preorderTraversalRec(root, preorderList);
        return preorderList;
    }

    public void preorderTraversalRec(TreeNode root, ArrayList<Integer> lst) {
        if (root == null) {
            return;
        }

        lst.add(root.val);
        preorderTraversalRec(root.left, lst);
        preorderTraversalRec(root.right, lst);
    }
}
