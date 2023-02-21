public class Solution {
    public int FindLengthOfLCIS(int[] nums) {
        if (nums.Length == 0) {
            return 0;
        }

        int length=1;
        int temp=1;

        for(int i=0; i < nums.Length-1; i++) {
            if(nums[i] < nums[i+1]) {
                temp++; 
                length = Math.Max(length, temp);
            } else {
                temp = 1;
            }
        }

        return length;
    }
}
