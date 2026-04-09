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

package math

import (
	"github.com/goplus/js/primitive"
)

const (
	XGoJSPackage = "Math"
)

// ----------------------------------------------------------------------------

// Mathematical constants.
const (
	E       = 2.718281828459045 // https://oeis.org/A001113
	LN10    = 2.302585092994046
	LN2     = 0.6931471805599453 // https://oeis.org/A002162
	LOG10E  = 1 / LN10
	LOG2E   = 1 / LN2
	PI      = 3.141592653589793  // https://oeis.org/A000796
	SQRT1_2 = 0.7071067811865476 // https://oeis.org/A002162
	SQRT_2  = 1.4142135623730951 // https://oeis.org/A002193
)

// ----------------------------------------------------------------------------

func Abs(primitive.Number) primitive.Number
func Acos(primitive.Number) primitive.Number
func Acosh(primitive.Number) primitive.Number
func Asin(primitive.Number) primitive.Number
func Asinh(primitive.Number) primitive.Number
func Atan(primitive.Number) primitive.Number
func Atan2(y, x primitive.Number) primitive.Number
func Atanh(primitive.Number) primitive.Number
func Cbrt(primitive.Number) primitive.Number
func Ceil(primitive.Number) primitive.Number
func Clz32(primitive.Number) primitive.Number
func Cos(primitive.Number) primitive.Number
func Cosh(primitive.Number) primitive.Number
func Exp(primitive.Number) primitive.Number
func Expm1(primitive.Number) primitive.Number
func Floor(primitive.Number) primitive.Number
func F16round(primitive.Number) primitive.Number
func Fround(primitive.Number) primitive.Number
func Hypot(p, q primitive.Number) primitive.Number
func Imul(primitive.Number, primitive.Number) primitive.Number
func Log(primitive.Number) primitive.Number
func Log10(primitive.Number) primitive.Number
func Log1p(primitive.Number) primitive.Number
func Log2(primitive.Number) primitive.Number
func Max(...primitive.Number) primitive.Number
func Min(...primitive.Number) primitive.Number
func Pow(x, y primitive.Number) primitive.Number
func Random() primitive.Number
func Round(primitive.Number) primitive.Number
func Sign(x primitive.Number) primitive.Number
func Sin(primitive.Number) primitive.Number
func Sinh(primitive.Number) primitive.Number
func Sqrt(primitive.Number) primitive.Number
func SumPrecise([]primitive.Number) primitive.Number
func Tan(primitive.Number) primitive.Number
func Tanh(primitive.Number) primitive.Number
func Trunc(primitive.Number) primitive.Number

// ----------------------------------------------------------------------------
