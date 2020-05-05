/**
 * 121. Best Time to Buy and Sell Stock
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock
 */
class Solution {

	public static void main(String[] args) {
		int[] prices = new int[]{7,1,5,3,6,4};
		System.out.println("暴力法：");
		System.out.println(maxProfit1(prices));
		System.out.println("差值法：");
		System.out.println(maxProfit2(prices));
	}

	public static int maxProfit1(int[] prices) {
		/**
		 * 暴力法
		 * Time Complexity O(n^2) 循环运行n*(n-1)/2次
		 * Space Complexity O(1) 只使用了常数个变量
		 */
		int profit = 0;
		for (int i = 0; i < prices.length - 1; i++) {
			for (int j = i+1; j < prices.length; j++) {
				if (prices[j] - prices[i] > profit) {
					profit = prices[j] - prices[i];
				}
			}
		}
		return profit;
	}

	public static int maxProfit2(int[] prices) {
		/**
		 * 一次遍历
		 * Time Complexity O(n) 遍历一次prices数组
		 * Space Complexity O(1) 使用了常数个变量
		 */
		int price = Integer.MAX_VALUE;
		int profit = 0;

		for (int i = 0; i < prices.length; i++) {
			if (prices[i] < price) {
				price = prices[i];
			} else if (prices[i] - price > profit) {
				profit = prices[i] - price;
			}
		}

		return profit;
	}

}
