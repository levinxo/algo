include(FetchContent)

FetchContent_Declare(
  fmt 
  GIT_REPOSITORY    https://github.com/fmtlib/fmt.git
  GIT_TAG           7bdf0628b1276379886c7f6dda2cef2b3b374f0b   #v7.1.3
)

FetchContent_MakeAvailable(fmt)

