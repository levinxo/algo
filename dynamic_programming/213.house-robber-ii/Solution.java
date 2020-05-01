/**
 * https://leetcode.com/problems/house-robber-ii/description/
 * 213. House Robber II
 */

import java.lang.Math;
import java.util.Arrays;

public class Solution {

    public static void main(String[] args) {
        if (args.length == 0) {
            return;
        }
        System.out.println(args.length + "个住户的最大抢劫量：");
        int[] nums = Arrays.asList(args).stream().mapToInt(Integer::parseInt).toArray();
        System.out.println(rob(nums));
    }

    public static int rob(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        int length = nums.length;
        if (length == 1) {
            return nums[0];
        }
        return Math.max(rob(nums, 0, length - 2), rob(nums, 1, length - 1));
    }

    public static int rob(int[] nums, int first, int last) {
        int prev2 = 0, prev1 = 0;
        for (int i = first; i <= last; i++) {
            int curr = Math.max(prev1, prev2 + nums[i]);
            prev2 = prev1;
            prev1 = curr;
        }
        return prev1;
    }

}