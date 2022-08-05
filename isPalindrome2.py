class Solution(object):
    def isPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """
        
        str = ""
        valid = "abcdefghijklmnopqrstuvwxyz0123456789"
        
        for i in range(len(s)):
            if s[i].lower() in valid:
                str += s[i].lower()
                               
        return str == str[::-1]
