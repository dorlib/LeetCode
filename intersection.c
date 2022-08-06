

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* intersection(int* nums1, int nums1Size, int* nums2, int nums2Size, int* returnSize){
    //initialize 2 arrays with 1001 places
    int numbers1[1001] = {0};
    int numbers2[1001] = {0};
    // allocating space
    int* res = calloc(1001, sizeof(int));
    
    // counting how many times each element apears in nums1 and in nums2
    // counting with indexes - if 3 apears 4 times in nums1, the value in index 3 in numbers1 will be 4
    for (int i = 0; i < nums1Size; i++) {
        numbers1[nums1[i]]++;
    }
    for (int j = 0; j < nums2Size; j++) { 
        numbers2[nums2[j]]++;
    }
    
    *returnSize = 0;
    
    int counter  = 0;
    for (int i = 0; i < 1001; i++) {
        //if in both lists the value in index i is more than 0, so the number of index appear in both nums1 and nums2
        if (numbers1[i] && numbers2[i]) {
            (*returnSize)++;
            res[counter] = i;
            counter++;
        }
    }
    return res;
}
