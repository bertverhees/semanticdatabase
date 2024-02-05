package primitives

// Package decimal implements an arbitrary precision fixed-point decimal.
//
// The zero-value of a Decimal is 0, as you would expect.
//
// The best way to create a new Decimal is to use decimal.NewFromString, ex:
//
//	n, err := decimal.NewFromString("-123.4567")
//	n.String() // output: "-123.4567"
//
// To use Decimal as part of a struct:
//
//	type StructName struct {
//	    Number Decimal
//	}
//
// Note: This can "only" represent numbers with a maximum of 2^31 digits after the decimal point.
import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

// DivisionPrecision is the number of decimal places in the result when it
// doesn't divide exactly.
//
// Example:
//
//	d1 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3))
//	d1.String() // output: "0.6666666666666667"
//	d2 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(30000))
//	d2.String() // output: "0.0000666666666667"
//	d3 := decimal.NewFromFloat(20000).Div(decimal.NewFromFloat(3))
//	d3.String() // output: "6666.6666666666666667"
//	decimal.DivisionPrecision = 3
//	d4 := decimal.NewFromFloat(2).Div(decimal.NewFromFloat(3))
//	d4.String() // output: "0.667"
var DivisionPrecision = 16

// MarshalJSONWithoutQuotes should be set to true if you want the decimal to
// be JSON marshaled as a number, instead of as a string.
// WARNING: this is dangerous for decimals with many digits, since many JSON
// unmarshallers (ex: Javascript's) will unmarshal JSON numbers to IEEE 754
// double-precision floating point numbers, which means you can potentially
// silently lose precision.
var MarshalJSONWithoutQuotes = false

// ExpMaxIterations specifies the maximum number of iterations needed to calculate
// precise natural exponent value using ExpHullAbrham method.
var ExpMaxIterations = 1000

// Zero constant, to make computations faster.
// Zero should never be compared with == or != directly, please use decimal.Equal or decimal.Cmp instead.
var Zero = New(0, 1)

var zeroInt = big.NewInt(0)
var oneInt = big.NewInt(1)
var twoInt = big.NewInt(2)
var fourInt = big.NewInt(4)
var fiveInt = big.NewInt(5)
var tenInt = big.NewInt(10)
var twentyInt = big.NewInt(20)

var factorials = []Decimal{New(1, 0)}

// Decimal represents a fixed-point decimal. It is immutable.
// number = value * 10 ^ exp
type Decimal struct {
	value *big.Int

	// NOTE(vadim): this must be an int32, because we cast it to float64 during
	// calculations. If exp is 64 bit, we might lose precision.
	// If we cared about being able to represent every possible decimal, we
	// could make exp a *big.Int but it would hurt performance and numbers
	// like that are unrealistic.
	exp int32
}

// New returns a new fixed-point decimal, value * 10 ^ exp.
func New(value int64, exp int32) Decimal {
	return Decimal{
		value: big.NewInt(value),
		exp:   exp,
	}
}

// NewFromInt converts an int64 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(value int64) Decimal {
	return Decimal{
		value: big.NewInt(value),
		exp:   0,
	}
}

// NewFromInt32 converts an int32 to Decimal.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt32(value int32) Decimal {
	return Decimal{
		value: big.NewInt(int64(value)),
		exp:   0,
	}
}

// NewFromBigInt returns a new Decimal from a big.Int, value * 10 ^ exp
func NewFromBigInt(value *big.Int, exp int32) Decimal {
	return Decimal{
		value: new(big.Int).Set(value),
		exp:   exp,
	}
}

// NewFromString returns a new Decimal from a string representation.
// Trailing zeroes are not trimmed.
//
// Example:
//
//	d, err := NewFromString("-123.45")
//	d2, err := NewFromString(".0001")
//	d3, err := NewFromString("1.47000")
func NewFromString(value string) (Decimal, error) {
	originalInput := value
	var intString string
	var exp int64

	// Check if number is using scientific notation
	eIndex := strings.IndexAny(value, "Ee")
	if eIndex != -1 {
		expInt, err := strconv.ParseInt(value[eIndex+1:], 10, 32)
		if err != nil {
			if e, ok := err.(*strconv.NumError); ok && e.Err == strconv.ErrRange {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", value)
			}
			return Decimal{}, fmt.Errorf("can't convert %s to decimal: exponent is not numeric", value)
		}
		value = value[:eIndex]
		exp = expInt
	}

	pIndex := -1
	vLen := len(value)
	for i := 0; i < vLen; i++ {
		if value[i] == '.' {
			if pIndex > -1 {
				return Decimal{}, fmt.Errorf("can't convert %s to decimal: too many .s", value)
			}
			pIndex = i
		}
	}

	if pIndex == -1 {
		// There is no decimal point, we can just parse the original string as
		// an int
		intString = value
	} else {
		if pIndex+1 < vLen {
			intString = value[:pIndex] + value[pIndex+1:]
		} else {
			intString = value[:pIndex]
		}
		expInt := -len(value[pIndex+1:])
		exp += int64(expInt)
	}

	var dValue *big.Int
	// strconv.ParseInt is faster than new(big.Int).SetString so this is just a shortcut for strings we know won't overflow
	if len(intString) <= 18 {
		parsed64, err := strconv.ParseInt(intString, 10, 64)
		if err != nil {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", value)
		}
		dValue = big.NewInt(parsed64)
	} else {
		dValue = new(big.Int)
		_, ok := dValue.SetString(intString, 10)
		if !ok {
			return Decimal{}, fmt.Errorf("can't convert %s to decimal", value)
		}
	}

	if exp < math.MinInt32 || exp > math.MaxInt32 {
		// NOTE(vadim): I doubt a string could realistically be this long
		return Decimal{}, fmt.Errorf("can't convert %s to decimal: fractional part too long", originalInput)
	}

	return Decimal{
		value: dValue,
		exp:   int32(exp),
	}, nil
}

// NewFromFormattedString returns a new Decimal from a formatted string representation.
// The second argument - replRegexp, is a regular expression that is used to find characters that should be
// removed from given decimal string representation. All matched characters will be replaced with an empty string.
//
// Example:
//
//	r := regexp.MustCompile("[$,]")
//	d1, err := NewFromFormattedString("$5,125.99", r)
//
//	r2 := regexp.MustCompile("[_]")
//	d2, err := NewFromFormattedString("1_000_000", r2)
//
//	r3 := regexp.MustCompile("[USD\\s]")
//	d3, err := NewFromFormattedString("5000 USD", r3)
func NewFromFormattedString(value string, replRegexp *regexp.Regexp) (Decimal, error) {
	parsedValue := replRegexp.ReplaceAllString(value, "")
	d, err := NewFromString(parsedValue)
	if err != nil {
		return Decimal{}, err
	}
	return d, nil
}

// RequireFromString returns a new Decimal from a string representation
// or panics if NewFromString had returned an error.
//
// Example:
//
//	d := RequireFromString("-123.45")
//	d2 := RequireFromString(".0001")
func RequireFromString(value string) Decimal {
	dec, err := NewFromString(value)
	if err != nil {
		panic(err)
	}
	return dec
}

// NewFromFloat converts a float64 to Decimal.
//
// The converted number will contain the number of significant digits that can be
// represented in a float with reliable roundtrip.
// This is typically 15 digits, but may be more in some cases.
// See https://www.exploringbinary.com/decimal-precision-of-binary-floating-point-numbers/ for more information.
//
// For slightly faster conversion, use NewFromFloatWithExponent where you can specify the precision in absolute terms.
//
// NOTE: this will panic on NaN, +/-inf
func NewFromFloat(value float64) Decimal {
	if value == 0 {
		return New(0, 0)
	}
	return newFromFloat(value, math.Float64bits(value), &float64info)
}

// NewFromFloat32 converts a float32 to Decimal.
//
// The converted number will contain the number of significant digits that can be
// represented in a float with reliable roundtrip.
// This is typically 6-8 digits depending on the input.
// See https://www.exploringbinary.com/decimal-precision-of-binary-floating-point-numbers/ for more information.
//
// For slightly faster conversion, use NewFromFloatWithExponent where you can specify the precision in absolute terms.
//
// NOTE: this will panic on NaN, +/-inf
func NewFromFloat32(value float32) Decimal {
	if value == 0 {
		return New(0, 0)
	}
	// XOR is workaround for https://github.com/golang/go/issues/26285
	a := math.Float32bits(value) ^ 0x80808080
	return newFromFloat(float64(value), uint64(a)^0x80808080, &float32info)
}

func newFromFloat(val float64, bits uint64, flt *floatInfo) Decimal {
	if math.IsNaN(val) || math.IsInf(val, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", val))
	}
	exp := int(bits>>flt.mantbits) & (1<<flt.expbits - 1)
	mant := bits & (uint64(1)<<flt.mantbits - 1)

	switch exp {
	case 0:
		// denormalized
		exp++

	default:
		// add implicit top bit
		mant |= uint64(1) << flt.mantbits
	}
	exp += flt.bias

	var d decimal
	d.Assign(mant)
	d.Shift(exp - int(flt.mantbits))
	d.neg = bits>>(flt.expbits+flt.mantbits) != 0

	roundShortest(&d, mant, exp, flt)
	// If less than 19 digits, we can do calculation in an int64.
	if d.nd < 19 {
		tmp := int64(0)
		m := int64(1)
		for i := d.nd - 1; i >= 0; i-- {
			tmp += m * int64(d.d[i]-'0')
			m *= 10
		}
		if d.neg {
			tmp *= -1
		}
		return Decimal{value: big.NewInt(tmp), exp: int32(d.dp) - int32(d.nd)}
	}
	dValue := new(big.Int)
	dValue, ok := dValue.SetString(string(d.d[:d.nd]), 10)
	if ok {
		return Decimal{value: dValue, exp: int32(d.dp) - int32(d.nd)}
	}

	return NewFromFloatWithExponent(val, int32(d.dp)-int32(d.nd))
}

// NewFromFloatWithExponent converts a float64 to Decimal, with an arbitrary
// number of fractional digits.
//
// Example:
//
//	NewFromFloatWithExponent(123.456, -2).String() // output: "123.46"
func NewFromFloatWithExponent(value float64, exp int32) Decimal {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		panic(fmt.Sprintf("Cannot create a Decimal from %v", value))
	}

	bits := math.Float64bits(value)
	mant := bits & (1<<52 - 1)
	exp2 := int32((bits >> 52) & (1<<11 - 1))
	sign := bits >> 63

	if exp2 == 0 {
		// specials
		if mant == 0 {
			return Decimal{}
		}
		// subnormal
		exp2++
	} else {
		// normal
		mant |= 1 << 52
	}

	exp2 -= 1023 + 52

	// normalizing base-2 values
	for mant&1 == 0 {
		mant = mant >> 1
		exp2++
	}

	// maximum number of fractional base-10 digits to represent 2^N exactly cannot be more than -N if N<0
	if exp < 0 && exp < exp2 {
		if exp2 < 0 {
			exp = exp2
		} else {
			exp = 0
		}
	}

	// representing 10^M * 2^N as 5^M * 2^(M+N)
	exp2 -= exp

	temp := big.NewInt(1)
	dMant := big.NewInt(int64(mant))

	// applying 5^M
	if exp > 0 {
		temp = temp.SetInt64(int64(exp))
		temp = temp.Exp(fiveInt, temp, nil)
	} else if exp < 0 {
		temp = temp.SetInt64(-int64(exp))
		temp = temp.Exp(fiveInt, temp, nil)
		dMant = dMant.Mul(dMant, temp)
		temp = temp.SetUint64(1)
	}

	// applying 2^(M+N)
	if exp2 > 0 {
		dMant = dMant.Lsh(dMant, uint(exp2))
	} else if exp2 < 0 {
		temp = temp.Lsh(temp, uint(-exp2))
	}

	// rounding and downscaling
	if exp > 0 || exp2 < 0 {
		halfDown := new(big.Int).Rsh(temp, 1)
		dMant = dMant.Add(dMant, halfDown)
		dMant = dMant.Quo(dMant, temp)
	}

	if sign == 1 {
		dMant = dMant.Neg(dMant)
	}

	return Decimal{
		value: dMant,
		exp:   exp,
	}
}

// Copy returns a copy of decimal with the same value and exponent, but a different pointer to value.
func (d Decimal) Copy() Decimal {
	d.ensureInitialized()
	return Decimal{
		value: new(big.Int).Set(d.value),
		exp:   d.exp,
	}
}

// rescale returns a rescaled version of the decimal. Returned
// decimal may be less precise if the given exponent is bigger
// than the initial exponent of the Decimal.
// NOTE: this will truncate, NOT round
//
// Example:
//
//	d := New(12345, -4)
//	d2 := d.rescale(-1)
//	d3 := d2.rescale(-4)
//	println(d1)
//	println(d2)
//	println(d3)
//
// Output:
//
//	1.2345
//	1.2
//	1.2000
func (d Decimal) rescale(exp int32) Decimal {
	d.ensureInitialized()

	if d.exp == exp {
		return Decimal{
			new(big.Int).Set(d.value),
			d.exp,
		}
	}

	// NOTE(vadim): must convert exps to float64 before - to prevent overflow
	diff := math.Abs(float64(exp) - float64(d.exp))
	value := new(big.Int).Set(d.value)

	expScale := new(big.Int).Exp(tenInt, big.NewInt(int64(diff)), nil)
	if exp > d.exp {
		value = value.Quo(value, expScale)
	} else if exp < d.exp {
		value = value.Mul(value, expScale)
	}

	return Decimal{
		value: value,
		exp:   exp,
	}
}

// Abs returns the absolute value of the decimal.
func (d Decimal) Abs() Decimal {
	if !d.IsNegative() {
		return d
	}
	d.ensureInitialized()
	d2Value := new(big.Int).Abs(d.value)
	return Decimal{
		value: d2Value,
		exp:   d.exp,
	}
}

// Add returns d + d2.
func (d Decimal) Add(d2 Decimal) Decimal {
	rd, rd2 := RescalePair(d, d2)

	d3Value := new(big.Int).Add(rd.value, rd2.value)
	return Decimal{
		value: d3Value,
		exp:   rd.exp,
	}
}

// Sub returns d - d2.
func (d Decimal) Sub(d2 Decimal) Decimal {
	rd, rd2 := RescalePair(d, d2)

	d3Value := new(big.Int).Sub(rd.value, rd2.value)
	return Decimal{
		value: d3Value,
		exp:   rd.exp,
	}
}

// Neg returns -d.
func (d Decimal) Neg() Decimal {
	d.ensureInitialized()
	val := new(big.Int).Neg(d.value)
	return Decimal{
		value: val,
		exp:   d.exp,
	}
}

// Mul returns d * d2.
func (d Decimal) Mul(d2 Decimal) Decimal {
	d.ensureInitialized()
	d2.ensureInitialized()

	expInt64 := int64(d.exp) + int64(d2.exp)
	if expInt64 > math.MaxInt32 || expInt64 < math.MinInt32 {
		// NOTE(vadim): better to panic than give incorrect results, as
		// Decimals are usually used for money
		panic(fmt.Sprintf("exponent %v overflows an int32!", expInt64))
	}

	d3Value := new(big.Int).Mul(d.value, d2.value)
	return Decimal{
		value: d3Value,
		exp:   int32(expInt64),
	}
}

// Shift shifts the decimal in base 10.
// It shifts left when shift is positive and right if shift is negative.
// In simpler terms, the given value for shift is added to the exponent
// of the decimal.
func (d Decimal) Shift(shift int32) Decimal {
	d.ensureInitialized()
	return Decimal{
		value: new(big.Int).Set(d.value),
		exp:   d.exp + shift,
	}
}

// Div returns d / d2. If it doesn't divide exactly, the result will have
// DivisionPrecision digits after the decimal point.
func (d Decimal) Div(d2 Decimal) Decimal {
	return d.DivRound(d2, int32(DivisionPrecision))
}

// QuoRem does division with remainder
// d.QuoRem(d2,precision) returns quotient q and remainder r such that
//
//	d = d2 * q + r, q an integer multiple of 10^(-precision)
//	0 <= r < abs(d2) * 10 ^(-precision) if d>=0
//	0 >= r > -abs(d2) * 10 ^(-precision) if d<0
//
// Note that precision<0 is allowed as input.
func (d Decimal) QuoRem(d2 Decimal, precision int32) (Decimal, Decimal) {
	d.ensureInitialized()
	d2.ensureInitialized()
	if d2.value.Sign() == 0 {
		panic("decimal division by 0")
	}
	scale := -precision
	e := int64(d.exp) - int64(d2.exp) - int64(scale)
	if e > math.MaxInt32 || e < math.MinInt32 {
		panic("overflow in decimal QuoRem")
	}
	var aa, bb, expo big.Int
	var scalerest int32
	// d = a 10^ea
	// d2 = b 10^eb
	if e < 0 {
		aa = *d.value
		expo.SetInt64(-e)
		bb.Exp(tenInt, &expo, nil)
		bb.Mul(d2.value, &bb)
		scalerest = d.exp
		// now aa = a
		//     bb = b 10^(scale + eb - ea)
	} else {
		expo.SetInt64(e)
		aa.Exp(tenInt, &expo, nil)
		aa.Mul(d.value, &aa)
		bb = *d2.value
		scalerest = scale + d2.exp
		// now aa = a ^ (ea - eb - scale)
		//     bb = b
	}
	var q, r big.Int
	q.QuoRem(&aa, &bb, &r)
	dq := Decimal{value: &q, exp: scale}
	dr := Decimal{value: &r, exp: scalerest}
	return dq, dr
}

// DivRound divides and rounds to a given precision
// i.e. to an integer multiple of 10^(-precision)
//
//	for a positive quotient digit 5 is rounded up, away from 0
//	if the quotient is negative then digit 5 is rounded down, away from 0
//
// Note that precision<0 is allowed as input.
func (d Decimal) DivRound(d2 Decimal, precision int32) Decimal {
	// QuoRem already checks initialization
	q, r := d.QuoRem(d2, precision)
	// the actual rounding decision is based on comparing r*10^precision and d2/2
	// instead compare 2 r 10 ^precision and d2
	var rv2 big.Int
	rv2.Abs(r.value)
	rv2.Lsh(&rv2, 1)
	// now rv2 = abs(r.value) * 2
	r2 := Decimal{value: &rv2, exp: r.exp + precision}
	// r2 is now 2 * r * 10 ^ precision
	var c = r2.Cmp(d2.Abs())

	if c < 0 {
		return q
	}

	if d.value.Sign()*d2.value.Sign() < 0 {
		return q.Sub(New(1, -precision))
	}

	return q.Add(New(1, -precision))
}

// Mod returns d % d2.
func (d Decimal) Mod(d2 Decimal) Decimal {
	_, r := d.QuoRem(d2, 0)
	return r
}

// Pow returns d to the power d2
func (d Decimal) Pow(d2 Decimal) Decimal {
	var temp Decimal
	if d2.IntPart() == 0 {
		return NewFromFloat(1)
	}
	temp = d.Pow(d2.Div(NewFromFloat(2)))
	if d2.IntPart()%2 == 0 {
		return temp.Mul(temp)
	}
	if d2.IntPart() > 0 {
		return temp.Mul(temp).Mul(d)
	}
	return temp.Mul(temp).Div(d)
}

// ExpHullAbrham calculates the natural exponent of decimal (e to the power of d) using Hull-Abraham algorithm.
// OverallPrecision argument specifies the overall precision of the result (integer part + decimal part).
//
// ExpHullAbrham is faster than ExpTaylor for small precision values, but it is much slower for large precision values.
//
// Example:
//
//	NewFromFloat(26.1).ExpHullAbrham(2).String()    // output: "220000000000"
//	NewFromFloat(26.1).ExpHullAbrham(20).String()   // output: "216314672147.05767284"
func (d Decimal) ExpHullAbrham(overallPrecision uint32) (Decimal, error) {
	// Algorithm based on Variable precision exponential function.
	// ACM Transactions on Mathematical Software by T. E. Hull & A. Abrham.
	if d.IsZero() {
		return Decimal{oneInt, 0}, nil
	}

	currentPrecision := overallPrecision

	// Algorithm does not work if currentPrecision * 23 < |x|.
	// Precision is automatically increased in such cases, so the value can be calculated precisely.
	// If newly calculated precision is higher than ExpMaxIterations the currentPrecision will not be changed.
	f := d.Abs().InexactFloat64()
	if ncp := f / 23; ncp > float64(currentPrecision) && ncp < float64(ExpMaxIterations) {
		currentPrecision = uint32(math.Ceil(ncp))
	}

	// fail if abs(d) beyond an over/underflow threshold
	overflowThreshold := New(23*int64(currentPrecision), 0)
	if d.Abs().Cmp(overflowThreshold) > 0 {
		return Decimal{}, fmt.Errorf("over/underflow threshold, exp(x) cannot be calculated precisely")
	}

	// Return 1 if abs(d) small enough; this also avoids later over/underflow
	overflowThreshold2 := New(9, -int32(currentPrecision)-1)
	if d.Abs().Cmp(overflowThreshold2) <= 0 {
		return Decimal{oneInt, d.exp}, nil
	}

	// t is the smallest integer >= 0 such that the corresponding abs(d/k) < 1
	t := d.exp + int32(d.NumDigits()) // Add d.NumDigits because the paper assumes that d.value [0.1, 1)

	if t < 0 {
		t = 0
	}

	k := New(1, t)                                     // reduction factor
	r := Decimal{new(big.Int).Set(d.value), d.exp - t} // reduced argument
	p := int32(currentPrecision) + t + 2               // precision for calculating the sum

	// Determine n, the number of therms for calculating sum
	// use first Newton step (1.435p - 1.182) / log10(p/abs(r))
	// for solving appropriate equation, along with directed
	// roundings and simple rational bound for log10(p/abs(r))
	rf := r.Abs().InexactFloat64()
	pf := float64(p)
	nf := math.Ceil((1.453*pf - 1.182) / math.Log10(pf/rf))
	if nf > float64(ExpMaxIterations) || math.IsNaN(nf) {
		return Decimal{}, fmt.Errorf("exact value cannot be calculated in <=ExpMaxIterations iterations")
	}
	n := int64(nf)

	tmp := New(0, 0)
	sum := New(1, 0)
	one := New(1, 0)
	for i := n - 1; i > 0; i-- {
		tmp.value.SetInt64(i)
		sum = sum.Mul(r.DivRound(tmp, p))
		sum = sum.Add(one)
	}

	ki := k.IntPart()
	res := New(1, 0)
	for i := ki; i > 0; i-- {
		res = res.Mul(sum)
	}

	resNumDigits := int32(res.NumDigits())

	var roundDigits int32
	if resNumDigits > abs(res.exp) {
		roundDigits = int32(currentPrecision) - resNumDigits - res.exp
	} else {
		roundDigits = int32(currentPrecision)
	}

	res = res.Round(roundDigits)

	return res, nil
}

// ExpTaylor calculates the natural exponent of decimal (e to the power of d) using Taylor series expansion.
// Precision argument specifies how precise the result must be (number of digits after decimal point).
// Negative precision is allowed.
//
// ExpTaylor is much faster for large precision values than ExpHullAbrham.
//
// Example:
//
//	d, err := NewFromFloat(26.1).ExpTaylor(2).String()
//	d.String()  // output: "216314672147.06"
//
//	NewFromFloat(26.1).ExpTaylor(20).String()
//	d.String()  // output: "216314672147.05767284062928674083"
//
//	NewFromFloat(26.1).ExpTaylor(-10).String()
//	d.String()  // output: "220000000000"
func (d Decimal) ExpTaylor(precision int32) (Decimal, error) {
	// Note(mwoss): Implementation can be optimized by exclusively using big.Int API only
	if d.IsZero() {
		return Decimal{oneInt, 0}.Round(precision), nil
	}

	var epsilon Decimal
	var divPrecision int32
	if precision < 0 {
		epsilon = New(1, -1)
		divPrecision = 8
	} else {
		epsilon = New(1, -precision-1)
		divPrecision = precision + 1
	}

	decAbs := d.Abs()
	pow := d.Abs()
	factorial := New(1, 0)

	result := New(1, 0)

	for i := int64(1); ; {
		step := pow.DivRound(factorial, divPrecision)
		result = result.Add(step)

		// Stop Taylor series when current step is smaller than epsilon
		if step.Cmp(epsilon) < 0 {
			break
		}

		pow = pow.Mul(decAbs)

		i++

		// Calculate next factorial number or retrieve cached value
		if len(factorials) >= int(i) && !factorials[i-1].IsZero() {
			factorial = factorials[i-1]
		} else {
			// To avoid any race conditions, firstly the zero value is appended to a slice to create
			// a spot for newly calculated factorial. After that, the zero value is replaced by calculated
			// factorial using the index notation.
			factorial = factorials[i-2].Mul(New(i, 0))
			factorials = append(factorials, Zero)
			factorials[i-1] = factorial
		}
	}

	if d.Sign() < 0 {
		result = New(1, 0).DivRound(result, precision+1)
	}

	result = result.Round(precision)
	return result, nil
}

// Ln calculates natural logarithm of d.
// Precision argument specifies how precise the result must be (number of digits after decimal point).
// Negative precision is allowed.
//
// Example:
//
//	d1, err := NewFromFloat(13.3).Ln(2)
//	d1.String()  // output: "2.59"
//
//	d2, err := NewFromFloat(579.161).Ln(10)
//	d2.String()  // output: "6.3615805046"
func (d Decimal) Ln(precision int32) (Decimal, error) {
	// Algorithm based on The Use of Iteration Methods for Approximating the Natural Logarithm,
	// James F. Epperson, The American Mathematical Monthly, Vol. 96, No. 9, November 1989, pp. 831-835.
	if d.IsNegative() {
		return Decimal{}, fmt.Errorf("cannot calculate natural logarithm for negative decimals")
	}

	if d.IsZero() {
		return Decimal{}, fmt.Errorf("cannot represent natural logarithm of 0, result: -infinity")
	}

	calcPrecision := precision + 2
	z := d.Copy()

	var comp1, comp3, comp2, comp4, reduceAdjust Decimal
	comp1 = z.Sub(Decimal{oneInt, 0})
	comp3 = Decimal{oneInt, -1}

	// for decimal in range [0.9, 1.1] where ln(d) is close to 0
	usePowerSeries := false

	if comp1.Abs().Cmp(comp3) <= 0 {
		usePowerSeries = true
	} else {
		// reduce input decimal to range [0.1, 1)
		expDelta := int32(z.NumDigits()) + z.exp
		z.exp -= expDelta

		// Input decimal was reduced by factor of 10^expDelta, thus we will need to add
		// ln(10^expDelta) = expDelta * ln(10)
		// to the result to compensate that
		ln10 := ln10.withPrecision(calcPrecision)
		reduceAdjust = NewFromInt32(expDelta)
		reduceAdjust = reduceAdjust.Mul(ln10)

		comp1 = z.Sub(Decimal{oneInt, 0})

		if comp1.Abs().Cmp(comp3) <= 0 {
			usePowerSeries = true
		} else {
			// initial estimate using floats
			zFloat := z.InexactFloat64()
			comp1 = NewFromFloat(math.Log(zFloat))
		}
	}

	epsilon := Decimal{oneInt, -calcPrecision}

	if usePowerSeries {
		// Power Series - https://en.wikipedia.org/wiki/Logarithm#Power_series
		// Calculating n-th term of formula: ln(z+1) = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
		// until the difference between current and next term is smaller than epsilon.
		// Coverage quite fast for decimals close to 1.0

		// z + 2
		comp2 = comp1.Add(Decimal{twoInt, 0})
		// z / (z + 2)
		comp3 = comp1.DivRound(comp2, calcPrecision)
		// 2 * (z / (z + 2))
		comp1 = comp3.Add(comp3)
		comp2 = comp1.Copy()

		for n := 1; ; n++ {
			// 2 * (z / (z+2))^(2n+1)
			comp2 = comp2.Mul(comp3).Mul(comp3)

			// 1 / (2n+1) * 2 * (z / (z+2))^(2n+1)
			comp4 = NewFromInt(int64(2*n + 1))
			comp4 = comp2.DivRound(comp4, calcPrecision)

			// comp1 = 2 sum [ 1 / (2n+1) * (z / (z+2))^(2n+1) ]
			comp1 = comp1.Add(comp4)

			if comp4.Abs().Cmp(epsilon) <= 0 {
				break
			}
		}
	} else {
		// Halley's Iteration.
		// Calculating n-th term of formula: a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z),
		// until the difference between current and next term is smaller than epsilon
		for {
			// exp(a_n)
			comp3, _ = comp1.ExpTaylor(calcPrecision)
			// exp(a_n) - z
			comp2 = comp3.Sub(z)
			// 2 * (exp(a_n) - z)
			comp2 = comp2.Add(comp2)
			// exp(a_n) + z
			comp4 = comp3.Add(z)
			// 2 * (exp(a_n) - z) / (exp(a_n) + z)
			comp3 = comp2.DivRound(comp4, calcPrecision)
			// comp1 = a_(n+1) = a_n - 2 * (exp(a_n) - z) / (exp(a_n) + z)
			comp1 = comp1.Sub(comp3)

			if comp3.Abs().Cmp(epsilon) <= 0 {
				break
			}
		}
	}

	comp1 = comp1.Add(reduceAdjust)

	return comp1.Round(precision), nil
}

// NumDigits returns the number of digits of the decimal coefficient (d.Value)
// Note: Current implementation is extremely slow for large decimals and/or decimals with large fractional part
func (d Decimal) NumDigits() int {
	d.ensureInitialized()
	// Note(mwoss): It can be optimized, unnecessary cast of big.Int to string
	if d.IsNegative() {
		return len(d.value.String()) - 1
	}
	return len(d.value.String())
}

// IsInteger returns true when decimal can be represented as an integer value, otherwise, it returns false.
func (d Decimal) IsInteger() bool {
	// The most typical case, all decimal with exponent higher or equal 0 can be represented as integer
	if d.exp >= 0 {
		return true
	}
	// When the exponent is negative we have to check every number after the decimal place
	// If all of them are zeroes, we are sure that given decimal can be represented as an integer
	var r big.Int
	q := new(big.Int).Set(d.value)
	for z := abs(d.exp); z > 0; z-- {
		q.QuoRem(q, tenInt, &r)
		if r.Cmp(zeroInt) != 0 {
			return false
		}
	}
	return true
}

// Abs calculates absolute value of any int32. Used for calculating absolute value of decimal's exponent.
func abs(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (d Decimal) Cmp(d2 Decimal) int {
	d.ensureInitialized()
	d2.ensureInitialized()

	if d.exp == d2.exp {
		return d.value.Cmp(d2.value)
	}

	rd, rd2 := RescalePair(d, d2)

	return rd.value.Cmp(rd2.value)
}

// Compare compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (d Decimal) Compare(d2 Decimal) int {
	return d.Cmp(d2)
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (d Decimal) Equal(d2 Decimal) bool {
	return d.Cmp(d2) == 0
}

// Deprecated: Equals is deprecated, please use Equal method instead.
func (d Decimal) Equals(d2 Decimal) bool {
	return d.Equal(d2)
}

// GreaterThan (GT) returns true when d is greater than d2.
func (d Decimal) GreaterThan(d2 Decimal) bool {
	return d.Cmp(d2) == 1
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (d Decimal) GreaterThanOrEqual(d2 Decimal) bool {
	cmp := d.Cmp(d2)
	return cmp == 1 || cmp == 0
}

// LessThan (LT) returns true when d is less than d2.
func (d Decimal) LessThan(d2 Decimal) bool {
	return d.Cmp(d2) == -1
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (d Decimal) LessThanOrEqual(d2 Decimal) bool {
	cmp := d.Cmp(d2)
	return cmp == -1 || cmp == 0
}

// Sign returns:
//
//	-1 if d <  0
//	 0 if d == 0
//	+1 if d >  0
func (d Decimal) Sign() int {
	if d.value == nil {
		return 0
	}
	return d.value.Sign()
}

// IsPositive return
//
//	true if d > 0
//	false if d == 0
//	false if d < 0
func (d Decimal) IsPositive() bool {
	return d.Sign() == 1
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (d Decimal) IsNegative() bool {
	return d.Sign() == -1
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (d Decimal) IsZero() bool {
	return d.Sign() == 0
}

// Exponent returns the exponent, or scale component of the decimal.
func (d Decimal) Exponent() int32 {
	return d.exp
}

// Coefficient returns the coefficient of the decimal. It is scaled by 10^Exponent()
func (d Decimal) Coefficient() *big.Int {
	d.ensureInitialized()
	// we copy the coefficient so that mutating the result does not mutate the Decimal.
	return new(big.Int).Set(d.value)
}

// CoefficientInt64 returns the coefficient of the decimal as int64. It is scaled by 10^Exponent()
// If coefficient cannot be represented in an int64, the result will be undefined.
func (d Decimal) CoefficientInt64() int64 {
	d.ensureInitialized()
	return d.value.Int64()
}

// IntPart returns the integer component of the decimal.
func (d Decimal) IntPart() int64 {
	scaledD := d.rescale(0)
	return scaledD.value.Int64()
}

// BigInt returns integer component of the decimal as a BigInt.
func (d Decimal) BigInt() *big.Int {
	scaledD := d.rescale(0)
	i := &big.Int{}
	i.SetString(scaledD.String(), 10)
	return i
}

// BigFloat returns decimal as BigFloat.
// Be aware that casting decimal to BigFloat might cause a loss of precision.
func (d Decimal) BigFloat() *big.Float {
	f := &big.Float{}
	f.SetString(d.String())
	return f
}

// Rat returns a rational number representation of the decimal.
func (d Decimal) Rat() *big.Rat {
	d.ensureInitialized()
	if d.exp <= 0 {
		// NOTE(vadim): must negate after casting to prevent int32 overflow
		denom := new(big.Int).Exp(tenInt, big.NewInt(-int64(d.exp)), nil)
		return new(big.Rat).SetFrac(d.value, denom)
	}

	mul := new(big.Int).Exp(tenInt, big.NewInt(int64(d.exp)), nil)
	num := new(big.Int).Mul(d.value, mul)
	return new(big.Rat).SetFrac(num, oneInt)
}

// Float64 returns the nearest float64 value for d and a bool indicating
// whether f represents d exactly.
// For more details, see the documentation for big.Rat.Float64
func (d Decimal) Float64() (f float64, exact bool) {
	return d.Rat().Float64()
}

// InexactFloat64 returns the nearest float64 value for d.
// It doesn't indicate if the returned value represents d exactly.
func (d Decimal) InexactFloat64() float64 {
	f, _ := d.Float64()
	return f
}

// String returns the string representation of the decimal
// with the fixed point.
//
// Example:
//
//	d := New(-12345, -3)
//	println(d.String())
//
// Output:
//
//	-12.345
func (d Decimal) String() string {
	return d.string(true)
}

// StringFixed returns a rounded fixed-point string with places digits after
// the decimal point.
//
// Example:
//
//	NewFromFloat(0).StringFixed(2) // output: "0.00"
//	NewFromFloat(0).StringFixed(0) // output: "0"
//	NewFromFloat(5.45).StringFixed(0) // output: "5"
//	NewFromFloat(5.45).StringFixed(1) // output: "5.5"
//	NewFromFloat(5.45).StringFixed(2) // output: "5.45"
//	NewFromFloat(5.45).StringFixed(3) // output: "5.450"
//	NewFromFloat(545).StringFixed(-1) // output: "550"
func (d Decimal) StringFixed(places int32) string {
	rounded := d.Round(places)
	return rounded.string(false)
}

// StringFixedBank returns a banker rounded fixed-point string with places digits
// after the decimal point.
//
// Example:
//
//	NewFromFloat(0).StringFixedBank(2) // output: "0.00"
//	NewFromFloat(0).StringFixedBank(0) // output: "0"
//	NewFromFloat(5.45).StringFixedBank(0) // output: "5"
//	NewFromFloat(5.45).StringFixedBank(1) // output: "5.4"
//	NewFromFloat(5.45).StringFixedBank(2) // output: "5.45"
//	NewFromFloat(5.45).StringFixedBank(3) // output: "5.450"
//	NewFromFloat(545).StringFixedBank(-1) // output: "540"
func (d Decimal) StringFixedBank(places int32) string {
	rounded := d.RoundBank(places)
	return rounded.string(false)
}

// StringFixedCash returns a Swedish/Cash rounded fixed-point string. For
// more details see the documentation at function RoundCash.
func (d Decimal) StringFixedCash(interval uint8) string {
	rounded := d.RoundCash(interval)
	return rounded.string(false)
}

// Round rounds the decimal to places decimal places.
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Example:
//
//	NewFromFloat(5.45).Round(1).String() // output: "5.5"
//	NewFromFloat(545).Round(-1).String() // output: "550"
func (d Decimal) Round(places int32) Decimal {
	if d.exp == -places {
		return d
	}
	// truncate to places + 1
	ret := d.rescale(-places - 1)

	// add sign(d) * 0.5
	if ret.value.Sign() < 0 {
		ret.value.Sub(ret.value, fiveInt)
	} else {
		ret.value.Add(ret.value, fiveInt)
	}

	// floor for positive numbers, ceil for negative numbers
	_, m := ret.value.DivMod(ret.value, tenInt, new(big.Int))
	ret.exp++
	if ret.value.Sign() < 0 && m.Cmp(zeroInt) != 0 {
		ret.value.Add(ret.value, oneInt)
	}

	return ret
}

// RoundCeil rounds the decimal towards +infinity.
//
// Example:
//
//	NewFromFloat(545).RoundCeil(-2).String()   // output: "600"
//	NewFromFloat(500).RoundCeil(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundCeil(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundCeil(1).String() // output: "-1.4"
func (d Decimal) RoundCeil(places int32) Decimal {
	if d.exp >= -places {
		return d
	}

	rescaled := d.rescale(-places)
	if d.Equal(rescaled) {
		return d
	}

	if d.value.Sign() > 0 {
		rescaled.value.Add(rescaled.value, oneInt)
	}

	return rescaled
}

// RoundFloor rounds the decimal towards -infinity.
//
// Example:
//
//	NewFromFloat(545).RoundFloor(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundFloor(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundFloor(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundFloor(1).String() // output: "-1.5"
func (d Decimal) RoundFloor(places int32) Decimal {
	if d.exp >= -places {
		return d
	}

	rescaled := d.rescale(-places)
	if d.Equal(rescaled) {
		return d
	}

	if d.value.Sign() < 0 {
		rescaled.value.Sub(rescaled.value, oneInt)
	}

	return rescaled
}

// RoundUp rounds the decimal away from zero.
//
// Example:
//
//	NewFromFloat(545).RoundUp(-2).String()   // output: "600"
//	NewFromFloat(500).RoundUp(-2).String()   // output: "500"
//	NewFromFloat(1.1001).RoundUp(2).String() // output: "1.11"
//	NewFromFloat(-1.454).RoundUp(1).String() // output: "-1.5"
func (d Decimal) RoundUp(places int32) Decimal {
	if d.exp >= -places {
		return d
	}

	rescaled := d.rescale(-places)
	if d.Equal(rescaled) {
		return d
	}

	if d.value.Sign() > 0 {
		rescaled.value.Add(rescaled.value, oneInt)
	} else if d.value.Sign() < 0 {
		rescaled.value.Sub(rescaled.value, oneInt)
	}

	return rescaled
}

// RoundDown rounds the decimal towards zero.
//
// Example:
//
//	NewFromFloat(545).RoundDown(-2).String()   // output: "500"
//	NewFromFloat(-500).RoundDown(-2).String()   // output: "-500"
//	NewFromFloat(1.1001).RoundDown(2).String() // output: "1.1"
//	NewFromFloat(-1.454).RoundDown(1).String() // output: "-1.4"
func (d Decimal) RoundDown(places int32) Decimal {
	if d.exp >= -places {
		return d
	}

	rescaled := d.rescale(-places)
	if d.Equal(rescaled) {
		return d
	}
	return rescaled
}

// RoundBank rounds the decimal to places decimal places.
// If the final digit to round is equidistant from the nearest two integers the
// rounded value is taken as the even number
//
// If places < 0, it will round the integer part to the nearest 10^(-places).
//
// Examples:
//
//	NewFromFloat(5.45).RoundBank(1).String() // output: "5.4"
//	NewFromFloat(545).RoundBank(-1).String() // output: "540"
//	NewFromFloat(5.46).RoundBank(1).String() // output: "5.5"
//	NewFromFloat(546).RoundBank(-1).String() // output: "550"
//	NewFromFloat(5.55).RoundBank(1).String() // output: "5.6"
//	NewFromFloat(555).RoundBank(-1).String() // output: "560"
func (d Decimal) RoundBank(places int32) Decimal {

	round := d.Round(places)
	remainder := d.Sub(round).Abs()

	half := New(5, -places-1)
	if remainder.Cmp(half) == 0 && round.value.Bit(0) != 0 {
		if round.value.Sign() < 0 {
			round.value.Add(round.value, oneInt)
		} else {
			round.value.Sub(round.value, oneInt)
		}
	}

	return round
}

// RoundCash aka Cash/Penny/öre rounding rounds decimal to a specific
// interval. The amount payable for a cash transaction is rounded to the nearest
// multiple of the minimum currency unit available. The following intervals are
// available: 5, 10, 25, 50 and 100; any other number throws a panic.
//
//	  5:   5 cent rounding 3.43 => 3.45
//	 10:  10 cent rounding 3.45 => 3.50 (5 gets rounded up)
//	 25:  25 cent rounding 3.41 => 3.50
//	 50:  50 cent rounding 3.75 => 4.00
//	100: 100 cent rounding 3.50 => 4.00
//
// For more details: https://en.wikipedia.org/wiki/Cash_rounding
func (d Decimal) RoundCash(interval uint8) Decimal {
	var iVal *big.Int
	switch interval {
	case 5:
		iVal = twentyInt
	case 10:
		iVal = tenInt
	case 25:
		iVal = fourInt
	case 50:
		iVal = twoInt
	case 100:
		iVal = oneInt
	default:
		panic(fmt.Sprintf("Decimal does not support this Cash rounding interval `%d`. Supported: 5, 10, 25, 50, 100", interval))
	}
	dVal := Decimal{
		value: iVal,
	}

	// TODO: optimize those calculations to reduce the high allocations (~29 allocs).
	return d.Mul(dVal).Round(0).Div(dVal).Truncate(2)
}

// Floor returns the nearest integer value less than or equal to d.
func (d Decimal) Floor() Decimal {
	d.ensureInitialized()

	if d.exp >= 0 {
		return d
	}

	exp := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	exp.Exp(exp, big.NewInt(-int64(d.exp)), nil)

	z := new(big.Int).Div(d.value, exp)
	return Decimal{value: z, exp: 0}
}

// Ceil returns the nearest integer value greater than or equal to d.
func (d Decimal) Ceil() Decimal {
	d.ensureInitialized()

	if d.exp >= 0 {
		return d
	}

	exp := big.NewInt(10)

	// NOTE(vadim): must negate after casting to prevent int32 overflow
	exp.Exp(exp, big.NewInt(-int64(d.exp)), nil)

	z, m := new(big.Int).DivMod(d.value, exp, new(big.Int))
	if m.Cmp(zeroInt) != 0 {
		z.Add(z, oneInt)
	}
	return Decimal{value: z, exp: 0}
}

// Truncate truncates off digits from the number, without rounding.
//
// NOTE: precision is the last digit that will not be truncated (must be >= 0).
//
// Example:
//
//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
func (d Decimal) Truncate(precision int32) Decimal {
	d.ensureInitialized()
	if precision >= 0 && -precision > d.exp {
		return d.rescale(-precision)
	}
	return d
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Decimal) UnmarshalJSON(decimalBytes []byte) error {
	if string(decimalBytes) == "null" {
		return nil
	}

	str, err := unquoteIfQuoted(decimalBytes)
	if err != nil {
		return fmt.Errorf("error decoding string '%s': %s", decimalBytes, err)
	}

	decimal, err := NewFromString(str)
	*d = decimal
	if err != nil {
		return fmt.Errorf("error decoding string '%s': %s", str, err)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (d Decimal) MarshalJSON() ([]byte, error) {
	var str string
	if MarshalJSONWithoutQuotes {
		str = d.String()
	} else {
		str = "\"" + d.String() + "\""
	}
	return []byte(str), nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface. As a string representation
// is already used when encoding to text, this method stores that string as []byte
func (d *Decimal) UnmarshalBinary(data []byte) error {
	// Verify we have at least 4 bytes for the exponent. The GOB encoded value
	// may be empty.
	if len(data) < 4 {
		return fmt.Errorf("error decoding binary %v: expected at least 4 bytes, got %d", data, len(data))
	}

	// Extract the exponent
	d.exp = int32(binary.BigEndian.Uint32(data[:4]))

	// Extract the value
	d.value = new(big.Int)
	if err := d.value.GobDecode(data[4:]); err != nil {
		return fmt.Errorf("error decoding binary %v: %s", data, err)
	}

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (d Decimal) MarshalBinary() (data []byte, err error) {
	// Write the exponent first since it's a fixed size
	v1 := make([]byte, 4)
	binary.BigEndian.PutUint32(v1, uint32(d.exp))

	// Add the value
	var v2 []byte
	if v2, err = d.value.GobEncode(); err != nil {
		return
	}

	// Return the byte array
	data = append(v1, v2...)
	return
}

// Scan implements the sql.Scanner interface for database deserialization.
func (d *Decimal) Scan(value interface{}) error {
	// first try to see if the data is stored in database as a Numeric datatype
	switch v := value.(type) {

	case float32:
		*d = NewFromFloat(float64(v))
		return nil

	case float64:
		// numeric in sqlite3 sends us float64
		*d = NewFromFloat(v)
		return nil

	case int64:
		// at least in sqlite3 when the value is 0 in db, the data is sent
		// to us as an int64 instead of a float64 ...
		*d = New(v, 0)
		return nil

	default:
		// default is trying to interpret value stored as string
		str, err := unquoteIfQuoted(v)
		if err != nil {
			return err
		}
		*d, err = NewFromString(str)
		return err
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (d Decimal) Value() (driver.Value, error) {
	return d.String(), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization.
func (d *Decimal) UnmarshalText(text []byte) error {
	str := string(text)

	dec, err := NewFromString(str)
	*d = dec
	if err != nil {
		return fmt.Errorf("error decoding string '%s': %s", str, err)
	}

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (d Decimal) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

// GobEncode implements the gob.GobEncoder interface for gob serialization.
func (d Decimal) GobEncode() ([]byte, error) {
	return d.MarshalBinary()
}

// GobDecode implements the gob.GobDecoder interface for gob serialization.
func (d *Decimal) GobDecode(data []byte) error {
	return d.UnmarshalBinary(data)
}

// StringScaled first scales the decimal then calls .String() on it.
//
// Deprecated: buggy and unintuitive. Use StringFixed instead.
func (d Decimal) StringScaled(exp int32) string {
	return d.rescale(exp).String()
}

func (d Decimal) string(trimTrailingZeros bool) string {
	if d.exp >= 0 {
		return d.rescale(0).value.String()
	}

	abs := new(big.Int).Abs(d.value)
	str := abs.String()

	var intPart, fractionalPart string

	// NOTE(vadim): this cast to int will cause bugs if d.exp == INT_MIN
	// and you are on a 32-bit machine. Won't fix this super-edge case.
	dExpInt := int(d.exp)
	if len(str) > -dExpInt {
		intPart = str[:len(str)+dExpInt]
		fractionalPart = str[len(str)+dExpInt:]
	} else {
		intPart = "0"

		num0s := -dExpInt - len(str)
		fractionalPart = strings.Repeat("0", num0s) + str
	}

	if trimTrailingZeros {
		i := len(fractionalPart) - 1
		for ; i >= 0; i-- {
			if fractionalPart[i] != '0' {
				break
			}
		}
		fractionalPart = fractionalPart[:i+1]
	}

	number := intPart
	if len(fractionalPart) > 0 {
		number += "." + fractionalPart
	}

	if d.value.Sign() < 0 {
		return "-" + number
	}

	return number
}

func (d *Decimal) ensureInitialized() {
	if d.value == nil {
		d.value = new(big.Int)
	}
}

// Min returns the smallest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Min(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Min with 0 arguments.
func Min(first Decimal, rest ...Decimal) Decimal {
	ans := first
	for _, item := range rest {
		if item.Cmp(ans) < 0 {
			ans = item
		}
	}
	return ans
}

// Max returns the largest Decimal that was passed in the arguments.
//
// To call this function with an array, you must do:
//
//	Max(arr[0], arr[1:]...)
//
// This makes it harder to accidentally call Max with 0 arguments.
func Max(first Decimal, rest ...Decimal) Decimal {
	ans := first
	for _, item := range rest {
		if item.Cmp(ans) > 0 {
			ans = item
		}
	}
	return ans
}

// Sum returns the combined total of the provided first and rest Decimals
func Sum(first Decimal, rest ...Decimal) Decimal {
	total := first
	for _, item := range rest {
		total = total.Add(item)
	}

	return total
}

// Avg returns the average value of the provided first and rest Decimals
func Avg(first Decimal, rest ...Decimal) Decimal {
	count := New(int64(len(rest)+1), 0)
	sum := Sum(first, rest...)
	return sum.Div(count)
}

// RescalePair rescales two decimals to common exponential value (minimal exp of both decimals)
func RescalePair(d1 Decimal, d2 Decimal) (Decimal, Decimal) {
	d1.ensureInitialized()
	d2.ensureInitialized()

	if d1.exp < d2.exp {
		return d1, d2.rescale(d1.exp)
	} else if d1.exp > d2.exp {
		return d1.rescale(d2.exp), d2
	}

	return d1, d2
}

func unquoteIfQuoted(value interface{}) (string, error) {
	var bytes []byte

	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'", value, value)
	}

	// If the amount is quoted, strip the quotes
	if len(bytes) > 2 && bytes[0] == '"' && bytes[len(bytes)-1] == '"' {
		bytes = bytes[1 : len(bytes)-1]
	}
	return string(bytes), nil
}

// NullDecimal represents a nullable decimal with compatibility for
// scanning null values from the database.
type NullDecimal struct {
	Decimal Decimal
	Valid   bool
}

func NewNullDecimal(d Decimal) NullDecimal {
	return NullDecimal{
		Decimal: d,
		Valid:   true,
	}
}

// Scan implements the sql.Scanner interface for database deserialization.
func (d *NullDecimal) Scan(value interface{}) error {
	if value == nil {
		d.Valid = false
		return nil
	}
	d.Valid = true
	return d.Decimal.Scan(value)
}

// Value implements the driver.Valuer interface for database serialization.
func (d NullDecimal) Value() (driver.Value, error) {
	if !d.Valid {
		return nil, nil
	}
	return d.Decimal.Value()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *NullDecimal) UnmarshalJSON(decimalBytes []byte) error {
	if string(decimalBytes) == "null" {
		d.Valid = false
		return nil
	}
	d.Valid = true
	return d.Decimal.UnmarshalJSON(decimalBytes)
}

// MarshalJSON implements the json.Marshaler interface.
func (d NullDecimal) MarshalJSON() ([]byte, error) {
	if !d.Valid {
		return []byte("null"), nil
	}
	return d.Decimal.MarshalJSON()
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for XML
// deserialization
func (d *NullDecimal) UnmarshalText(text []byte) error {
	str := string(text)

	// check for empty XML or XML without body e.g., <tag></tag>
	if str == "" {
		d.Valid = false
		return nil
	}
	if err := d.Decimal.UnmarshalText(text); err != nil {
		d.Valid = false
		return err
	}
	d.Valid = true
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for XML
// serialization.
func (d NullDecimal) MarshalText() (text []byte, err error) {
	if !d.Valid {
		return []byte{}, nil
	}
	return d.Decimal.MarshalText()
}

// Trig functions

// Atan returns the arctangent, in radians, of x.
func (d Decimal) Atan() Decimal {
	if d.Equal(NewFromFloat(0.0)) {
		return d
	}
	if d.GreaterThan(NewFromFloat(0.0)) {
		return d.satan()
	}
	return d.Neg().satan().Neg()
}

func (d Decimal) xatan() Decimal {
	P0 := NewFromFloat(-8.750608600031904122785e-01)
	P1 := NewFromFloat(-1.615753718733365076637e+01)
	P2 := NewFromFloat(-7.500855792314704667340e+01)
	P3 := NewFromFloat(-1.228866684490136173410e+02)
	P4 := NewFromFloat(-6.485021904942025371773e+01)
	Q0 := NewFromFloat(2.485846490142306297962e+01)
	Q1 := NewFromFloat(1.650270098316988542046e+02)
	Q2 := NewFromFloat(4.328810604912902668951e+02)
	Q3 := NewFromFloat(4.853903996359136964868e+02)
	Q4 := NewFromFloat(1.945506571482613964425e+02)
	z := d.Mul(d)
	b1 := P0.Mul(z).Add(P1).Mul(z).Add(P2).Mul(z).Add(P3).Mul(z).Add(P4).Mul(z)
	b2 := z.Add(Q0).Mul(z).Add(Q1).Mul(z).Add(Q2).Mul(z).Add(Q3).Mul(z).Add(Q4)
	z = b1.Div(b2)
	z = d.Mul(z).Add(d)
	return z
}

// satan reduces its argument (known to be positive)
// to the range [0, 0.66] and calls xatan.
func (d Decimal) satan() Decimal {
	Morebits := NewFromFloat(6.123233995736765886130e-17) // pi/2 = PIO2 + Morebits
	Tan3pio8 := NewFromFloat(2.41421356237309504880)      // tan(3*pi/8)
	pi := NewFromFloat(3.14159265358979323846264338327950288419716939937510582097494459)

	if d.LessThanOrEqual(NewFromFloat(0.66)) {
		return d.xatan()
	}
	if d.GreaterThan(Tan3pio8) {
		return pi.Div(NewFromFloat(2.0)).Sub(NewFromFloat(1.0).Div(d).xatan()).Add(Morebits)
	}
	return pi.Div(NewFromFloat(4.0)).Add((d.Sub(NewFromFloat(1.0)).Div(d.Add(NewFromFloat(1.0)))).xatan()).Add(NewFromFloat(0.5).Mul(Morebits))
}

// sin coefficients
var _sin = [...]Decimal{
	NewFromFloat(1.58962301576546568060e-10), // 0x3de5d8fd1fd19ccd
	NewFromFloat(-2.50507477628578072866e-8), // 0xbe5ae5e5a9291f5d
	NewFromFloat(2.75573136213857245213e-6),  // 0x3ec71de3567d48a1
	NewFromFloat(-1.98412698295895385996e-4), // 0xbf2a01a019bfdf03
	NewFromFloat(8.33333333332211858878e-3),  // 0x3f8111111110f7d0
	NewFromFloat(-1.66666666666666307295e-1), // 0xbfc5555555555548
}

// Sin returns the sine of the radian argument x.
func (d Decimal) Sin() Decimal {
	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if d.Equal(NewFromFloat(0.0)) {
		return d
	}
	// make argument positive but save the sign
	sign := false
	if d.LessThan(NewFromFloat(0.0)) {
		d = d.Neg()
		sign = true
	}

	j := d.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	y := NewFromFloat(float64(j)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if j&1 == 1 {
		j++
		y = y.Add(NewFromFloat(1.0))
	}
	j &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if j > 3 {
		sign = !sign
		j -= 4
	}
	z := d.Sub(y.Mul(PI4A)).Sub(y.Mul(PI4B)).Sub(y.Mul(PI4C)) // Extended precision modular arithmetic
	zz := z.Mul(z)

	if j == 1 || j == 2 {
		w := zz.Mul(zz).Mul(_cos[0].Mul(zz).Add(_cos[1]).Mul(zz).Add(_cos[2]).Mul(zz).Add(_cos[3]).Mul(zz).Add(_cos[4]).Mul(zz).Add(_cos[5]))
		y = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(zz)).Add(w)
	} else {
		y = z.Add(z.Mul(zz).Mul(_sin[0].Mul(zz).Add(_sin[1]).Mul(zz).Add(_sin[2]).Mul(zz).Add(_sin[3]).Mul(zz).Add(_sin[4]).Mul(zz).Add(_sin[5])))
	}
	if sign {
		y = y.Neg()
	}
	return y
}

// cos coefficients
var _cos = [...]Decimal{
	NewFromFloat(-1.13585365213876817300e-11), // 0xbda8fa49a0861a9b
	NewFromFloat(2.08757008419747316778e-9),   // 0x3e21ee9d7b4e3f05
	NewFromFloat(-2.75573141792967388112e-7),  // 0xbe927e4f7eac4bc6
	NewFromFloat(2.48015872888517045348e-5),   // 0x3efa01a019c844f5
	NewFromFloat(-1.38888888888730564116e-3),  // 0xbf56c16c16c14f91
	NewFromFloat(4.16666666666665929218e-2),   // 0x3fa555555555554b
}

// Cos returns the cosine of the radian argument x.
func (d Decimal) Cos() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	// make argument positive
	sign := false
	if d.LessThan(NewFromFloat(0.0)) {
		d = d.Neg()
	}

	j := d.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	y := NewFromFloat(float64(j)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if j&1 == 1 {
		j++
		y = y.Add(NewFromFloat(1.0))
	}
	j &= 7 // octant modulo 2Pi radians (360 degrees)
	// reflect in x axis
	if j > 3 {
		sign = !sign
		j -= 4
	}
	if j > 1 {
		sign = !sign
	}

	z := d.Sub(y.Mul(PI4A)).Sub(y.Mul(PI4B)).Sub(y.Mul(PI4C)) // Extended precision modular arithmetic
	zz := z.Mul(z)

	if j == 1 || j == 2 {
		y = z.Add(z.Mul(zz).Mul(_sin[0].Mul(zz).Add(_sin[1]).Mul(zz).Add(_sin[2]).Mul(zz).Add(_sin[3]).Mul(zz).Add(_sin[4]).Mul(zz).Add(_sin[5])))
	} else {
		w := zz.Mul(zz).Mul(_cos[0].Mul(zz).Add(_cos[1]).Mul(zz).Add(_cos[2]).Mul(zz).Add(_cos[3]).Mul(zz).Add(_cos[4]).Mul(zz).Add(_cos[5]))
		y = NewFromFloat(1.0).Sub(NewFromFloat(0.5).Mul(zz)).Add(w)
	}
	if sign {
		y = y.Neg()
	}
	return y
}

var _tanP = [...]Decimal{
	NewFromFloat(-1.30936939181383777646e+4), // 0xc0c992d8d24f3f38
	NewFromFloat(1.15351664838587416140e+6),  // 0x413199eca5fc9ddd
	NewFromFloat(-1.79565251976484877988e+7), // 0xc1711fead3299176
}
var _tanQ = [...]Decimal{
	NewFromFloat(1.00000000000000000000e+0),
	NewFromFloat(1.36812963470692954678e+4),  //0x40cab8a5eeb36572
	NewFromFloat(-1.32089234440210967447e+6), //0xc13427bc582abc96
	NewFromFloat(2.50083801823357915839e+7),  //0x4177d98fc2ead8ef
	NewFromFloat(-5.38695755929454629881e+7), //0xc189afe03cbe5a31
}

// Tan returns the tangent of the radian argument x.
func (d Decimal) Tan() Decimal {

	PI4A := NewFromFloat(7.85398125648498535156e-1)                             // 0x3fe921fb40000000, Pi/4 split into three parts
	PI4B := NewFromFloat(3.77489470793079817668e-8)                             // 0x3e64442d00000000,
	PI4C := NewFromFloat(2.69515142907905952645e-15)                            // 0x3ce8469898cc5170,
	M4PI := NewFromFloat(1.273239544735162542821171882678754627704620361328125) // 4/pi

	if d.Equal(NewFromFloat(0.0)) {
		return d
	}

	// make argument positive but save the sign
	sign := false
	if d.LessThan(NewFromFloat(0.0)) {
		d = d.Neg()
		sign = true
	}

	j := d.Mul(M4PI).IntPart()    // integer part of x/(Pi/4), as integer for tests on the phase angle
	y := NewFromFloat(float64(j)) // integer part of x/(Pi/4), as float

	// map zeros to origin
	if j&1 == 1 {
		j++
		y = y.Add(NewFromFloat(1.0))
	}

	z := d.Sub(y.Mul(PI4A)).Sub(y.Mul(PI4B)).Sub(y.Mul(PI4C)) // Extended precision modular arithmetic
	zz := z.Mul(z)

	if zz.GreaterThan(NewFromFloat(1e-14)) {
		w := zz.Mul(_tanP[0].Mul(zz).Add(_tanP[1]).Mul(zz).Add(_tanP[2]))
		x := zz.Add(_tanQ[1]).Mul(zz).Add(_tanQ[2]).Mul(zz).Add(_tanQ[3]).Mul(zz).Add(_tanQ[4])
		y = z.Add(z.Mul(w.Div(x)))
	} else {
		y = z
	}
	if j&2 == 2 {
		y = NewFromFloat(-1.0).Div(y)
	}
	if sign {
		y = y.Neg()
	}
	return y
}

type floatInfo struct {
	mantbits uint
	expbits  uint
	bias     int
}

var float32info = floatInfo{23, 8, -127}
var float64info = floatInfo{52, 11, -1023}

// roundShortest rounds d (= mant * 2^exp) to the shortest number of digits
// that will let the original floating point value be precisely reconstructed.
func roundShortest(d *decimal, mant uint64, exp int, flt *floatInfo) {
	// If mantissa is zero, the number is zero; stop now.
	if mant == 0 {
		d.nd = 0
		return
	}

	// Compute upper and lower such that any decimal number
	// between upper and lower (possibly inclusive)
	// will round to the original floating point number.

	// We may see at once that the number is already shortest.
	//
	// Suppose d is not denormal, so that 2^exp <= d < 10^dp.
	// The closest shorter number is at least 10^(dp-nd) away.
	// The lower/upper bounds computed below are at distance
	// at most 2^(exp-mantbits).
	//
	// So the number is already shortest if 10^(dp-nd) > 2^(exp-mantbits),
	// or equivalently log2(10)*(dp-nd) > exp-mantbits.
	// It is true if 332/100*(dp-nd) >= exp-mantbits (log2(10) > 3.32).
	minexp := flt.bias + 1 // minimum possible exponent
	if exp > minexp && 332*(d.dp-d.nd) >= 100*(exp-int(flt.mantbits)) {
		// The number is already shortest.
		return
	}

	// d = mant << (exp - mantbits)
	// Next highest floating point number is mant+1 << exp-mantbits.
	// Our upper bound is halfway between, mant*2+1 << exp-mantbits-1.
	upper := new(decimal)
	upper.Assign(mant*2 + 1)
	upper.Shift(exp - int(flt.mantbits) - 1)

	// d = mant << (exp - mantbits)
	// Next lowest floating point number is mant-1 << exp-mantbits,
	// unless mant-1 drops the significant bit and exp is not the minimum exp,
	// in which case the next lowest is mant*2-1 << exp-mantbits-1.
	// Either way, call it mantlo << explo-mantbits.
	// Our lower bound is halfway between, mantlo*2+1 << explo-mantbits-1.
	var mantlo uint64
	var explo int
	if mant > 1<<flt.mantbits || exp == minexp {
		mantlo = mant - 1
		explo = exp
	} else {
		mantlo = mant*2 - 1
		explo = exp - 1
	}
	lower := new(decimal)
	lower.Assign(mantlo*2 + 1)
	lower.Shift(explo - int(flt.mantbits) - 1)

	// The upper and lower bounds are possible outputs only if
	// the original mantissa is even, so that IEEE round-to-even
	// would round to the original mantissa and not the neighbors.
	inclusive := mant%2 == 0

	// As we walk the digits we want to know whether rounding up would fall
	// within the upper bound. This is tracked by upperdelta:
	//
	// If upperdelta == 0, the digits of d and upper are the same so far.
	//
	// If upperdelta == 1, we saw a difference of 1 between d and upper on a
	// previous digit and subsequently only 9s for d and 0s for upper.
	// (Thus rounding up may fall outside the bound, if it is exclusive.)
	//
	// If upperdelta == 2, then the difference is greater than 1
	// and we know that rounding up falls within the bound.
	var upperdelta uint8

	// Now we can figure out the minimum number of digits required.
	// Walk along until d has distinguished itself from upper and lower.
	for ui := 0; ; ui++ {
		// lower, d, and upper may have the decimal points at different
		// places. In this case upper is the longest, so we iterate from
		// ui==0 and start li and mi at (possibly) -1.
		mi := ui - upper.dp + d.dp
		if mi >= d.nd {
			break
		}
		li := ui - upper.dp + lower.dp
		l := byte('0') // lower digit
		if li >= 0 && li < lower.nd {
			l = lower.d[li]
		}
		m := byte('0') // middle digit
		if mi >= 0 {
			m = d.d[mi]
		}
		u := byte('0') // upper digit
		if ui < upper.nd {
			u = upper.d[ui]
		}

		// Okay to round down (truncate) if lower has a different digit
		// or if lower is inclusive and is exactly the result of rounding
		// down (i.e., and we have reached the final digit of lower).
		okdown := l != m || inclusive && li+1 == lower.nd

		switch {
		case upperdelta == 0 && m+1 < u:
			// Example:
			// m = 12345xxx
			// u = 12347xxx
			upperdelta = 2
		case upperdelta == 0 && m != u:
			// Example:
			// m = 12345xxx
			// u = 12346xxx
			upperdelta = 1
		case upperdelta == 1 && (m != '9' || u != '0'):
			// Example:
			// m = 1234598x
			// u = 1234600x
			upperdelta = 2
		}
		// Okay to round up if upper has a different digit and either upper
		// is inclusive or upper is bigger than the result of rounding up.
		okup := upperdelta > 0 && (inclusive || upperdelta > 1 || ui+1 < upper.nd)

		// If it's okay to do either, then round to the nearest one.
		// If it's okay to do only one, do it.
		switch {
		case okdown && okup:
			d.Round(mi + 1)
			return
		case okdown:
			d.RoundDown(mi + 1)
			return
		case okup:
			d.RoundUp(mi + 1)
			return
		}
	}
}

const (
	strLn10 = "2.302585092994045684017991454684364207601101488628772976033327900967572609677352480235997205089598298341967784042286248633409525465082806756666287369098781689482907208325554680843799894826233198528393505308965377732628846163366222287698219886746543667474404243274365155048934314939391479619404400222105101714174800368808401264708068556774321622835522011480466371565912137345074785694768346361679210180644507064800027750268491674655058685693567342067058113642922455440575892572420824131469568901675894025677631135691929203337658714166023010570308963457207544037084746994016826928280848118428931484852494864487192780967627127577539702766860595249671667418348570442250719796500471495105049221477656763693866297697952211071826454973477266242570942932258279850258550978526538320760672631716430950599508780752371033310119785754733154142180842754386359177811705430982748238504564801909561029929182431823752535770975053956518769751037497088869218020518933950723853920514463419726528728696511086257149219884997874887377134568620916705849807828059751193854445009978131146915934666241071846692310107598438319191292230792503747298650929009880391941702654416816335727555703151596113564846546190897042819763365836983716328982174407366009162177850541779276367731145041782137660111010731042397832521894898817597921798666394319523936855916447118246753245630912528778330963604262982153040874560927760726641354787576616262926568298704957954913954918049209069438580790032763017941503117866862092408537949861264933479354871737451675809537088281067452440105892444976479686075120275724181874989395971643105518848195288330746699317814634930000321200327765654130472621883970596794457943468343218395304414844803701305753674262153675579814770458031413637793236291560128185336498466942261465206459942072917119370602444929358037007718981097362533224548366988505528285966192805098447175198503666680874970496982273220244823343097169111136813588418696549323714996941979687803008850408979618598756579894836445212043698216415292987811742973332588607915912510967187510929248475023930572665446276200923068791518135803477701295593646298412366497023355174586195564772461857717369368404676577047874319780573853271810933883496338813069945569399346101090745616033312247949360455361849123333063704751724871276379140924398331810164737823379692265637682071706935846394531616949411701841938119405416449466111274712819705817783293841742231409930022911502362192186723337268385688273533371925103412930705632544426611429765388301822384091026198582888433587455960453004548370789052578473166283701953392231047527564998119228742789713715713228319641003422124210082180679525276689858180956119208391760721080919923461516952599099473782780648128058792731993893453415320185969711021407542282796298237068941764740642225757212455392526179373652434440560595336591539160312524480149313234572453879524389036839236450507881731359711238145323701508413491122324390927681724749607955799151363982881058285740538000653371655553014196332241918087621018204919492651483892"
)

var (
	ln10 = newConstApproximation(strLn10)
)

type constApproximation struct {
	exact          Decimal
	approximations []Decimal
}

func newConstApproximation(value string) constApproximation {
	parts := strings.Split(value, ".")
	coeff, fractional := parts[0], parts[1]

	coeffLen := len(coeff)
	maxPrecision := len(fractional)

	var approximations []Decimal
	for p := 1; p < maxPrecision; p *= 2 {
		r := RequireFromString(value[:coeffLen+p])
		approximations = append(approximations, r)
	}

	return constApproximation{
		RequireFromString(value),
		approximations,
	}
}

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Multiprecision decimal numbers.
// For floating-point formatting only; not general purpose.
// Only operations are assign and (binary) left/right shift.
// Can do binary floating point in multiprecision decimal precisely
// because 2 divides 10; cannot do decimal floating point
// in multiprecision binary precisely.

type decimal struct {
	d     [800]byte // digits, big-endian representation
	nd    int       // number of digits used
	dp    int       // decimal point
	neg   bool      // negative flag
	trunc bool      // discarded nonzero digits beyond d[:nd]
}

func (a *decimal) String() string {
	n := 10 + a.nd
	if a.dp > 0 {
		n += a.dp
	}
	if a.dp < 0 {
		n += -a.dp
	}

	buf := make([]byte, n)
	w := 0
	switch {
	case a.nd == 0:
		return "0"

	case a.dp <= 0:
		// zeros fill space between decimal point and digits
		buf[w] = '0'
		w++
		buf[w] = '.'
		w++
		w += digitZero(buf[w : w+-a.dp])
		w += copy(buf[w:], a.d[0:a.nd])

	case a.dp < a.nd:
		// decimal point in middle of digits
		w += copy(buf[w:], a.d[0:a.dp])
		buf[w] = '.'
		w++
		w += copy(buf[w:], a.d[a.dp:a.nd])

	default:
		// zeros fill space between digits and decimal point
		w += copy(buf[w:], a.d[0:a.nd])
		w += digitZero(buf[w : w+a.dp-a.nd])
	}
	return string(buf[0:w])
}

func digitZero(dst []byte) int {
	for i := range dst {
		dst[i] = '0'
	}
	return len(dst)
}

// trim trailing zeros from number.
// (They are meaningless; the decimal point is tracked
// independent of the number of digits.)
func trim(a *decimal) {
	for a.nd > 0 && a.d[a.nd-1] == '0' {
		a.nd--
	}
	if a.nd == 0 {
		a.dp = 0
	}
}

// Assign v to a.
func (a *decimal) Assign(v uint64) {
	var buf [24]byte

	// Write reversed decimal in buf.
	n := 0
	for v > 0 {
		v1 := v / 10
		v -= 10 * v1
		buf[n] = byte(v + '0')
		n++
		v = v1
	}

	// Reverse again to produce forward decimal in a.d.
	a.nd = 0
	for n--; n >= 0; n-- {
		a.d[a.nd] = buf[n]
		a.nd++
	}
	a.dp = a.nd
	trim(a)
}

// Maximum shift that we can do in one pass without overflow.
// A uint has 32 or 64 bits, and we have to be able to accommodate 9<<k.
const uintSize = 32 << (^uint(0) >> 63)
const maxShift = uintSize - 4

// Binary shift right (/ 2) by k bits.  k <= maxShift to avoid overflow.
func rightShift(a *decimal, k uint) {
	r := 0 // read pointer
	w := 0 // write pointer

	// Pick up enough leading digits to cover first shift.
	var n uint
	for ; n>>k == 0; r++ {
		if r >= a.nd {
			if n == 0 {
				// a == 0; shouldn't get here, but handle anyway.
				a.nd = 0
				return
			}
			for n>>k == 0 {
				n = n * 10
				r++
			}
			break
		}
		c := uint(a.d[r])
		n = n*10 + c - '0'
	}
	a.dp -= r - 1

	var mask uint = (1 << k) - 1

	// Pick up a digit, put down a digit.
	for ; r < a.nd; r++ {
		c := uint(a.d[r])
		dig := n >> k
		n &= mask
		a.d[w] = byte(dig + '0')
		w++
		n = n*10 + c - '0'
	}

	// Put down extra digits.
	for n > 0 {
		dig := n >> k
		n &= mask
		if w < len(a.d) {
			a.d[w] = byte(dig + '0')
			w++
		} else if dig > 0 {
			a.trunc = true
		}
		n = n * 10
	}

	a.nd = w
	trim(a)
}

// Cheat sheet for left shift: table indexed by shift count giving
// number of new digits that will be introduced by that shift.
//
// For example, leftcheats[4] = {2, "625"}.  That means that
// if we are shifting by 4 (multiplying by 16), it will add 2 digits
// when the string prefix is "625" through "999", and one fewer digit
// if the string prefix is "000" through "624".
//
// Credit for this trick goes to Ken.

type leftCheat struct {
	delta  int    // number of new digits
	cutoff string // minus one digit if original < a.
}

var leftcheats = []leftCheat{
	// Leading digits of 1/2^i = 5^i.
	// 5^23 is not an exact 64-bit floating point number,
	// so have to use bc for the math.
	// Go up to 60 to be large enough for 32bit and 64bit platforms.
	/*
		seq 60 | sed 's/^/5^/' | bc |
		awk 'BEGIN{ print "\t{ 0, \"\" }," }
		{
			log2 = log(2)/log(10)
			printf("\t{ %d, \"%s\" },\t// * %d\n",
				int(log2*NR+1), $0, 2**NR)
		}'
	*/
	{0, ""},
	{1, "5"},                                           // * 2
	{1, "25"},                                          // * 4
	{1, "125"},                                         // * 8
	{2, "625"},                                         // * 16
	{2, "3125"},                                        // * 32
	{2, "15625"},                                       // * 64
	{3, "78125"},                                       // * 128
	{3, "390625"},                                      // * 256
	{3, "1953125"},                                     // * 512
	{4, "9765625"},                                     // * 1024
	{4, "48828125"},                                    // * 2048
	{4, "244140625"},                                   // * 4096
	{4, "1220703125"},                                  // * 8192
	{5, "6103515625"},                                  // * 16384
	{5, "30517578125"},                                 // * 32768
	{5, "152587890625"},                                // * 65536
	{6, "762939453125"},                                // * 131072
	{6, "3814697265625"},                               // * 262144
	{6, "19073486328125"},                              // * 524288
	{7, "95367431640625"},                              // * 1048576
	{7, "476837158203125"},                             // * 2097152
	{7, "2384185791015625"},                            // * 4194304
	{7, "11920928955078125"},                           // * 8388608
	{8, "59604644775390625"},                           // * 16777216
	{8, "298023223876953125"},                          // * 33554432
	{8, "1490116119384765625"},                         // * 67108864
	{9, "7450580596923828125"},                         // * 134217728
	{9, "37252902984619140625"},                        // * 268435456
	{9, "186264514923095703125"},                       // * 536870912
	{10, "931322574615478515625"},                      // * 1073741824
	{10, "4656612873077392578125"},                     // * 2147483648
	{10, "23283064365386962890625"},                    // * 4294967296
	{10, "116415321826934814453125"},                   // * 8589934592
	{11, "582076609134674072265625"},                   // * 17179869184
	{11, "2910383045673370361328125"},                  // * 34359738368
	{11, "14551915228366851806640625"},                 // * 68719476736
	{12, "72759576141834259033203125"},                 // * 137438953472
	{12, "363797880709171295166015625"},                // * 274877906944
	{12, "1818989403545856475830078125"},               // * 549755813888
	{13, "9094947017729282379150390625"},               // * 1099511627776
	{13, "45474735088646411895751953125"},              // * 2199023255552
	{13, "227373675443232059478759765625"},             // * 4398046511104
	{13, "1136868377216160297393798828125"},            // * 8796093022208
	{14, "5684341886080801486968994140625"},            // * 17592186044416
	{14, "28421709430404007434844970703125"},           // * 35184372088832
	{14, "142108547152020037174224853515625"},          // * 70368744177664
	{15, "710542735760100185871124267578125"},          // * 140737488355328
	{15, "3552713678800500929355621337890625"},         // * 281474976710656
	{15, "17763568394002504646778106689453125"},        // * 562949953421312
	{16, "88817841970012523233890533447265625"},        // * 1125899906842624
	{16, "444089209850062616169452667236328125"},       // * 2251799813685248
	{16, "2220446049250313080847263336181640625"},      // * 4503599627370496
	{16, "11102230246251565404236316680908203125"},     // * 9007199254740992
	{17, "55511151231257827021181583404541015625"},     // * 18014398509481984
	{17, "277555756156289135105907917022705078125"},    // * 36028797018963968
	{17, "1387778780781445675529539585113525390625"},   // * 72057594037927936
	{18, "6938893903907228377647697925567626953125"},   // * 144115188075855872
	{18, "34694469519536141888238489627838134765625"},  // * 288230376151711744
	{18, "173472347597680709441192448139190673828125"}, // * 576460752303423488
	{19, "867361737988403547205962240695953369140625"}, // * 1152921504606846976
}

// Is the leading prefix of b lexicographically less than s?
func prefixIsLessThan(b []byte, s string) bool {
	for i := 0; i < len(s); i++ {
		if i >= len(b) {
			return true
		}
		if b[i] != s[i] {
			return b[i] < s[i]
		}
	}
	return false
}

// Binary shift left (* 2) by k bits.  k <= maxShift to avoid overflow.
func leftShift(a *decimal, k uint) {
	delta := leftcheats[k].delta
	if prefixIsLessThan(a.d[0:a.nd], leftcheats[k].cutoff) {
		delta--
	}

	r := a.nd         // read index
	w := a.nd + delta // write index

	// Pick up a digit, put down a digit.
	var n uint
	for r--; r >= 0; r-- {
		n += (uint(a.d[r]) - '0') << k
		quo := n / 10
		rem := n - 10*quo
		w--
		if w < len(a.d) {
			a.d[w] = byte(rem + '0')
		} else if rem != 0 {
			a.trunc = true
		}
		n = quo
	}

	// Put down extra digits.
	for n > 0 {
		quo := n / 10
		rem := n - 10*quo
		w--
		if w < len(a.d) {
			a.d[w] = byte(rem + '0')
		} else if rem != 0 {
			a.trunc = true
		}
		n = quo
	}

	a.nd += delta
	if a.nd >= len(a.d) {
		a.nd = len(a.d)
	}
	a.dp += delta
	trim(a)
}

// Binary shift left (k > 0) or right (k < 0).
func (a *decimal) Shift(k int) {
	switch {
	case a.nd == 0:
		// nothing to do: a == 0
	case k > 0:
		for k > maxShift {
			leftShift(a, maxShift)
			k -= maxShift
		}
		leftShift(a, uint(k))
	case k < 0:
		for k < -maxShift {
			rightShift(a, maxShift)
			k += maxShift
		}
		rightShift(a, uint(-k))
	}
}

// If we chop a at nd digits, should we round up?
func shouldRoundUp(a *decimal, nd int) bool {
	if nd < 0 || nd >= a.nd {
		return false
	}
	if a.d[nd] == '5' && nd+1 == a.nd { // exactly halfway - round to even
		// if we truncated, a little higher than what's recorded - always round up
		if a.trunc {
			return true
		}
		return nd > 0 && (a.d[nd-1]-'0')%2 != 0
	}
	// not halfway - digit tells all
	return a.d[nd] >= '5'
}

// Round a to nd digits (or fewer).
// If nd is zero, it means we're rounding
// just to the left of the digits, as in
// 0.09 -> 0.1.
func (a *decimal) Round(nd int) {
	if nd < 0 || nd >= a.nd {
		return
	}
	if shouldRoundUp(a, nd) {
		a.RoundUp(nd)
	} else {
		a.RoundDown(nd)
	}
}

// Round a down to nd digits (or fewer).
func (a *decimal) RoundDown(nd int) {
	if nd < 0 || nd >= a.nd {
		return
	}
	a.nd = nd
	trim(a)
}

// Round a up to nd digits (or fewer).
func (a *decimal) RoundUp(nd int) {
	if nd < 0 || nd >= a.nd {
		return
	}

	// round up
	for i := nd - 1; i >= 0; i-- {
		c := a.d[i]
		if c < '9' { // can stop after this digit
			a.d[i]++
			a.nd = i + 1
			return
		}
	}

	// Number is all 9s.
	// Change to single 1 with adjusted decimal point.
	a.d[0] = '1'
	a.nd = 1
	a.dp++
}

// Extract integer part, rounded appropriately.
// No guarantees about overflow.
func (a *decimal) RoundedInteger() uint64 {
	if a.dp > 20 {
		return 0xFFFFFFFFFFFFFFFF
	}
	var i int
	n := uint64(0)
	for i = 0; i < a.dp && i < a.nd; i++ {
		n = n*10 + uint64(a.d[i]-'0')
	}
	for ; i < a.dp; i++ {
		n *= 10
	}
	if shouldRoundUp(a, a.dp) {
		n++
	}
	return n
}
