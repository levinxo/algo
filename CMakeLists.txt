cmake_minimum_required(VERSION 3.14)

project(algo)


set(CMAKE_BUILD_TYPE "Debug")
set(CMAKE_CXX_STANDARD 11)
#add_definitions(-std=c++11)
set(CMAKE_CXX_FLAGS_DEBUG "$ENV{CXXFLAGS} -O0 -Wall -g -ggdb")
set(CMAKE_CXX_FLAGS_RELEASE "$ENV{CXXFLAGS} -O3 -Wall")

include_directories(${PROJECT_SOURCE_DIR})

add_subdirectory(src/util)
add_subdirectory(src/sort)
add_subdirectory(src/bfs)

set(EXTRA_LIBS ${EXTRA_LIBS} sort bfs)

set(CMAKE_MODULE_PATH "${CMAKE_SOURCE_DIR}/cmake;${CMAKE_MODULE_PATH}")
include(thirdparty)

add_executable(algo src/main.cc)
target_link_libraries(algo ${EXTRA_LIBS})

# 启用测试
enable_testing()

# 测试程序是否成功运行
add_test(test_run algo)

# 指定安装路径
install(TARGETS algo DESTINATION bin)

