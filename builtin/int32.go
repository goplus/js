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
)

// ----------------------------------------------------------------------------

// int32 is the set of all signed 32-bit integers.
// Range: -2147483648 through 2147483647.
type Int32 primitive.Number

// Int32_Cast: func int32(v int) int32
func Int32_Cast__0(v primitive.Number) Int32 {
	return Int32(v.XGo_Or(0))
}

// XGo_Add implements func (a int32) + (b int32) int32
func (a Int32) XGo_Add(b Int32) Int32 {
	return Int32((primitive.Number(a) + primitive.Number(b)).XGo_Or(0))
}

// XGo_Sub implements func (a int32) - (b int32) int32
func (a Int32) XGo_Sub(b Int32) Int32 {
	return Int32((primitive.Number(a) - primitive.Number(b)).XGo_Or(0))
}

// XGo_Mul implements func (a int32) * (b int32) int32
func (a Int32) XGo_Mul(b Int32) Int32 {
	return Int32((primitive.Number(a) * primitive.Number(b)).XGo_Or(0))
}

// XGo_Quo implements func (a int32) / (b int32) int32
func (a Int32) XGo_Quo(b Int32) Int32 {
	return Int32((primitive.Number(a) / primitive.Number(b)).XGo_Or(0))
}

// XGo_Rem implements func (a int32) % (b int32) int32
func (a Int32) XGo_Rem(b Int32) Int32 {
	return Int32(primitive.Number(a).XGo_Rem(primitive.Number(b)))
}

// XGo_Or implements func (a int32) | (b int32) int32
func (a Int32) XGo_Or(b Int32) Int32 {
	return Int32(primitive.Number(a).XGo_Or(primitive.Number(b)))
}

// XGo_And implements func (a int32) & (b int32) int32
func (a Int32) XGo_And(b Int32) Int32 {
	return Int32(primitive.Number(a).XGo_And(primitive.Number(b)))
}

// XGo_Xor implements func (a int32) ^ (b int32) int32
func (a Int32) XGo_Xor(b Int32) Int32 {
	return Int32(primitive.Number(a).XGo_Xor(primitive.Number(b)))
}

// XGo_AndNot implements func (a int32) &^ (b int32) int32
func (a Int32) XGo_AndNot(b Int32) Int32 {
	return Int32(primitive.Number(a).XGo_And(primitive.Number(b).XGo_Not()))
}

// XGo_Lsh implements func (a int32) << (b uint) int32
func (a Int32) XGo_Lsh(b Uint) Int32 {
	return Int32(primitive.Number(a).XGo_Lsh(primitive.Number(b)))
}

// XGo_Rsh implements func (a int32) >> (b uint) int32
func (a Int32) XGo_Rsh(b Uint) Int32 {
	return Int32(primitive.Number(a).XGo_Rsh(primitive.Number(b)))
}

// XGo_Neg implements func -(a int32) int32
func (a Int32) XGo_Neg() Int32 {
	return Int32((-primitive.Number(a)).XGo_Or(0))
}

// XGo_Not implements func ^(a int32) int32
func (a Int32) XGo_Not() Int32 {
	return Int32(primitive.Number(a).XGo_Not())
}

// XGo_LT implements func (a int32) < (b int32) bool
func (a Int32) XGo_LT(b Int32) bool {
	return primitive.Number(a) < primitive.Number(b)
}

// XGo_LE implements func (a int32) <= (b int32) bool
func (a Int32) XGo_LE(b Int32) bool {
	return primitive.Number(a) <= primitive.Number(b)
}

// XGo_GT implements func (a int32) > (b int32) bool
func (a Int32) XGo_GT(b Int32) bool {
	return primitive.Number(a) > primitive.Number(b)
}

// XGo_GE implements func (a int32) >= (b int32) bool
func (a Int32) XGo_GE(b Int32) bool {
	return primitive.Number(a) >= primitive.Number(b)
}

// ----------------------------------------------------------------------------
