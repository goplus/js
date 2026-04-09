/*
Copyright 2026 The XGo Authors (xgo.dev)
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package builtin

import (
	"github.com/goplus/js/internal/universe"
)

// ----------------------------------------------------------------------------

// int32 is the set of all signed 32-bit integers.
// Range: -2147483648 through 2147483647.
type Int32 universe.Number

// Int32_Cast: func int32(v int) int32
func Int32_Cast__0(v universe.Number) Int32 {
	return Int32(v.XGo_Or(0))
}

// XGo_Or implements func (a int32) | (b int32) int32
func (a Int32) XGo_Or(b Int32) Int32 {
	return Int32(universe.Number(a).XGo_Or(universe.Number(b)))
}

// ----------------------------------------------------------------------------
