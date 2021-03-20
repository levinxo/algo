
aux_source_directory(. FILE_LIB_SRCS)
add_library(bfs ${FILE_LIB_SRCS})

target_link_libraries(bfs fmt)

