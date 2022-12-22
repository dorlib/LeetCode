# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def mergeNodes(self, head):
        newList = ListNode(0)
        current = newList
        sum = 0

        while head:
            if head.val != 0:
                sum += head.val
            else:
                if sum > 0:
                    current.next = ListNode(sum)
                    current = current.next
                    sum = 0
            head = head.next
        return newList.next
