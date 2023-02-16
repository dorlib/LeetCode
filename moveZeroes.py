class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """

        arrLength = len(nums)
        counter = 0

        for i in range(arrLength):
            if nums[i] == 0:
                counter += 1
            else:
                nums[i-counter] = nums[i]
            
        for i in range(counter):
            nums[arrLength - i - 1] = 0

        return nums
