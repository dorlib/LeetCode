"""
# Definition for a Node.
class Node(object):
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""

class Solution(object):
    def postorder(self, root):
        """
        :type root: Node
        :rtype: List[int]
        """ 
        if root and root.children == []:
            return [root.val]

        result = []
        
        if root:
            for i in range(len(root.children)):
                result += self.postorder(root.children[i])
        
            return result + [root.val]
