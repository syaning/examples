cmake_minimum_required(VERSION 3.18)
project(velox_demo LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_FIND_LIBRARY_PREFERENCE "STATIC")
set(CMAKE_FIND_LIBRARY_SUFFIXES ".a";".so")

find_path(VELOX_INCLUDE_DIR velox
          HINTS /usr/local/include)
find_library(VELOX_LIBRARY velox
             HINTS /usr/local/lib/velox
                   /usr/local/lib)

include(FindPackageHandleStandardArgs)
find_package_handle_standard_args(Velox REQUIRED
    VELOX_INCLUDE_DIR VELOX_LIBRARY)

if(VELOX_FOUND)
    add_library(velox UNKNOWN IMPORTED)
    set_target_properties(velox PROPERTIES
        IMPORTED_LOCATION ${VELOX_LIBRARY}
        INTERFACE_INCLUDE_DIRECTORIES ${VELOX_INCLUDE_DIR}
    )
endif()

add_executable(demo demo.cpp)
target_link_libraries(demo PRIVATE
    velox
    folly
    boost_context
    double-conversion
    event
    fmt
    gflags
    glog
    re2
    unwind
)