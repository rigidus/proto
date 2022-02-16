package main

/*
     #cgo CFLAGS: -I${SRCDIR}/ctestlib
     #cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/ctestlib
     #cgo LDFLAGS: -L${SRCDIR}/ctestlib
     #cgo LDFLAGS: -ltest

     #include <test.h>
*/
import "C"
import (
    "fmt"
    "unsafe"
)

//go:linkname overflowError runtime.overflowError
var overflowError error

//go:linkname divideError runtime.divideError
var divideError error

// Int is represented as an array of 4 uint64, in little-endian order,
// so that Int[3] is the most significant, and Int[0] is the least significant
type uint256b [4]uint64

// NewInt returns a new zero-initialized uint256b
func NewUint256b() *uint256b {
    return &uint256b{}
}

// SetUint64 sets z to the value x
func (z *uint256b) SetUint64(x uint64) *uint256b {
    z[3], z[2], z[1], z[0] = 0, 0, 0, x
    return z
}

// IsZero returns true if z == 0
func (z * uint256b) IsZero() bool {
    return (z[0] | z[1] | z[2] | z[3]) == 0
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

// Add64 returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
//
// This function's execution time does not depend on the inputs.
func Add64(x, y, carry uint64) (sum, carryOut uint64) {
    sum = x + y + carry
    // The sum will overflow if both top bits are set (x & y) or if one of them
    // is (x | y), and a carry from the lower place happened. If such a carry
    // happens, the top bit will be 1 + 0 + 1 = 0 (&^ sum).
    carryOut = ((x & y) | ((x | y) &^ sum)) >> 63
    return
}

// umulHop computes (hi * 2^64 + lo) = z + (x * y)
func umulHop(z, x, y uint64) (hi, lo uint64) {
    hi, lo = Mul64(x, y)
    lo, carry := Add64(lo, z, 0)
    hi, _ = Add64(hi, 0, carry)
    return hi, lo
}

// umulStep computes (hi * 2^64 + lo) = z + (x * y) + carry.
func umulStep(z, x, y, carry uint64) (hi, lo uint64) {
    hi, lo = Mul64(x, y)
    lo, carry = Add64(lo, carry, 0)
    hi, _ = Add64(hi, 0, carry)
    lo, carry = Add64(lo, z, 0)
    hi, _ = Add64(hi, 0, carry)
    return hi, lo
}

// Mul sets z to the product x*y
func (z *uint256b) Mul(x, y *uint256b) *uint256b {
    var (
        res              uint256b
        carry            uint64
        res1, res2, res3 uint64
    )

    carry, res[0] = Mul64(x[0], y[0])
    carry, res1 = umulHop(carry, x[1], y[0])
    carry, res2 = umulHop(carry, x[2], y[0])
    res3 = x[3]*y[0] + carry

    carry, res[1] = umulHop(res1, x[0], y[1])
    carry, res2 = umulStep(res2, x[1], y[1], carry)
    res3 = res3 + x[2]*y[1] + carry

    carry, res[2] = umulHop(res2, x[0], y[2])
    res3 = res3 + x[1]*y[2] + carry

    res[3] = res3 + x[0]*y[3]

    return z.Set(&res)
}

// Set sets z to x and returns z.
func (z *uint256b) Set(x *uint256b) *uint256b {
    *z = *x
    return z
}

// Add sets z to the sum x+y
func (z *uint256b) Add(x, y *uint256b) *uint256b {
    var carry uint64
    z[0], carry = Add64(x[0], y[0], 0)
    z[1], carry = Add64(x[1], y[1], carry)
    z[2], carry = Add64(x[2], y[2], carry)
    z[3], _ = Add64(x[3], y[3], carry)
    return z
}

// Sub64 returns the difference of x, y and borrow: diff = x - y - borrow.
// The borrow input must be 0 or 1; otherwise the behavior is undefined.
// The borrowOut output is guaranteed to be 0 or 1.
//
// This function's execution time does not depend on the inputs.
func Sub64(x, y, borrow uint64) (diff, borrowOut uint64) {
    diff = x - y - borrow
    // See Sub32 for the bit logic.
    borrowOut = ((^x & y) | (^(x ^ y) & diff)) >> 63
    return
}

// Lt returns true if z < x
func (z *uint256b) Lt(x *uint256b) bool {
    // z < x <=> z - x < 0 i.e. when subtraction overflows.
    _, carry := Sub64(z[0], x[0], 0)
    _, carry = Sub64(z[1], x[1], carry)
    _, carry = Sub64(z[2], x[2], carry)
    _, carry = Sub64(z[3], x[3], carry)
    return carry != 0
}

// Gt returns true if z > x
func (z *uint256b) Gt(x *uint256b) bool {
    return x.Lt(z)
}

// Clear sets z to 0
func (z *uint256b) Clear() *uint256b {
    z[3], z[2], z[1], z[0] = 0, 0, 0, 0
    return z
}

// SetOne sets z to 1
func (z *uint256b) SetOne() *uint256b {
    z[3], z[2], z[1], z[0] = 0, 0, 0, 1
    return z
}

// Eq returns true if z == x
func (z *uint256b) Eq(x *uint256b) bool {
    return (z[0] == x[0]) && (z[1] == x[1]) && (z[2] == x[2]) && (z[3] == x[3])
}

// IsUint64 reports whether z can be represented as a uint64.
func (z *uint256b) IsUint64() bool {
    return (z[1] | z[2] | z[3]) == 0
}

// Uint64 returns the lower 64-bits of z
func (z *uint256b) Uint64() uint64 {
    return z[0]
}

var len8tab = [256]uint8{
    0x00, 0x01, 0x02, 0x02, 0x03, 0x03, 0x03, 0x03, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04, 0x04,
    0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
    0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
    0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06, 0x06,
    0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
    0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
    0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
    0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07, 0x07,
    0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
    0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
    0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
    0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
    0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
    0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
    0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
    0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08,
}

// Len64 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len64(x uint64) (n int) {
    if x >= 1<<32 {
        x >>= 32
        n = 32
    }
    if x >= 1<<16 {
        x >>= 16
        n += 16
    }
    if x >= 1<<8 {
        x >>= 8
        n += 8
    }
    return n + int(len8tab[x])
}

// Div64 returns the quotient and remainder of (hi, lo) divided by y:
// quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper
// half in parameter hi and the lower half in parameter lo.
// Div64 panics for y == 0 (division by zero) or y <= hi (quotient overflow).
func Div64(hi, lo, y uint64) (quo, rem uint64) {
    const (
        two32  = 1 << 32
        mask32 = two32 - 1
    )
    if y == 0 {
        panic(divideError)
    }
    if y <= hi {
        panic(overflowError)
    }

    s := uint(LeadingZeros64(y))
    y <<= s

    yn1 := y >> 32
    yn0 := y & mask32
    un32 := (hi) << (s) | (lo) >> (64-s)
    un10 := lo << s
    un1 := un10 >> 32
    un0 := un10 & mask32
    q1 := un32 / yn1
    rhat := un32 - q1*yn1

    for q1 >= two32 || q1*yn0 > two32*rhat+un1 {
        q1--
        rhat += yn1
        if rhat >= two32 {
            break
        }
    }

    un21 := un32*two32 + un1 - q1*y
    q0 := un21 / yn1
    rhat = un21 - q0*yn1

    for q0 >= two32 || q0*yn0 > two32*rhat+un0 {
        q0--
        rhat += yn1
        if rhat >= two32 {
            break
        }
    }

    return q1*two32 + q0, (un21*two32 + un0 - q0*y) >> s
}

// reciprocal2by1 computes <^d, ^0> / d.
func reciprocal2by1(d uint64) uint64 {
    reciprocal, _ := Div64(^d, ^uint64(0), d)
    return reciprocal
}

// subMulTo computes x -= y * multiplier.
// Requires len(x) >= len(y).
func subMulTo(x, y []uint64, multiplier uint64) uint64 {

    var borrow uint64
    for i := 0; i < len(y); i++ {
        s, carry1 := Sub64(x[i], borrow, 0)
        ph, pl := Mul64(y[i], multiplier)
        t, carry2 := Sub64(s, pl, 0)
        x[i] = t
        borrow = ph + carry1 + carry2
    }
    return borrow
}

// udivrem2by1 divides <uh, ul> / d and produces both quotient and remainder.
// It uses the provided d's reciprocal.
// Implementation ported from https://github.com/chfast/intx and is based on
// "Improved division by invariant integers", Algorithm 4.
func udivrem2by1(uh, ul, d, reciprocal uint64) (quot, rem uint64) {
    qh, ql := Mul64(reciprocal, uh)
    ql, carry := Add64(ql, ul, 0)
    qh, _ = Add64(qh, uh, carry)
    qh++

    r := ul - qh*d

    if r > ql {
        qh--
        r += d
    }

    if r >= d {
        qh++
        r -= d
    }

    return qh, r
}

// addTo computes x += y.
// Requires len(x) >= len(y).
func addTo(x, y []uint64) uint64 {
    var carry uint64
    for i := 0; i < len(y); i++ {
        x[i], carry = Add64(x[i], y[i], carry)
    }
    return carry
}

// udivremKnuth implements the division of u by normalized multiple word d from the Knuth's division algorithm.
// The quotient is stored in provided quot - len(u)-len(d) words.
// Updates u to contain the remainder - len(d) words.
func udivremKnuth(quot, u, d []uint64) {
    dh := d[len(d)-1]
    dl := d[len(d)-2]
    reciprocal := reciprocal2by1(dh)

    for j := len(u) - len(d) - 1; j >= 0; j-- {
        u2 := u[j+len(d)]
        u1 := u[j+len(d)-1]
        u0 := u[j+len(d)-2]

        var qhat, rhat uint64
        if u2 >= dh { // Division overflows.
            qhat = ^uint64(0)
            // TODO: Add "qhat one to big" adjustment (not needed for correctness, but helps avoiding "add back" case).
        } else {
            qhat, rhat = udivrem2by1(u2, u1, dh, reciprocal)
            ph, pl := Mul64(qhat, dl)
            if ph > rhat || (ph == rhat && pl > u0) {
                qhat--
                // TODO: Add "qhat one to big" adjustment (not needed for correctness, but helps avoiding "add back" case).
            }
        }

        // Multiply and subtract.
        borrow := subMulTo(u[j:], d, qhat)
        u[j+len(d)] = u2 - borrow
        if u2 < borrow { // Too much subtracted, add back.
            qhat--
            u[j+len(d)] += addTo(u[j:], d)
        }

        quot[j] = qhat // Store quotient digit.
    }
}

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int { return 64 - Len64(x) }

// udivremBy1 divides u by single normalized word d and produces both quotient and remainder.
// The quotient is stored in provided quot.
func udivremBy1(quot, u []uint64, d uint64) (rem uint64) {
    reciprocal := reciprocal2by1(d)
    rem = u[len(u)-1] // Set the top word as remainder.
    for j := len(u) - 2; j >= 0; j-- {
        quot[j], rem = udivrem2by1(rem, u[j], d, reciprocal)
    }
    return rem
}

// udivrem divides u by d and produces both quotient and remainder.
// The quotient is stored in provided quot - len(u)-len(d)+1 words.
// It loosely follows the Knuth's division algorithm (sometimes referenced as "schoolbook" division) using 64-bit words.
// See Knuth, Volume 2, section 4.3.1, Algorithm D.
func udivrem(quot, u []uint64, d *uint256b) (rem uint256b) {
    var dLen int
    for i := len(d) - 1; i >= 0; i-- {
        if d[i] != 0 {
            dLen = i + 1
            break
        }
    }

    shift := uint(LeadingZeros64(d[dLen-1]))

    var dnStorage uint256b
    dn := dnStorage[:dLen]
    for i := dLen - 1; i > 0; i-- {
        dn[i] = (d[i] << shift) | (d[i-1] >> (64 - shift))
    }
    dn[0] = d[0] << shift

    var uLen int
    for i := len(u) - 1; i >= 0; i-- {
        if u[i] != 0 {
            uLen = i + 1
            break
        }
    }

    var unStorage [9]uint64
    un := unStorage[:uLen+1]
    un[uLen] = u[uLen-1] >> (64 - shift)
    for i := uLen - 1; i > 0; i-- {
        un[i] = (u[i] << shift) | (u[i-1] >> (64 - shift))
    }
    un[0] = u[0] << shift

    // TODO: Skip the highest word of numerator if not significant.

    if dLen == 1 {
        r := udivremBy1(quot, un, dn[0])
        rem.SetUint64(r >> shift)
        return rem
    }

    udivremKnuth(quot, un, dn)

    for i := 0; i < dLen-1; i++ {
        rem[i] = (un[i] >> shift) | (un[i+1] << (64 - shift))
    }
    rem[dLen-1] = un[dLen-1] >> shift

    return rem
}

// Div sets z to the quotient x/y for returns z.
// If y == 0, z is set to 0
func (z *uint256b) Div(x, y *uint256b) *uint256b {
    if y.IsZero() || y.Gt(x) {
        return z.Clear()
    }
    if x.Eq(y) {
        return z.SetOne()
    }
    // Shortcut some cases
    if x.IsUint64() {
        return z.SetUint64(x.Uint64() / y.Uint64())
    }

    // At this point, we know
    // x/y ; x > y > 0

    var quot uint256b
    udivrem(quot[:], x[:], y)
    return z.Set(&quot)
}

func CalcAmountOut(x, y, a, mulv1, mulv2 *uint256b) *uint256b {
    if x.IsZero() && a.IsZero() {
        return NewUint256b().SetUint64(0)
    }

    /*
     *        a*y*997
     *     -------------
     *     x*1000 + a*997
     */
    var mulA, mulX uint256b
    mulA.Mul(a, mulv1)
    mulX.Mul(x, mulv2)      // vpmuldq // vpmuludq // vpmulld
    mulX.Add(&mulX, &mulA)	// vpaddq // vpmaxsw
    mulA.Mul(y, &mulA)      // vpmullq
    mulA.Div(&mulA, &mulX)
    return &mulA
}



func main() {
    var x *uint256b = NewUint256b().SetUint64(100)
    var y *uint256b = NewUint256b().SetUint64(200)
    var amount *uint256b = NewUint256b().SetUint64(400)
    var mulv1 *uint256b = NewUint256b().SetUint64(500)
    var mulv2 *uint256b = NewUint256b().SetUint64(600)
    fmt.Println(CalcAmountOut(x, y, amount, mulv1, mulv2))
    // fmt.Println(fmt.CalculateGas(fromToken, toToken, amount))
    fmt.Println("Send byte array")
    data := []byte{1, 4, 2}
    cBytes := (*C.uchar)(unsafe.Pointer(&data[0]))
    cBytesLength := C.size_t(len(data))
    fmt.Print("bytes: ")
    C.print_buffer(cBytes, cBytesLength)
    fmt.Println()
}
