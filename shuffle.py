class Solution(object):
    def shuffle(self, nums, n):
        """
        :type nums: List[int]
        :type n: int
        :rtype: List[int]
        """
        lst1 = nums[:n]
        lst2 = nums[n:]
        
        i = 0
        j = 0
        res = [0] * 2 * n
        while i < n:
            res[j] = lst1[i]
            res[j+1] = lst2[i]
            i += 1
            j += 2
        return res
