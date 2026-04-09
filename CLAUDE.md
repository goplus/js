## Overview

It implements Go's built-in types targeting JavaScript as the compilation backend. The compilation model is [XGo](https://github.com/goplus/xgo) (formerly Go+), so operator overloading follows XGo's method-naming convention rather than Go's built-in operators.

## Repository Layout

```
.
├── primitive/  # JS primitive wrappers: Number, String, Boolean, …
└── builtin/    # Go built-in types exposed to XGo ← primary work area
```

## Type Mapping

Every Go built-in numeric/boolean/string type maps to a named type in `builtin`:

| Go type      | `builtin` type  |
|--------------|-----------------|
| `bool`       | `Bool`          |
| `int`        | `Int`           |
| `int8`       | `Int8`          |
| `int16`      | `Int16`         |
| `int32`      | `Int32`         |
| `int64`      | `Int64`         |
| `uint`       | `Uint`          |
| `uint8`      | `Uint8`         |
| `uint16`     | `Uint16`        |
| `uint32`     | `Uint32`        |
| `uint64`     | `Uint64`        |
| `uintptr`    | `Uintptr`       |
| `float32`    | `Float32`       |
| `float64`    | `Float64`       |
| `complex64`  | `Complex64`     |
| `complex128` | `Complex128`    |
| `string`     | `String`        |
| `byte`       | `Byte` (alias)  |
| `rune`       | `Rune` (alias)  |

## XGoJSPackage Packages

Some packages (e.g. `math`) are built-in JavaScript packages, identified by the `XGoJSPackage = true` constant. For these packages:

- **Do not implement functions in Go.** Only declare function prototypes (no function bodies).
- The XGo compiler maps the declared functions to their JavaScript equivalents at compile time.
- Constants should still be defined as Go constants.

## XGo Operator Overloading Convention

XGo maps Go operators to methods using the `XGo_` prefix. Every `builtin` type must implement the applicable subset below.

### Binary arithmetic (integers and floats)

| Go operator | Method name | Signature example                 |
|-------------|-------------|-----------------------------------|
| `+`         | `XGo_Add`   | `func (a Int) XGo_Add(b Int) Int` |
| `-`         | `XGo_Sub`   | `func (a Int) XGo_Sub(b Int) Int` |
| `*`         | `XGo_Mul`   | `func (a Int) XGo_Mul(b Int) Int` |
| `/`         | `XGo_Quo`   | `func (a Int) XGo_Quo(b Int) Int` |
| `%`         | `XGo_Rem`   | `func (a Int) XGo_Rem(b Int) Int` |

### Shift operators (integers only)
 
| Go operator | Method name | Signature example                      |
|-------------|-------------|----------------------------------------|
| `<<`        | `XGo_Lsh`   | `func (a Int) XGo_Lsh(b Uint) Int`     |
| `>>`        | `XGo_Rsh`   | `func (a Int) XGo_Rsh(b Uint) Int`     |

Shift count must be `Uint` (unsigned), matching the Go spec.

### Bitwise operators (integers only)

| Go operator | Method name  | Signature example                     |
|-------------|--------------|---------------------------------------|
| `\|`        | `XGo_Or`     | `func (a Int) XGo_Or(b Int) Int`      |
| `&`         | `XGo_And`    | `func (a Int) XGo_And(b Int) Int`     |
| `^`         | `XGo_Xor`    | `func (a Int) XGo_Xor(b Int) Int`     |
| `&^`        | `XGo_AndNot` | `func (a Int) XGo_AndNot(b Int) Int`  |

### Unary operators

| Go operator | Method name | Signature example                     |
|-------------|-------------|---------------------------------------|
| `-x`        | `XGo_Neg`   | `func (a Int) XGo_Neg() Int`          |
| `^x`        | `XGo_Not`   | `func (a Int) XGo_Not() Int`          |
| `!x`        | `XGo_LNot`  | `func (a Bool) XGo_LNot() Bool`       |

### Comparison operators
 
| Go operator | Method name | Signature example                     |
|-------------|-------------|---------------------------------------|
| `<`         | `XGo_LT`    | `func (a Int) XGo_LT(b Int) Bool`     |
| `<=`        | `XGo_LE`    | `func (a Int) XGo_LE(b Int) Bool`     |
| `>`         | `XGo_GT`    | `func (a Int) XGo_GT(b Int) Bool`     |
| `>=`        | `XGo_GE`    | `func (a Int) XGo_GE(b Int) Bool`     |
