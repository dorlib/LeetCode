class Solution(object):
    def numDifferentIntegers(self, word):
        """
        :type word: str
        :rtype: int
        """
        alphabet = "abcdefghijklmnopqrstuvwxyz"
        word = list(word)
        result = []

        for i in range(len(word)):
            if word[i] in alphabet:
                word[i] = '_'
        
        word = "".join(word)
        splitWord = re.split('\_*', word)

        for num in splitWord:
            if num not in result and num != "" and not self.checkIfEqaul(num, result):
                result.append(num)

        return len(result)

    def checkIfEqaul(self, num, lst):
        for number in lst:
            if int(number) == int(num):
                return True
        
        return False


