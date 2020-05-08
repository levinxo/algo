/**
 * 122. Best Time to Buy and Sell Stock II
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
 */
class Solution {


    public static void main(String[] args) {
        int[] prices = new int[]{7,1,5,3,6,4};
        System.out.println("贪心算法：");
        System.out.println(maxProfit1(prices));
        System.out.println("波峰波谷法：");
        System.out.println(maxProfit2(prices));
        System.out.println("动态规划：");
        System.out.println(maxProfit3(prices));
    }


    public static int maxProfit1(int[] prices) {
        /**
         * 贪心算法：在每一步都求当前最优解
         * Time Complexity O(n) 循环一次prices数组
         * Space Complexity O(1) 使用了有限个数的常量
         */
        int profit = 0;
        for (int i = 1; i < prices.length; i++) {
            if (prices[i] - prices[i-1] > 0) {
                profit += prices[i] - prices[i-1];
            }
        }
        return profit;
    }


    public static int maxProfit2(int[] prices) {
        /**
         * 波峰波谷法，每个波峰-波谷就是一个盈利点
         * Time Complexity O(n) 循环一次prices数组
         * Space Complexity O(1) 使用了有限个数的常量
         */
        if (prices == null || prices.length == 0) {
            return 0;
        }

        int profit = 0;
        int max = prices[0], min = prices[0];
        int i = 0;
        while (i < prices.length - 1) {
            // 波谷判断
            while (i < prices.length - 1 && prices[i] >= prices[i+1]) {
                i++;
            }
            min = prices[i];
            // 波峰判断
            while (i < prices.length -1 && prices[i] <= prices[i+1]) {
                i++;
            }
            max = prices[i];
            profit += max - min;
        }
        return profit;
    }


    public static int maxProfit3(int[] prices) {
        /**
         * 和121题思路一样，只是状态方程有了变化
         * buy = max(buy, sell-prices[0])
         * sell = max(sell, buy+prices[0])
         *
         * 仔细观察可以发现，其实就是一个dp二维数组
         * dp[i][0/1]，代表第i天的最大收益，第二列0和1分别表示卖和买的状态
         */
        if (prices == null || prices.length <= 1) {
            return 0;
        }

        int buy = -prices[0];
        int sale = 0;

        for (int i = 1; i < prices.length; i++) {
            buy = Math.max(buy, sale-prices[i]);
            sale = Math.max(sale, buy+prices[i]);
        }
        return sale;
    }


}
