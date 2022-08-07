class Solution(object):
    def firstUniqChar(self, s):
        """
        :type s: str
        :rtype: int
        """
        
        dict = {}
        
        for char in s:
            if char not in dict:
                dict[char] = 1
            else:
                dict[char] += 1
                
        for char in s:
            if dict[char] == 1:
                return s.index(char)
        
        return -1
