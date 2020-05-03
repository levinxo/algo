/**
 * https://leetcode.com/problems/coin-change/description/
 * https://www.lintcode.com/problem/coin-change/description
 * 322. Coin Change
 */

import java.lang.Math;

public class Solution {

	public static void main(String[] args) {
		if (args.length != 1) {
			return;
		}
		int[] coins = new int[]{1, 2, 5};
		int amount = Integer.parseInt(args[0]);
		//System.out.println(coinChange1(coins, amount));
		System.out.println(coinChange2(coins, amount));
	}

	public static int coinChange1(int[] coins, int amount) {
		int res = recursive1(coins, amount);
		if (res == Integer.MAX_VALUE) {
			return -1;
		}
		return res;
	}

	public static int recursive1(int[] coins, int amount) {
		if (amount == 0) {
			return 0;
		}
		int res = Integer.MAX_VALUE;
		for (int coin: coins) {
			if (amount >= coin) {
				int sub = recursive1(coins, amount - coin);
				if (sub == Integer.MAX_VALUE) {
					continue;
				}
				res = Math.min(sub + 1, res);
			}
		}
		return res;
	}

	public static int coinChange2(int[] coins, int amount) {
		// 1.确定状态，最后一步和子问题：
		// 最后一枚硬币为k，那么子问题为amount-k最少用多少枚硬币来换
		// 2.转移方程：
		// min{f(x-coin1) + 1, f(x-coin2) + 1}
		// 3.初始条件和边界情况：
		// f(0) = 0
		// amount要大于等于硬币的值，以及dp往前追溯时，应该是有解的
		// 4.计算顺序
		// 从amount=1一直计算到等于amount

		int[] dp = new int[amount+1];
		dp[0] = 0;
		int len = coins.length;

		int i, j;
		for (i = 1; i < dp.length; i++) {
			dp[i] = Integer.MAX_VALUE;
			for (j = 0; j < len; j++) {
				if (i >= coins[j] && dp[i-coins[j]] != Integer.MAX_VALUE) {
				    dp[i] = Math.min(dp[i-coins[j]] + 1, dp[i]);
				}
			}
		}

		if (dp[amount] == Integer.MAX_VALUE) {
			dp[amount] = -1;
		}
		return dp[amount];
	}

}
