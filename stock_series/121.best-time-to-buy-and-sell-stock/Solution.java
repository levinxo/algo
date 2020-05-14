/**
 * 121. Best Time to Buy and Sell Stock
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock
 */
class Solution {

	public static void main(String[] args) {
		int[] prices = new int[]{7,1,5,3,6,4};
		System.out.println("暴力法：");
		System.out.println(maxProfit1(prices));
		System.out.println("贪心算法：");
		System.out.println(maxProfit2(prices));
		System.out.println("动态规划：");
		System.out.println(maxProfit3(prices));
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
		 * 贪心算法，一次遍历
         * 遍历到i时，我们能知道prices[:i]中的最小值，也能知道prices[i:]的最大收益
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

    public static int maxProfit3(int[] prices) {
        /**
         * 动态规划算法。
         * 1.最后一步和子问题分析
         * 每天结束时，我们只会有两种状态，一种是已经买了，一种是已经卖了
         * 买了的，可能是之前买的，也可能是今天买的
         * 卖了的，可能是之前卖的，也可能是今天卖的
         * 2.转移方程
         * buy = max(buy, -prices[i])
         * sell = max(sell, buy+prices[i])
         * 3.初始值和边界条件
         * buy = -prices[0]，第一天就持有股票
         * sell = 0，第一天买了就卖
         * 4.计算顺序，从头开始就好
		 * 
		 * 仔细观察可以发现，其实就是一个dp二维数组
		 * dp[i][0/1]，代表第i天的最大收益，第二列0和1分别表示卖和买的状态
         */

        if (prices == null || prices.length <= 1) {
            return 0;
        }

        // initialization
        int buy = -prices[0];
        int sell = 0;

        for (int i = 1; i < prices.length; i++) {
			sell = Math.max(sell, buy+prices[i]);
			buy = Math.max(buy, -prices[i]);
            System.out.println("buy: " + buy + " sell: " + sell);
            // buy: -1 sell: 0
			// buy: -1 sell: 4
			// buy: -1 sell: 4
			// buy: -1 sell: 5
			// buy: -1 sell: 5
        }
        return sell;
    }


}
