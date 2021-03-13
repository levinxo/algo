
aux_source_directory(. FILE_LIB_SRCS)
add_library(util ${FILE_LIB_SRCS})

target_link_libraries(util fmt)

# 指定 util 库的安装路径
install (TARGETS util DESTINATION lib)
install (FILES array.h DESTINATION include)

