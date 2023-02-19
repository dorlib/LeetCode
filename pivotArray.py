class Solution(object):
    def pivotArray(self, nums, pivot):
        """
        :type nums: List[int]
        :type pivot: int
        :rtype: List[int]
        """
        smaller = []
        bigger = []
        equal = []

        for num in nums:
            if num > pivot:
                bigger.append(num)
            elif num < pivot:
                smaller.append(num)
            elif num == pivot:
                equal.append(num)
        
        return smaller + equal + bigger
        
