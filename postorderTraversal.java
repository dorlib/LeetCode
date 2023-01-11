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
    public List<Integer> postorderTraversal(TreeNode root) {
        ArrayList<Integer> postOrderList = new ArrayList<>();
        postorderTraversalRec(root, postOrderList);
        return postOrderList;
    }

    public void postorderTraversalRec(TreeNode root, ArrayList<Integer> lst) {
        if (root == null) {
            return;
        }

        postorderTraversalRec(root.left, lst);
        postorderTraversalRec(root.right, lst);
        lst.add(root.val);
    }
}
