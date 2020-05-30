public class BinarySearch {


    public static int binarySearch(int[] arr, int value) {
        if (arr == null || arr.length == 0) {
            return -1;
        }

        // 初始化头尾index
        int l = 0, r = arr.length - 1;

        while (l <= r) {
            // 计算中间值
            int m = l + (r - l) / 2;    // 每次循环要重新计算m

            // 更新头尾index
            if (arr[m] == value) {
                return m;
            } else if (arr[m] > value) {
                r = m - 1;
            } else {
                l = m + 1;
            }
        }
        return -1;
    }


    public static void main(String[] args) {
        int[] arr = new int[]{1,3,5,7,9,11,13,15,17};
        int value = 3;
        int index = binarySearch(arr, value);
        System.out.println(index);
    }


}
