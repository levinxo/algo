/**
 * 540. Single Element in a Sorted Array
 * https://leetcode.com/problems/single-element-in-a-sorted-array/
 */

public class Solution {


    public static void main(String[] args){
        int[] nums = new int[]{1,1,2,3,3,4,4,8,8};
        System.out.println(singleNonDuplicate(nums));
    }


    public static int singleNonDuplicate(int[] nums) {
        int l = 0, r = nums.length - 1;
        while (l < r) {
            int m = l + (r - l) / 2;
            if ((m & 1) == 1) {
                // 判断index是奇数，就减1，保持index为偶数
                // 此处不会越界，因为开始的index=0为偶数
                m--;
            }
            if (nums[m] == nums[m + 1]) {
                // 此处不会越界，因为在这个循环内部l != r
                l = m + 2;
            } else {
                r = m;
            }
        }
        return nums[l];
    }


}
