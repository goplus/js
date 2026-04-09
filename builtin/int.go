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

// int is a signed integer type that is at least 32 bits in size. It is a
// distinct type, however, and not an alias for, say, int32.
//
// The size of int is 32 bits for JS.
type Int primitive.Number

// Int_Cast: func int(v float64) int
func Int_Cast__0(v primitive.Number) Int {
	return Int(v.JS_Or(0))
}

// JS_Add implements func (a int) + (b int) int
func (a Int) XGo_Add(b Int) Int {
	return Int((primitive.Number(a) + primitive.Number(b)).JS_Or(0))
}

// JS_Sub implements func (a int) - (b int) int
func (a Int) XGo_Sub(b Int) Int {
	return Int((primitive.Number(a) - primitive.Number(b)).JS_Or(0))
}

// JS_Mul implements func (a int) * (b int) int
func (a Int) XGo_Mul(b Int) Int {
	return Int(math.Imul(primitive.Number(a), primitive.Number(b)))
}

// JS_Quo implements func (a int) / (b int) int
func (a Int) XGo_Quo(b Int) Int {
	return Int((primitive.Number(a) / primitive.Number(b)).JS_Or(0))
}

// JS_Rem implements func (a int) % (b int) int
func (a Int) XGo_Rem(b Int) Int {
	return Int(primitive.Number(a).JS_Rem(primitive.Number(b)))
}

// JS_Or implements func (a int) | (b int) int
func (a Int) XGo_Or(b Int) Int {
	return Int(primitive.Number(a).JS_Or(primitive.Number(b)))
}

// JS_And implements func (a int) & (b int) int
func (a Int) XGo_And(b Int) Int {
	return Int(primitive.Number(a).JS_And(primitive.Number(b)))
}

// JS_Xor implements func (a int) ^ (b int) int
func (a Int) XGo_Xor(b Int) Int {
	return Int(primitive.Number(a).JS_Xor(primitive.Number(b)))
}

// JS_AndNot implements func (a int) &^ (b int) int
func (a Int) XGo_AndNot(b Int) Int {
	return Int(primitive.Number(a).JS_And(primitive.Number(b).JS_Not()))
}

// JS_Lsh implements func (a int) << (b uint) int
func (a Int) XGo_Lsh(b Uint) Int {
	return Int(primitive.Number(a).JS_Lsh(primitive.Number(b)))
}

// JS_Rsh implements func (a int) >> (b uint) int
func (a Int) XGo_Rsh(b Uint) Int {
	return Int(primitive.Number(a).JS_Rsh(primitive.Number(b)))
}

// JS_Neg implements func -(a int) int
func (a Int) XGo_Neg() Int {
	return Int((-primitive.Number(a)).JS_Or(0))
}

// JS_Not implements func ^(a int) int
func (a Int) XGo_Not() Int {
	return Int(primitive.Number(a).JS_Not())
}

// JS_LT implements func (a int) < (b int) bool
func (a Int) XGo_LT(b Int) bool {
	return primitive.Number(a) < primitive.Number(b)
}

// JS_LE implements func (a int) <= (b int) bool
func (a Int) XGo_LE(b Int) bool {
	return primitive.Number(a) <= primitive.Number(b)
}

// JS_GT implements func (a int) > (b int) bool
func (a Int) XGo_GT(b Int) bool {
	return primitive.Number(a) > primitive.Number(b)
}

// JS_GE implements func (a int) >= (b int) bool
func (a Int) XGo_GE(b Int) bool {
	return primitive.Number(a) >= primitive.Number(b)
}

// ----------------------------------------------------------------------------
