/**
 * 714. Best Time to Buy and Sell Stock with Transaction Fee
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
 */

import java.lang.Math;

class Solution {


    public static void main(String[] args){
        int[] prices = new int[]{1, 3, 2, 8, 4, 9};
        int fee = 2;
        System.out.println(maxProfit(prices, fee));
    }


    public static int maxProfit(int[] prices, int fee) {
        /**
         * 1.最后一步和子问题分析
         * 股票依然是只有买和卖两种状态
         * 只是买入时需要交个手续费
         * 2.状态转移方程
         * dp[i][0] = max{dp[i-1][0], dp[i-1][1] + prices[i]}
         * dp[i][1] = max{dp[i-1][1], dp[i-1][0] - prices[i] - fee}
         * 3.初始化和边界条件
         * dp[0][0] = 0;
         * dp[0][1] = -prices[0] - fee;
         * 4.计算顺序
         * 从头计算即可
         *
         * 参考之前的股票问题，下面是二维数组简化为Space Complexity O(1)的写法
         */

        if (prices == null || prices.length <= 1) {
            return 0;
        }

        int length = prices.length;

        // initialization
        int buy = -prices[0] - fee;
        int sell = 0;

        for (int i = 1; i < length; i++) {
            int tmp_buy = buy;
            buy = Math.max(buy, sell - prices[i] - fee);
            sell = Math.max(sell, tmp_buy + prices[i]);
        }
        return sell;
    }


}
