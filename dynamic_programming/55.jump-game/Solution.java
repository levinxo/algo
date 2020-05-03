/**
 * https://leetcode.com/problems/jump-game/
 * https://www.lintcode.com/problem/jump-game/description
 * 55. Jump Game
 */

public class Solution {


	public static void main(String[] args) {
		int[] nums = new int[]{2,3,1,1,4};
		System.out.println(canJump1(nums));
	}


	public static boolean canJump1(int[] nums) {
		/**
		 * 存在型的动态规划
		 * 1.确定状态：最后一步和子问题
		 * 最终位置设为j，且最后一跳从i起跳，那么i+nums[i] >= j
		 * 子问题变为是否能跳到i
		 * 2.转移方程
		 * dp[j] = (dp[i] AND i+nums[i] >= j)，其中0<=i<j
		 * 3.初始条件和边界情况
		 * dp[0] = true
		 * 4.计算顺序
		 * 从头开始
		 */

		if (nums == null || nums.length == 0) {
			return false;
		}

		int n = nums.length;
		boolean dp[] = new boolean[n];
		dp[0] = true;	// initialization

		for (int j = 1; j < n; ++j) {
			for (int i = 0; i < j; i++) {
				if (dp[i] && i+nums[i] >= j) {
					dp[j] = true;
					break;
				}
			}
		}

		return dp[n-1];
	}


}
