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
        System.out.println(maxProfit(prices));
    }

    public static int maxProfit(int[] prices) {
        /**
         * 1.最后一步和子问题分析
         * 通过dp[i]代表某天的最大收益，但找不到dp[i-1]和dp[i]的关系。
         * 那么尝试使用状态来表示。
         * dp[i][j][0]表示第i天已进行了j次交易，0表示当天最终不持有股票。
         * dp[i][j][1]表示第i天已进行了j次交易，1表示当天最终持有股票。
         * 其中交易次数在买入的时候就加一次
         *
         * 2.状态转移方程
         * 当天最终不持有股票，最大收益情况分析：
         * 无论卖出与否，j都不会变，因为卖出不算交易次数
         * a.前一天不持有股票 b.前一天持有股票，今天卖了，因此有利润prices[i]
         * dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
         *
         * 当天最终持有股票，最大收益情况分析：
         * a.前一天有股票 b.前一天没有股票，今天买了，因此要消耗prices[i]来买入
         * dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
         *
         * 3.初始化和边界条件
         * dp[0][j][0] = 0
         * dp[0][j][1] = -prices[i]
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
            for (int j = 1; j <= K; j++) {
                if (i == 0) {
                    dp[i][j][0] = 0;
                    dp[i][j][1] = -prices[i];
                    System.out.println("dp["+i+"]["+j+"][0]: " + dp[i][j][0] + " dp["+i+"]["+j+"][1]: " + dp[i][j][1]);
                    continue;           // initialization后跳过
                }
                dp[i][j][0] = Math.max(dp[i-1][j][0], dp[i-1][j][1] + prices[i]);
                dp[i][j][1] = Math.max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i]);
                System.out.println("dp["+i+"]["+j+"][0]: " + dp[i][j][0] + " dp["+i+"]["+j+"][1]: " + dp[i][j][1]);
            }
        }

        return dp[length-1][K][0];
    }


}
