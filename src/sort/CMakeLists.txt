
aux_source_directory(. FILE_LIB_SRCS)
add_library(sort ${FILE_LIB_SRCS})

target_link_libraries(sort util fmt)

# 指定安装路径
install (TARGETS sort DESTINATION lib)
install (FILES sort.h DESTINATION include)

