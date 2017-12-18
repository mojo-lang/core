cmake_minimum_required(VERSION 3.7)
project(core)

set(CMAKE_CXX_STANDARD 14)

INCLUDE_DIRECTORIES(
        /usr/local/include
        ${PROJECT_SOURCE_DIR}/../../aimap/ncraft/include
)

set(SOURCE_FILES
        include/mojo/core/any.hpp
        include/mojo/core/array.hpp
        include/mojo/core/bool.hpp
        include/mojo/core/dictionary.hpp
        include/mojo/core/error.hpp
        include/mojo/core/null.hpp
        include/mojo/core/numeric.hpp
        include/mojo/core/object.hpp
        include/mojo/core/optional.hpp
        include/mojo/core/range.hpp
        include/mojo/core/result.hpp
        include/mojo/core/string.hpp
        include/mojo/core/tuple.hpp
        include/mojo/core/union.hpp
        src/mojo/mojo_core.cpp
        )

#add_executable(core ${SOURCE_FILES})