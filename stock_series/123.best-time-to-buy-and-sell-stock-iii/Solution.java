/**
 * 123. Best Time to Buy and Sell Stock III
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/
 * https://www.lintcode.com/problem/best-time-to-buy-and-sell-stock-iii/
 */
import java.lang.Math;


public class Solution {

    public static void main(String[] args) {
        int[] prices = new int[]{3,3,5,0,0,3,1,4};
        //int[] prices = new int[]{1,2,3,4,5};
        //int[] prices = new int[]{7,6,4,3,1};
        System.out.println(maxProfit1(prices));
        System.out.println(maxProfit2(prices));
    }

    public static int maxProfit1(int[] prices) {
        /**
         * 1.最后一步和子问题分析
         * dp[i][j][0]表示第i天已进行了j次交易，0表示当天最终不持有股票。
         * dp[i][j][1]表示第i天已进行了j次交易，1表示当天最终持有股票。
         * 其中交易次数在买入的时候就加一次
         * 卖出不算交易次数
         *
         * 2.状态转移方程
         * 当天最终不持有股票，最大收益情况分析：
         * a.前一天不持有股票 b.前一天持有股票，今天卖了，因此有利润prices[i]
         * dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
         *
         * 当天最终持有股票，最大收益情况分析：
         * a.前一天有股票 b.前一天没有股票，今天买了，因此要消耗prices[i]来买入
         * dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
         *
         * 3.初始化和边界条件
         * dp[0][j][0] = 0
         * dp[0][j][1] = -prices[0]
         *
         * 4.遍历顺序，从头开始
         */

        if (prices == null || prices.length == 0) {
            return 0;
        }

        int length = prices.length;
        int K = 2;
        int[][][] dp = new int[length][K + 1][2];   // K+1

        for (int i = 0; i < length; i++) {
            //dp[i][0][0] = 0;   因为java初始化数组时初始值都为0，因此可省略
            for (int j = 1; j <= K; j++) {
                if (i == 0) {
                    dp[i][j][0] = 0;
                    dp[i][j][1] = -prices[i];
                    continue;           // initialization后跳过
                }
                dp[i][j][0] = Math.max(dp[i-1][j][0], dp[i-1][j][1] + prices[i]);
                dp[i][j][1] = Math.max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i]);
            }
        }

        return dp[length-1][K][0];
    }


    public static int maxProfit2(int[] prices) {
        /**
         * 更简单易懂的动态规划
         * 参考121题的思路
         * 1.最后一步和子问题
         * 第二次买时，必须是第一次卖了之后
         * 2.状态转移方程
         * buy1 = max(buy1, -prices[i])
         * sell1 = max(sell1, buy1+prices[i])
         * buy2 = max(buy2, sell1-prices[i])
         * sell2 = max(sell2, buy2+prices[i])
         * 3.初始值和边界条件
         * buy1, buy2 = -prices[0]
         * sell1, sell2 = 0
         * 4.计算顺序
         * 从头开始即可
         */

        if (prices == null || prices.length <= 1) {
            return 0;
        }

        // initialization
        int buy1 = -prices[0], buy2 = -prices[0], sell1 = 0, sell2 = 0;

        for (int i = 1; i < prices.length; i++) {
            buy1 = Math.max(buy1, -prices[i]);
            sell1 = Math.max(sell1, buy1+prices[i]);
            buy2 = Math.max(buy2, sell1-prices[i]);
            sell2 = Math.max(sell2, buy2+prices[i]);
        }
        return sell2;
    }


}
