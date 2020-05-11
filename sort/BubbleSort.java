import java.util.Arrays;

class BubbleSort {


    public static void main(String[] args) {
        int[] arr = new int[]{5,2,18,3,11,90,45,22};
        arr = bubbleSort(arr);
        System.out.println(Arrays.toString(arr));
    }


    /**
     * Time Complexity O(n^2)
     * Space Complexity O(1)
     */
    public static int[] bubbleSort(int[] arr) {
        if (arr == null || arr.length <= 1) {
            return arr;
        }

        for (int i = 0; i < arr.length - 1; i++) {
            boolean change = false;
            for (int j = 0; j < arr.length - i - 1; j++) {
                if (arr[j] > arr[j+1]) {
                    int tmp = arr[j+1];
                    arr[j+1] = arr[j];
                    arr[j] = tmp;
                    change = true;
                }
            }
            if (!change) {
                break;
            }
        }
        return arr;
    }


}
