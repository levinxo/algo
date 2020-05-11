import java.util.Arrays;

class QuickSort {


    public static void main(String[] args) {
        int[] arr = new int[]{5,2,18,3,11,90,45,22};
        quickSort(arr, 0, arr.length - 1);
        System.out.println(Arrays.toString(arr));
    }


    public static void quickSort(int[] arr, int start, int end) {
        if (arr == null || arr.length < 2 || start >= end) {
            return;
        }

        int i = start, j = end;
        int pivot = arr[start + (end - start) / 2];

        while (i <= j) {
            while (arr[i] < pivot) {
                i++;
            }
            while (arr[j] > pivot) {  // 使用>=时，会产生java.lang.StackOverflowError
                j--;
            }
            if (i < j) {
                int tmp = arr[j];
                arr[j] = arr[i];
                arr[i] = tmp;
                i++;
                j--;
            } else if (i == j) {
                i++;
            }
        }
        quickSort(arr, start, j);
        quickSort(arr, i, end);
    }


}
