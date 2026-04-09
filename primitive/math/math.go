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
	E   = 2.718281828459045 // https://oeis.org/A001113
	Pi  = 3.141592653589793 // https://oeis.org/A000796
	Phi = 1.618033988749895 // https://oeis.org/A001622

	Sqrt2   = 1.4142135623730951 // https://oeis.org/A002193
	SqrtE   = 1.6487212707001282 // https://oeis.org/A019774
	SqrtPi  = 1.7724538509055159 // https://oeis.org/A002161
	SqrtPhi = 1.272019649514069  // https://oeis.org/A139339

	Ln2    = 0.6931471805599453 // https://oeis.org/A002162
	Log2E  = 1 / Ln2
	Ln10   = 2.302585092994046
	Log10E = 1 / Ln10
)

// Floating-point limit values.
const (
	MaxFloat32             = 3.4028234663852886e+38  // 0x1p127 * (1 + (1 - 0x1p-23))
	SmallestNonzeroFloat32 = 1.401298464324817e-45   // 1 / 2**(127 - 1 + 23)
	MaxFloat64             = 1.7976931348623157e+308 // 0x1p1023 * (1 + (1 - 0x1p-52))
	SmallestNonzeroFloat64 = 5e-324                  // 1 / 2**(1023 - 1 + 52)
)

// Integer limit values.
const (
	MaxInt    = 1<<63 - 1
	MinInt    = -1 << 63
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
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
func Copysign(x, y primitive.Number) primitive.Number
func Cos(primitive.Number) primitive.Number
func Cosh(primitive.Number) primitive.Number
func Dim(x, y primitive.Number) primitive.Number
func Exp(primitive.Number) primitive.Number
func Exp2(primitive.Number) primitive.Number
func Expm1(primitive.Number) primitive.Number
func Floor(primitive.Number) primitive.Number
func Hypot(p, q primitive.Number) primitive.Number
func Imul(primitive.Number, primitive.Number) primitive.Number
func Inf(sign int) primitive.Number
func IsInf(f primitive.Number, sign int) bool
func IsNaN(f primitive.Number) bool
func Log(primitive.Number) primitive.Number
func Log10(primitive.Number) primitive.Number
func Log1p(primitive.Number) primitive.Number
func Log2(primitive.Number) primitive.Number
func Max(x, y primitive.Number) primitive.Number
func Min(x, y primitive.Number) primitive.Number
func Mod(x, y primitive.Number) primitive.Number
func NaN() primitive.Number
func Pow(x, y primitive.Number) primitive.Number
func Pow10(n int) primitive.Number
func Round(primitive.Number) primitive.Number
func Sin(primitive.Number) primitive.Number
func Sinh(primitive.Number) primitive.Number
func Signbit(x primitive.Number) bool
func Sqrt(primitive.Number) primitive.Number
func Tan(primitive.Number) primitive.Number
func Tanh(primitive.Number) primitive.Number
func Trunc(primitive.Number) primitive.Number

// ----------------------------------------------------------------------------
