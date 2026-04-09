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

// Float32_Cast: func float32(v float64) float32
func Float32_Cast__0(v primitive.Number) Float32 {
	return Float32(math.Fround(v))
}

// XGo_Add implements func (a float32) + (b float32) float32
func (a Float32) XGo_Add(b Float32) Float32 {
	return Float32(math.Fround(primitive.Number(a) + primitive.Number(b)))
}

// XGo_Sub implements func (a float32) - (b float32) float32
func (a Float32) XGo_Sub(b Float32) Float32 {
	return Float32(math.Fround(primitive.Number(a) - primitive.Number(b)))
}

// XGo_Mul implements func (a float32) * (b float32) float32
func (a Float32) XGo_Mul(b Float32) Float32 {
	return Float32(math.Fround(primitive.Number(a) * primitive.Number(b)))
}

// XGo_Quo implements func (a float32) / (b float32) float32
func (a Float32) XGo_Quo(b Float32) Float32 {
	return Float32(math.Fround(primitive.Number(a) / primitive.Number(b)))
}

// XGo_Rem implements func (a float32) % (b float32) float32
func (a Float32) XGo_Rem(b Float32) Float32 {
	return Float32(math.Fround(primitive.Number(a).JS_Rem(primitive.Number(b))))
}

// XGo_Neg implements func -(a float32) float32
func (a Float32) XGo_Neg() Float32 {
	return Float32(-primitive.Number(a))
}

// XGo_LT implements func (a float32) < (b float32) bool
func (a Float32) XGo_LT(b Float32) bool {
	return primitive.Number(a) < primitive.Number(b)
}

// XGo_LE implements func (a float32) <= (b float32) bool
func (a Float32) XGo_LE(b Float32) bool {
	return primitive.Number(a) <= primitive.Number(b)
}

// XGo_GT implements func (a float32) > (b float32) bool
func (a Float32) XGo_GT(b Float32) bool {
	return primitive.Number(a) > primitive.Number(b)
}

// XGo_GE implements func (a float32) >= (b float32) bool
func (a Float32) XGo_GE(b Float32) bool {
	return primitive.Number(a) >= primitive.Number(b)
}

// ----------------------------------------------------------------------------
