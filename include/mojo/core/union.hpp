#ifndef MOJO_CORE_UNION_HPP
#define MOJO_CORE_UNION_HPP

#include <ncraft/core/either.hpp>

namespace mojo {
namespace core {

template <typename... Ts>
using Union = ncraft::core::Either<Ts...>;

}

template <typename... Ts>
using Union = core::Union<Ts...>;

}

#endif //MOJO_CORE_UNION_HPP
