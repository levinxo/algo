import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

/**
 * https://leetcode-cn.com/problems/pascals-triangle/
 * 118. 杨辉三角
 */

class Solution {


    /**
     * 动态规划：基于前一行来构造下一行
     * Time Complexity O(n^2)
     * Space Complexity O(n^2)
     */
    public List<List<Integer>> generate(int numRows) {
        List<Integer> a = new ArrayList<>();
        List<Integer> b = Arrays.asList(1);
        List<List<Integer>> list = new ArrayList<>();

        int n = 0;

        while (n < numRows) {
            list.add(b);
            a = b;
            b = new ArrayList<>();
            for (int i = 0; i < a.size()+1; i++) {
                if (i == 0 || i == a.size()) {
                    b.add(1);
                } else {
                    b.add(a.get(i) + a.get(i-1));
                }
            }
            n++;
        }
        return list;
    }


}
