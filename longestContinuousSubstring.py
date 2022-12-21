class Solution(object):
    def longestContinuousSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        res = 1
        currentLength =  1

        for i in range(1, len(s)):
            if ord(s[i]) - ord(s[i-1]) == 1:
                currentLength += 1
            else:
                currentLength = 1

            res = max(res, currentLength)

        return res
