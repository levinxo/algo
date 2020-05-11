/**
 * 309. Best Time to Buy and Sell Stock with Cooldown
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
 */

import java.lang.Math;

class Solution {


    public static void main(String[] args) {
        int[] prices = new int[]{1,2,3,0,2};
        System.out.println(maxProfit1(prices));
        System.out.println(maxProfit2(prices));
    }

    public static int maxProfit1(int[] prices) {
        /**
         * 使用二维数组动态规划来解决
         * 1.最后一步和子问题
         * 参考122题，分为两种状态，买和卖
         * 因为有cooldown机制，在买入时，只有前两天之前卖出的才能买入
         * 2.状态转移方程
         * dp[i][0] = max{dp[i-1][0], dp[i-1][1]+prices[i]}
         * dp[i][1] = max{dp[i-1][1], dp[i-2][0]-prices[i]}
         * 3.初始值和边界条件
         * dp[0][0] = 0
         * dp[0][1] = -prices[0]
         * 注意i-2时不能小于0
         * 4.计算顺序
         * 从头计算即可
         */

        if (prices == null || prices.length <= 1) {
            return 0;
        }

        int length = prices.length;
        int[][] dp = new int[length][2];

        // initialization
        dp[0][0] = 0;
        dp[0][1] = -prices[0];

        for (int i = 1; i < length; i++) {
            dp[i][0] = Math.max(dp[i-1][0], dp[i-1][1] + prices[i]);
            dp[i][1] = Math.max(dp[i-1][1], (i >= 2 ? dp[i-2][0] : 0) - prices[i]);
        }
        return dp[length-1][0];
    }


    public static int maxProfit2(int[] prices) {
        /**
         * 针对二维数组dp做优化
         * Space Complexity O(1) 的动态规划写法
         */

        if  (prices == null || prices.length <= 1) {
            return 0;
        }

        int length = prices.length;

        // initialization
        int buy = -prices[0];
        int sell = 0;
        int prev_sell = 0;

        for (int i = 1; i < length; i++) {
            int tmp_sell = sell;
            sell = Math.max(sell, buy + prices[i]);
            buy = Math.max(buy, prev_sell - prices[i]);
            prev_sell = tmp_sell;
        }
        return sell;
    }


}
