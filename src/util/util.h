
#ifndef UTIL_UTIL_H_
#define UTIL_UTIL_H_

template <typename T>
void print_array(T*, int);

template <typename T>
T* generate_array(int);

template <typename T>
void assert_array_order(T*, int, bool);

template <typename F, typename... Args>
void test_sort(F, int, Args&... rest);

#endif	// UTIL_UTIL_H_

