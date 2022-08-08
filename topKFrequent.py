class Solution(object):    
    def topKFrequent(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        dict = {}
        
        for num in nums:
            if num in dict:
                dict[num] += 1
            else:
                dict[num] = 1
                
        return nlargest(k, dict, key=dict.get)
