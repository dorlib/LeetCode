class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        maximum = 0
        j = 0
        
        for i in range(1, len(prices)):
            maximum =  max(maximum,prices[i]-prices[j])
            if prices[i] < prices[j]:
                j = i
        return maximum
            
