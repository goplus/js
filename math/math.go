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
	"github.com/goplus/js/primitive/math"
)

const (
	E   = math.E
	Pi  = math.PI
	Phi = 1.61803398874989484820458683436563811772030917980576286213544862

	Sqrt2   = math.SQRT_2
	SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 // https://oeis.org/A019774
	SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 // https://oeis.org/A002161
	SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 // https://oeis.org/A139339

	Ln2    = math.LN2
	Log2E  = math.LOG2E
	Ln10   = math.LN10
	Log10E = math.LOG10E
)

// ----------------------------------------------------------------------------

//go:linkname Acos github.com/goplus/js/primitive/math.Acos
func Acos(x float64) float64

//go:linkname Acosh github.com/goplus/js/primitive/math.Acosh
func Acosh(x float64) float64

//go:linkname Asin github.com/goplus/js/primitive/math.Asin
func Asin(x float64) float64

//go:linkname Asinh github.com/goplus/js/primitive/math.Asinh
func Asinh(x float64) float64

//go:linkname Atan github.com/goplus/js/primitive/math.Atan
func Atan(x float64) float64

//go:linkname Atan2 github.com/goplus/js/primitive/math.Atan2
func Atan2(y, x float64) float64

//go:linkname Atanh github.com/goplus/js/primitive/math.Atanh
func Atanh(x float64) float64

//go:linkname Cbrt github.com/goplus/js/primitive/math.Cbrt
func Cbrt(x float64) float64

//go:linkname Ceil github.com/goplus/js/primitive/math.Ceil
func Ceil(x float64) float64

//go:linkname Cos github.com/goplus/js/primitive/math.Cos
func Cos(x float64) float64

//go:linkname Cosh github.com/goplus/js/primitive/math.Cosh
func Cosh(x float64) float64

//func Copysign(x, y float64) float64
//func Dim(x, y float64) float64
//func Erf(x float64) float64
//func Erfc(x float64) float64

//go:linkname Exp github.com/goplus/js/primitive/math.Exp
func Exp(x float64) float64

//func Exp2(x float64) float64

//go:linkname Expm1 github.com/goplus/js/primitive/math.Expm1
func Expm1(x float64) float64

//go:linkname Floor github.com/goplus/js/primitive/math.Floor
func Floor(x float64) float64

//func FMA(x, y, z float64) float64

//go:linkname Max github.com/goplus/js/primitive/math.Max
func Max(x, y float64) float64

//go:linkname Min github.com/goplus/js/primitive/math.Min
func Min(x, y float64) float64

//func Mod(x, y float64) float64
//func Frexp(f float64) (float64, int)
//func Gamma(x float64) float64

//go:linkname Hypot github.com/goplus/js/primitive/math.Hypot
func Hypot(x, y float64) float64

//func Ilogb(x float64) int
//func J0(x float64) float64
//func J1(x float64) float64
//func Jn(n int, x float64) float64
//func Ldexp(x float64, exp int) float64
//func Lgamma(x float64) (lgamma float64, sign int)

//go:linkname Log github.com/goplus/js/primitive/math.Log
func Log(x float64) float64

//go:linkname Log10 github.com/goplus/js/primitive/math.Log10
func Log10(x float64) float64

//go:linkname Log1p github.com/goplus/js/primitive/math.Log1p
func Log1p(x float64) float64

//go:linkname Log2 github.com/goplus/js/primitive/math.Log2
func Log2(x float64) float64

//func Logb(x float64) float64
//func Modf(f float64) (float64, float64)
//func Nextafter(x, y float64) float64

//go:linkname Pow github.com/goplus/js/primitive/math.Pow
func Pow(x, y float64) float64

//func Remainder(x, y float64) float64

//go:linkname Round github.com/goplus/js/primitive/math.Round
func Round(x float64) float64

//go:linkname Sin github.com/goplus/js/primitive/math.Sin
func Sin(x float64) float64

//go:linkname Sinh github.com/goplus/js/primitive/math.Sinh
func Sinh(x float64) float64

//go:linkname Sqrt github.com/goplus/js/primitive/math.Sqrt
func Sqrt(x float64) float64

//go:linkname Tan github.com/goplus/js/primitive/math.Tan
func Tan(x float64) float64

//go:linkname Tanh github.com/goplus/js/primitive/math.Tanh
func Tanh(x float64) float64

//go:linkname Trunc github.com/goplus/js/primitive/math.Trunc
func Trunc(x float64) float64

//func Tgamma(x float64) float64

// ----------------------------------------------------------------------------
