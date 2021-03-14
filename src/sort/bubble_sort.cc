#include <fmt/color.h>
#include "src/util/util.h"
#include "src/sort/sort.h"


namespace sort {


void bubble_sort(int *arr, int n) {
    if (arr == NULL) {
        return;
    }

    for (int i = 0; i < n - 1; i++) {
        bool change = false;
        for (int j = 0; j < n - 1 - i; j++) {
            if (arr[j] > arr[j+1]) {
                int tmp = arr[j];
                arr[j] = arr[j+1];
                arr[j+1] = tmp;
                change = true;
            }
        }
        if (!change) {
            break;
        }
    }
}


void test_bubble_sort() {
    fmt::print(fg(fmt::color::yellow), "Bubble Sort\n");

    int len = 10;
    test_sort(bubble_sort, len, len);
}


}   // namespace sort
