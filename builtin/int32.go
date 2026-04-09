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
	"github.com/goplus/js/primitive"
	"github.com/goplus/js/primitive/math"
)

// ----------------------------------------------------------------------------

// int32 is the set of all signed 32-bit integers.
// Range: -2147483648 through 2147483647.
type Int32 primitive.Number

// Int32_Cast: func int32(v float64) int32
func Int32_Cast__0(v primitive.Number) Int32 {
	return Int32(v.JS_Or(0))
}

// JS_Add implements func (a int32) + (b int32) int32
func (a Int32) XGo_Add(b Int32) Int32 {
	return Int32((primitive.Number(a) + primitive.Number(b)).JS_Or(0))
}

// JS_Sub implements func (a int32) - (b int32) int32
func (a Int32) XGo_Sub(b Int32) Int32 {
	return Int32((primitive.Number(a) - primitive.Number(b)).JS_Or(0))
}

// JS_Mul implements func (a int32) * (b int32) int32
func (a Int32) XGo_Mul(b Int32) Int32 {
	return Int32(math.Imul(primitive.Number(a), primitive.Number(b)))
}

// JS_Quo implements func (a int32) / (b int32) int32
func (a Int32) XGo_Quo(b Int32) Int32 {
	return Int32((primitive.Number(a) / primitive.Number(b)).JS_Or(0))
}

// JS_Rem implements func (a int32) % (b int32) int32
func (a Int32) XGo_Rem(b Int32) Int32 {
	return Int32(primitive.Number(a).JS_Rem(primitive.Number(b)))
}

// JS_Or implements func (a int32) | (b int32) int32
func (a Int32) XGo_Or(b Int32) Int32 {
	return Int32(primitive.Number(a).JS_Or(primitive.Number(b)))
}

// JS_And implements func (a int32) & (b int32) int32
func (a Int32) XGo_And(b Int32) Int32 {
	return Int32(primitive.Number(a).JS_And(primitive.Number(b)))
}

// JS_Xor implements func (a int32) ^ (b int32) int32
func (a Int32) XGo_Xor(b Int32) Int32 {
	return Int32(primitive.Number(a).JS_Xor(primitive.Number(b)))
}

// JS_AndNot implements func (a int32) &^ (b int32) int32
func (a Int32) XGo_AndNot(b Int32) Int32 {
	return Int32(primitive.Number(a).JS_And(primitive.Number(b).JS_Not()))
}

// JS_Lsh implements func (a int32) << (b uint) int32
func (a Int32) XGo_Lsh(b Uint) Int32 {
	return Int32(primitive.Number(a).JS_Lsh(primitive.Number(b)))
}

// JS_Rsh implements func (a int32) >> (b uint) int32
func (a Int32) XGo_Rsh(b Uint) Int32 {
	return Int32(primitive.Number(a).JS_Rsh(primitive.Number(b)))
}

// JS_Neg implements func -(a int32) int32
func (a Int32) XGo_Neg() Int32 {
	return Int32((-primitive.Number(a)).JS_Or(0))
}

// JS_Not implements func ^(a int32) int32
func (a Int32) XGo_Not() Int32 {
	return Int32(primitive.Number(a).JS_Not())
}

// JS_LT implements func (a int32) < (b int32) bool
func (a Int32) XGo_LT(b Int32) bool {
	return primitive.Number(a) < primitive.Number(b)
}

// JS_LE implements func (a int32) <= (b int32) bool
func (a Int32) XGo_LE(b Int32) bool {
	return primitive.Number(a) <= primitive.Number(b)
}

// JS_GT implements func (a int32) > (b int32) bool
func (a Int32) XGo_GT(b Int32) bool {
	return primitive.Number(a) > primitive.Number(b)
}

// JS_GE implements func (a int32) >= (b int32) bool
func (a Int32) XGo_GE(b Int32) bool {
	return primitive.Number(a) >= primitive.Number(b)
}

// ----------------------------------------------------------------------------
