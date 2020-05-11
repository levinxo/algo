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
         * 1.最后一步和子问题
         * 第j次买时，必须是j-1次卖了之后
         * 2.状态转移方程
         * buy[j] = max{buy[j], sell[j-1]-prices[i]}
         * sell[j] = max{sell[j], buy[j]+prices[i]}
         * 3.初始值和边界条件
         * 所有的buy和sell，初始状态都分别为-prices[0]和0
         * 4.计算顺序
         * 从头计算即可
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
                int tmp_buy = buy[j];
                buy[j] = Math.max(buy[j], (j >= 1 ? sell[j-1]: 0) - prices[i]);
                sell[j] = Math.max(sell[j], tmp_buy + prices[i]);
            }
        }
        return sell[k-1];
    }

    public static int greedy(int[] prices) {
        /**
         * 无限次的贪心算法
         * Time Complexity O(n) 遍历数组n次
         * Space Complexity O(1) 有限个变量
         */
        int profit = 0;
        for (int i = 0; i < prices.length - 1; i++) {
            if (prices[i+1] > prices[i]) {
                profit += prices[i+1] - prices[i];
            }
        }
        return profit;
    }


}
