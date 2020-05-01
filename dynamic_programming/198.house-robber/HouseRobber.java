/**
 * https://leetcode.com/problems/house-robber/description/
 * 198. House Robber
 * 定义数组dp来存储最大抢劫量，其中dp[i]表示抢到第i个住户时最大抢劫量
 * 情况分析：如果抢劫了dp[i]，那么dp[i-1]就不能抢了，可得转移方程：
 * dp[i] = nums[i] + dp[i-2]
 * dp[i] = dp[i-1]
 */

import java.lang.Math;
import java.util.Arrays;

public class HouseRobber {

    public static void main(String... args) {
        if (args.length == 0) {
            return;
        }
        System.out.println(args[0] + "个住户的最大抢劫量：");
        int[] nums = Arrays.asList(args).stream().mapToInt(Integer::parseInt).toArray();
        System.out.println(rob(nums));
    }

    public static int rob(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        int prev2 = 0, prev1 = 0;
        for (int i = 0; i < nums.length; i++) {
            int curr = Math.max(prev2 + nums[i], prev1);
            prev2 = prev1;
            prev1 = curr;
        }
        return prev1;
    }

}