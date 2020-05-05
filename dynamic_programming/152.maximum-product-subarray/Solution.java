/**
 * 152. Maximum Product Subarray
 * https://leetcode.com/problems/maximum-product-subarray/
 * https://www.lintcode.com/problem/maximum-product-subarray/description
 */
import java.lang.Math;


class Solution {


	public static void main(String[] args) {
		int[] nums = new int[]{2,3,-2,4};
		System.out.println(maxProduct(nums));
	}

	/**
	 * Time complexity O(n)
	 * 枚举了数组的长度
	 * Space complexity O(n)
	 * 消耗了等长的空间
	 */
	public static int maxProduct(int[] nums) {
		// 1.确定状态（最后一步和子问题）
		// subarray肯定有一个最后的元素i，考虑到元素i可能为正/负
		// 那么，可以求i-1的最大和最小值
		// 2.转移方程
		// maxDp[i] = max{maxDp[i-1]*nums[i], minDp[i-1]*nums[i], nums[i]}
		// minDp[i] = min{minDp[i-1]*nums[i], maxDp[i-1]*nums[i], nums[i]}
		// 3.初始值和边界情况
		// maxDp[0] = nums[0]
		// minDp[0] = nums[0]
		// 4.计算顺序
		// 从第0位开始

		if (nums.length == 0) {
			return 0;
		}

		int[] maxDp = new int[nums.length];
		int[] minDp = new int[nums.length];

		maxDp[0] = nums[0];		// initialization
		minDp[0] = nums[0];
		int result = nums[0];

		for (int i = 1; i < nums.length; i++) {
			int value = nums[i];
			maxDp[i] = Math.max(Math.max(maxDp[i-1]*value, minDp[i-1]*value), value);
			minDp[i] = Math.min(Math.min(maxDp[i-1]*value, minDp[i-1]*value), value);
			result = Math.max(result, maxDp[i]);
		}

		return result;
	}


}
