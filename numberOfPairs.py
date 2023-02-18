class Solution(object):
    def numberOfPairs(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        num_of_pairs = 0
        num_of_elements_left = len(nums)

        dict = {}

        for i in range(len(nums)):
            if nums[i] not in dict:
                dict[nums[i]] = 1
            else:
                dict[nums[i]] += 1
        
        for key in dict:
            while dict[key] > 1:
                dict[key] -= 2
                num_of_pairs += 1
                num_of_elements_left -= 2


        return [num_of_pairs, num_of_elements_left]


