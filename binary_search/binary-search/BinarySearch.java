public class BinarySearch {


    public static int binarySearch(int[] arr, int value) {
        if (arr == null || arr.length == 0) {
            return -1;
        }

        int low = 0, high = arr.length - 1;

        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (value < arr[mid]) {
                high = mid - 1;
            } else if (value > arr[mid]) {
                low = mid + 1;
            } else {
                return mid;
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
