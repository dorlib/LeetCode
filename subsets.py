class Solution(object):
    def subsets(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        result = [[]]
        for num in nums:
            for i in range (len(result)):
                result.append([num] + result[i])

        return result



        
