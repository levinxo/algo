/**
 * 42.trapping rain water
 * https://leetcode.com/problems/trapping-rain-water/
 *
 * 若用暴力解法，就是计算每个index的左侧最高和右侧最高，
 * 选最低的bar减去index的值得到可以盛的水。
 *
 * 用动态规划解法可以用两个数组记录每个index左侧最高和右侧最高，
 * 再遍历一次index计算出来可以盛的水。
 */
import java.lang.Math;

public class Solution {


    public static void main(String[] args) {
        int[] arr = new int[]{0,1,0,2,1,0,1,3,2,1,2,1};
        //int[] arr = new int[]{2,0,2};
        //int[] arr = new int[]{5,2,1,2,1,5};
        System.out.println(trap(arr));
    }


    /**
     * Time Complexity O(n)
     * Space Complexity O(n)
     */
    public static int trap(int[] arr) {
        if (arr == null || arr.length < 3) {
            return 0;
        }

        int length = arr.length;
        int ans = 0;

        int[] left = new int[length];
        int[] right = new int[length];

        left[0] = arr[0];
        right[length - 1] = arr[length - 1];
        // 为每一个index记录左边最高的
        for (int i = 1; i < length; i++) {
            left[i] = Math.max(arr[i], left[i-1]);
        }
        // 为每一个index记录右边最高的
        for (int i = length - 2; i >= 0; i--) {
            right[i] = Math.max(arr[i], right[i+1]);
        }

        // 计算最矮的bar能保证在当前index盛多少水
        for (int i = 1; i < length - 1; i++) {
            ans += Math.min(left[i], right[i]) - arr[i];
        }
        return ans;
    }


}
