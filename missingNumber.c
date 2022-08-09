int missingNumber(int* nums, int numsSize){
    int lst[numsSize+1];
    memset( lst, 0, (numsSize+1)*sizeof(int));
    
    for (int i = 0; i < numsSize; i++) {
        lst[nums[i]] = nums[i];
    }
    
    for (int i = 1; i < numsSize + 1; i++) {
        if (lst[i] == 0) {
            return i;
        }
    }
    return 0;
}


// faster solution

int missingNumber(int *nums, int numsSize) {
    int total = 0;
    int n = 0;
    for (int i = 0; i < numsSize; i++) {
        total += nums[i];
        if (nums[i] > n) {
            n = nums[i];
        }
    }
        
    if (n < numsSize) {
        return numsSize;
    }
    
    return ((n*(n+1))/2) - total;
}
