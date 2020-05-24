/**
 * 744. Find Smallest Letter Greater Than Target
 * https://leetcode.com/problems/find-smallest-letter-greater-than-target/
 */

public class Solution {


    public static char nextGreatestLetter(char[] letters, char target) {
        int l = 0, r = letters.length - 1;
        while (l <= r) {
            int m = l + (r - l) / 2;
            if (letters[m] <= target) { //即使在letters中找到了target，也继续往后+，因为需要找比target大一点的
                l = m + 1;
            } else {
                r = m - 1;
            }
        }
        // 若target在letters的左侧外围，则l=0，应该返回index=0
        // 若target在letters的右侧外围，则l=n，根据循环原则，应该返回index=0
        return l < letters.length? letters[l]: letters[0];
    }


    public static void main(String[] args) {
        char[] letters = new char[]{'c', 'f', 'j'};
        char target = 'g';
        System.out.println(nextGreatestLetter(letters, target));
    }


}
