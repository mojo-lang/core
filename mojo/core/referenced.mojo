// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/// Used in the recyle referenced stituation
///
/// ```
/// type Node {
///    name: String @1
///    parent: Refenched<Node> @2
///    children: [Node] @3
/// }
/// ```
/// to json
/// ```
/// {
///    "name": "node-0-0", 
///    "parent": "node-0",
///    "children": [{"name": "node-0-0-0"}]
/// }
/// ```
type Referenced<T> {
    key: String @1
    val: T @2 @ignored
}