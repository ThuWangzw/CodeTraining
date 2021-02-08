/*
    异或操作满足：交换律结合律分配律
 */
class Solution {
    public int singleNumber(int[] nums) {
        int num = nums[0];
        for(int i=1; i<nums.length; i++) {
            num = num ^ nums[i];
        }
        return num;
    }
}