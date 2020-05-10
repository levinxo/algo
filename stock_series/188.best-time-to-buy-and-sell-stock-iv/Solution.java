/**
 * 188. Best Time to Buy and Sell Stock IV
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
 */
import java.util.Arrays;
import java.lang.Math;

class Solution {

    public static void main(String[] args) {
        int k = 2;
        int[] prices = new int[]{3,2,6,5,0,3};

        System.out.println(maxProfit(k, prices));
    }

    public static int maxProfit(int k, int[] prices) {
        /**
         * 参考123题k=2的情况，当2*k>length时使用贪心算法
         */
        if (prices == null || prices.length <= 1 || k < 1) {
            return 0;
        }

        int length = prices.length;
        if (2 * k > length) {
            return greedy(prices);
        }

        int[] buy = new int[k];
        int[] sell = new int[k];

        // initialization
        Arrays.fill(buy, -prices[0]);

        for (int i = 1; i < length; i++) {
            for (int j = 0; j < k; j++) {
                buy[j] = Math.max(buy[j], (j >= 1 ? sell[j-1]: 0) - prices[i]);
                sell[j] = Math.max(sell[j], buy[j] + prices[i]);
            }
        }
        return sell[k-1];
    }

    public static int greedy(int[] prices) {
        int profit = 0;
        for (int i = 0; i < prices.length - 1; i++) {
            if (prices[i+1] > prices[i]) {
                profit += prices[i+1] - prices[i];
            }
        }
        return profit;
    }


}
