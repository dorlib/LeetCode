class Solution(object):
    def findGCD(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        min = self.findMin(nums)
        max = self.findMax(nums)

        if min == max:
            return min

        return  self.gcd(min, max)

    
    def findMin(self, nums):
        if len(nums) == 0:
            return 

        min = nums[0]

        for i in range(len(nums)):
            if nums[i] < min:
                min = nums[i]
        
        return min

    def findMax(self, nums):
        if len(nums) == 0:
            return 

        max = nums[0]

        for i in range(len(nums)):
            if nums[i] > max:
                max = nums[i]
        
        return max

    def gcd(self, a, b): 
        if (b==0):
            return (a)
        else:
            return self.gcd(b, a % b);
