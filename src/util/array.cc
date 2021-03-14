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

template void print_array(int*, int);


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

template int* generate_array(int);


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

template void assert_array_order(int*, int, bool);


// notice: `rest` only recv lvalue
template <typename F, typename... Args>
void test_sort(F sort_func, int len, Args&... rest){
    int *arr = generate_array<int>(len);

    fmt::print(fg(fmt::color::light_yellow), "origin array:\n");
    print_array<int>(arr, len);

    sort_func(arr, rest...);

    fmt::print("resorted array:\n");
    print_array<int>(arr, len);

    assert_array_order<int>(arr, len, true);
}

// quick sort
typedef void (*f1)(int*, int, int);
template void test_sort(f1, int, int&, int&);


