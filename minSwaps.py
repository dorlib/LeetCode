class Solution(object):
    def minSwaps(self, s):
        """
        :type s: str
        :rtype: int
        """
        balance = 0
        answer = 0

        for char in s:
            if char == '[':
                balance += 1
            else:
                balance -= 1
            
            answer = min(answer, balance)

        return (-answer + 1) // 2
