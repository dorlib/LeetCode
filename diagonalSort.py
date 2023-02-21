class Solution(object):
    def diagonalSort(self, mat):
        """
        :type mat: List[List[int]]
        :rtype: List[List[int]]
        """
        for i in range(1, len(mat) - 1):
            j = 0
            k = i
            temp = []

            while j < len(mat[0]) and k < len(mat):
                temp.append(mat[k][j])
                j += 1
                k += 1
                    
            temp.sort()
            n = 0
            j = 0
            k = i

            print(temp)

            while j < len(mat[0]) and k < len(mat):
                mat[k][j] = temp[n]
                j += 1
                k += 1
                n += 1
        
        for i in range(len(mat[0]) - 1):
            j = i
            k = 0
            temp = []

            while j < len(mat[0]) and k < len(mat):
                temp.append(mat[k][j])
                j += 1
                k += 1
            
            temp.sort()
            n = 0
            j = i
            k = 0

            while j < len(mat[0]) and k < len(mat):
                mat[k][j] = temp[n]
                j += 1
                k += 1
                n += 1
        
        return mat
