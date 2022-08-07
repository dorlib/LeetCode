# first solution - naive

class Solution(object):
    def countBits(self, n):
        """
        :type n: int
        :rtype: List[int]
        """
        ans = []
        for i in range(n+1):
            num = format(i,'b')
            ans.append(num.count('1'))
        
        return ans
    
# second solution - dp O(N)
class Solution(object):
    def countBits(self, n):
        """
        :type n: int
        :rtype: List[int]
        """
        dp = [0]
		
        for i in range(1, n + 1):
            exp = int(math.log(i) / math.log(2))
            num = 2**exp
            decimal = i % (num)
            if num == i:
                dp.append(1)
            else:
                dp.append(dp[num] + dp[decimal])
        return(dp)
