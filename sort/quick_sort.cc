/**
 * quick sort
 * 快速排序
 */

#include <iostream>
#include <vector>
#include <time.h>
#include <cstdlib>

void quickSort(std::vector<int> &arr, int start, int end) {
    if (arr.size() <= 1 || start >= end) {
        return;
    }

    int pivot = arr[start + (end - start) / 2];     // 这里要注意应该是数组的值
    int i = start, j = end;

    while (i <= j) {
        while (arr[i] < pivot) {
            i++;
        }
        while (arr[j] > pivot) {
            j--;
        }
        if (i < j) {
            int tmp = arr[i];
            arr[i] = arr[j];
            arr[j] = tmp;
            i++;
            j--;
        } else if (i == j) {      // 这里要注意判断i==j
            i++;
        }
    }
    quickSort(arr, start, j);
    quickSort(arr, i, end);
}

void print(std::vector<int> &arr){
    for (int i = 0; i < arr.size(); i++) {
        std::cout << arr[i] << " ";
    }
    std::cout << std::endl;
}

int main() {
    std::vector<int> arr;
    srand((unsigned) time(NULL));
    for (int i = 0; i < 10; i++) {
        arr.push_back(rand() % 100);
    }
    print(arr);
    quickSort(arr, 0, arr.size() - 1);
    print(arr);

    return 0;
}

