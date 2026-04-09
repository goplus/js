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

package primitive

// ----------------------------------------------------------------------------

type (
	String  = string
	Boolean = bool
)

// ----------------------------------------------------------------------------

type Number float64

func (a Number) XGo_Or(b Number) Number
func (a Number) XGo_And(b Number) Number
func (a Number) XGo_Xor(b Number) Number
func (a Number) XGo_AndNot(b Number) Number
func (a Number) XGo_Lsh(b Number) Number
func (a Number) XGo_Rsh(b Number) Number
func (a Number) XGo_Rem(b Number) Number
func (a Number) XGo_Neg() Number
func (a Number) XGo_Not() Number

// ----------------------------------------------------------------------------
