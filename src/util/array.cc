#include <time.h>
#include <assert.h>
#include <fmt/color.h>
#include "src/util/util.h"

template <typename T>
void print_array(T *arr, int n){
    for (int i = 0; i < n; i++) {
        fmt::print(fmt::emphasis::underline, "{} ", arr[i]);
    }
    fmt::print("\n");
}

template
void print_array(int*, int);


template <typename T>
T *generate_array(int n) {
    T *arr;
    arr = (T*) malloc(sizeof(T) * n);
    srand(time(NULL));
    for (int i = 0; i < n; i++) {
        arr[i] = rand() % 100;
    }
    return arr;
}

template 
int* generate_array(int);


template <typename T>
void assert_array_order(T *arr, int n, bool asc) {
    for (int i = 0; i < n - 1; i++) {
        if (asc) {
            assert(arr[i+1] >= arr[i]);
        } else {
            assert(arr[i+1] <= arr[i]);
        }
    }
}

template
void assert_array_order(int*, int, bool);

