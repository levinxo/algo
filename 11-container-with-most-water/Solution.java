import java.lang.Math;


/**
 * 11. 盛最多水的容器
 * https://leetcode-cn.com/problems/container-with-most-water/
 */
class Solution {


    /**
     * Time Complexity O(n)
     * Space Complexity O(1)
     */
    public static int maxArea(int[] height) {
        int left = 0, right = height.length - 1;

        int maxFill = 0;

        while (left < right) {
            int fill = Math.min(height[left], height[right]) * (right - left);
            maxFill = Math.max(maxFill, fill);
            if (height[left] <= height[right]) {
                left++;
            } else {
                right--;
            }
        }
        return maxFill;
    }

    public static void main(String[] args) {
        //int[] rainDrop = new int[]{1,8,6,2,5,4,8,3,7};
        int[] rainDrop = new int[]{2,3,4,5,18,17,6};
        System.out.println(maxArea(rainDrop));
    }

}
