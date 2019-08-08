// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

const CHAR_BYTES = 4
const HEX_BASE = 16
const BINARY_BASE = 2

function convertJoinEUI (joinEUI, fromBase, toBase) {
  return parseInt(joinEUI, fromBase).toString(toBase)
}

/**
 * Computes the join EUI prefix given `joinEUI` and its `length`.
 * @param {string} joinEUI - The join EUI.
 * @param {number} length - The length of the prefix.
 * @returns {string} - The join EUI prefix.
 */
function computePrefix (joinEUI, length = 0) {
  if (length % CHAR_BYTES !== 0) {
    const binaryPrefix = convertJoinEUI(joinEUI, HEX_BASE, BINARY_BASE).slice(0, length)

    return convertJoinEUI(binaryPrefix, BINARY_BASE, HEX_BASE)
  }

  return joinEUI.slice(0, length / CHAR_BYTES)
}

export default computePrefix
