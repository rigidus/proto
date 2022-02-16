// go tool compile -S asm.go > asm.asm
package main

import (
    "fmt"
    // "math/bits"
)

func myInput(msg string, p1 *uint64) {
    fmt.Print(msg)
    fmt.Scan(p1)
    fmt.Println()
}

// Mul64 returns the 128-bit product of x and y: (hi, lo) = x * y
// with the product bits' upper half returned in hi and the lower
// half returned in lo.
//
// This function's execution time does not depend on the inputs.
func Mul64(x, y uint64) (hi, lo uint64) {
    const mask32 = 1<<32 - 1
    x0 := x & mask32
    x1 := x >> 32
    y0 := y & mask32
    y1 := y >> 32
    w0 := x0 * y0
    t := x1*y0 + w0>>32
    w1 := t & mask32
    w2 := t >> 32
    w1 += x0 * y1
    hi = x1*y1 + w2 + w1>>32
    lo = x * y
    return
}

func myMul(x, y uint64) (uint64, uint64) {
    var hi, lo uint64
    hi, lo = Mul64(x, y)
    return hi, lo
}

func main() {
    var x, y, hi, lo uint64
    myInput("Enter X: ", &x)
    myInput("Enter Y: ", &y)
    hi, lo = myMul(x, y)
    fmt.Println(hi, lo)
}
