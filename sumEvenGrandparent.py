# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def sumEvenGrandparent(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if root == None:
            return 0
        
        if root.left == None and root.right == None:
            return 0
        
        res = 0
        if root.val % 2 == 0:

            if root.left:
                if root.left.left:
                    res += root.left.left.val
                if root.left.right:
                    res += root.left.right.val
            
            if root.right:
                if root.right.right:
                    res += root.right.right.val
                if root.right.left:
                    res += root.right.left.val

        return res + self.sumEvenGrandparent(root.left) + self.sumEvenGrandparent(root.right)      

            
            
            
