#include <fmt/color.h>
#include "src/util/util.h"


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
    int *arr = generate_array<int>(len);

    fmt::print(fg(fmt::color::light_yellow), "origin array:\n");
    print_array<int>(arr, len);

    quick_sort(arr, 0, len - 1);

    fmt::print("resorted array:\n");
    print_array<int>(arr, len);

    assert_array_order<int>(arr, len, true);
}


}      // namespace sort

