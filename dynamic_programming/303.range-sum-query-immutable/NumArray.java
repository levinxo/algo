/**
 * https://leetcode.com/problems/range-sum-query-immutable/description/
 * 303. Range Sum Query - Immutable
 * 思路：
 * 利用动态规划先算出数组的和
 * 之后每次调用sumRange时再用区间值减去区间值即可
 */

class NumArray {

    private int[] dp;

    public NumArray(int[] nums) {
        dp = new int[nums.length];
        for (int i = 0; i < nums.length; i++) {
            if (i == 0) {
                dp[i] = nums[i];
            } else {
                dp[i] = dp[i - 1] + nums[i];
            }
        }
    }

    public int sumRange(int i, int j) {
        if (i == 0) {
            return dp[j];
        } else {
            return dp[j] - dp[i - 1];
        }
    }

    public static void main(String[] args) {
        int[] nums = new int[] { -2, 0, 3, -5, 2, -1 };
        NumArray numArray = new NumArray(nums);

        System.out.println(numArray.sumRange(0, 2));
        System.out.println(numArray.sumRange(2, 5));
        System.out.println(numArray.sumRange(0, 5));
    }

}

/**
 * Your NumArray object will be instantiated and called as such: NumArray obj =
 * new NumArray(nums); int param_1 = obj.sumRange(i,j);
 */