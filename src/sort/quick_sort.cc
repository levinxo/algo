#include <fmt/color.h>
#include "src/util/util.h"
#include "src/sort/sort.h"


namespace sort {


void quick_sort(int *arr, int begin, int end) {
    if (arr == NULL || begin >= end) {
        return;
    }

    int i = begin, j = end;
    int pivot = arr[begin + (end - begin) / 2];

    while (i <= j) {
        while (arr[i] < pivot) {
            i++;
        }
        while (arr[j] > pivot) {
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
    quick_sort(arr, begin, j);
    quick_sort(arr, i, end);
}


void test_quick_sort() {
    fmt::print(fg(fmt::color::yellow), "Quick Sort\n");

    int len = 10;
    int start = 0;
    int end = len - 1;
    test_sort(quick_sort, len, start, end);
}


}      // namespace sort

