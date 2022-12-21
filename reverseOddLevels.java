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
    public TreeNode reverseOddLevels(TreeNode root) {
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);

        for (int level = 0; !queue.isEmpty(); level++) {
            int size = queue.size();
            TreeNode[] arr = new TreeNode[size];

            for (int i = 0; i < size; i++) {
                arr[i] = queue.poll();
                if (arr[i].left != null) {
                    queue.add(arr[i].left);
                }
                if (arr[i].right != null) {
                    queue.add(arr[i].right);
                }
            }

            if (level % 2 != 0) {
                reverse(arr);
            }
        }

        return root;
    }

    void reverse(TreeNode[] arr){
        for(int i=0, j=arr.length-1; i<j; i++, j--){
            int temp = arr[i].val;
            arr[i].val = arr[j].val;
            arr[j].val = temp;
        }
    }
}
