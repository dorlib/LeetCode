class Solution(object):
    def findTheDifference(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: str
        """
        s_dict = {}

        for i in range (len(s)):
            if s[i] in s_dict:
                s_dict[s[i]] += 1
            else:
                s_dict[s[i]] = 1
        
            
        for i in range (len(t)):
            if t[i] in s_dict:
                if s_dict[t[i]] < 1:
                    return t[i]
                else:
                    s_dict[t[i]] -= 1
            else:
                return t[i]
