package bitwiser

import (
	"github.com/pkg/errors"
	"math"
	"fmt"
	"strings"
	"regexp"
	"strconv"
)

//===================================================
// BitOperation
//===================================================
type BitOperation interface{
	And(byte, byte)byte
	Or(byte, byte)byte
	Xor(byte, byte)byte
	Not(byte)byte
	ShiftLeft (byte, uint)byte
	ShiftRight (byte, uint)byte
}

type bitwiseOp struct {
}

//===================================================
// Methods bitwiseOp
//===================================================

func (o *bitwiseOp) And(a byte, b byte)byte{
	return a & b
}

func (o *bitwiseOp) Or(a byte, b byte)byte{
	return a | b
}

func (o *bitwiseOp) Xor(a byte, b byte)byte{
	return a ^ b
}

func (o *bitwiseOp) Not(a byte)byte{
	return ^a
}

func (o *bitwiseOp) ShiftLeft (a byte, n uint) byte{
	return a << n
}

func (o *bitwiseOp) ShiftRight (a byte, n uint)byte{
	return a >> n
}

//===================================================
// ByteOperation
//===================================================

var  ErrBadFormat = errors.New("Bad formatted bits representation")
var  ErrInMatchString = errors.New("Internal error in MatchString operation")

type Bytes struct {
	b []byte
	s int
}

type ByteOperation interface{
	And(Bytes, Bytes)(*Bytes, error)
	Or(Bytes, Bytes)(*Bytes, error)
	Xor(Bytes, Bytes)(*Bytes, error)
	Not(Bytes)(*Bytes, error)
	ShiftLeft (Bytes, uint)(*Bytes, error)
	ShiftRight (Bytes, uint)(*Bytes, error)
}

type bytewiseOp struct {
	*bitwiseOp
}

//===================================================
// Methods Bytes
//===================================================

func (y Bytes) ToInt() int {
	s := int(0)
	for i, v := range y.b {
		s += int(v)*int(math.Pow(256.0, float64(y.s-i-1)))
	}
	return s
}

func (y Bytes) Array() []byte {
	return y.b
}

func (y Bytes) ToString() string {
	return fmt.Sprintf("0x%x", y.b)
}

//===================================================
// Methods bytewiseOp
//===================================================

func (o *bytewiseOp) check(a Bytes, b Bytes) error{
	if a.s != b.s {
		return errors.New("Params 'a' and 'b' must be same size")
	}
	return nil
}

func (o *bytewiseOp)And(a Bytes, b Bytes)(*Bytes, error){
	//
	if err := o.check(a, b); err != nil {
		return nil, err
	}
	//
	r := make([]byte, a.s)
	for i := 0; i < a.s; i++ {
		r[i] = o.bitwiseOp.And(a.b[i], b.b[i])
	}
	//
	return &Bytes{b:r, s:len(r)}, nil
}

func (o *bytewiseOp)Or(a Bytes, b Bytes)(*Bytes, error){
	//
	if err := o.check(a, b); err != nil {
		return nil, err
	}
	//
	r := make([]byte, a.s)
	for i := 0; i < a.s; i++ {
		r[i] = o.bitwiseOp.Or(a.b[i], b.b[i])
	}
	//
	return &Bytes{b:r, s:len(r)}, nil
}

func (o *bytewiseOp)Xor(a Bytes, b Bytes)(*Bytes, error){
	//
	if err := o.check(a, b); err != nil {
		return nil, err
	}
	//
	r := make([]byte, a.s)
	for i := 0; i < a.s; i++ {
		r[i] = o.bitwiseOp.Xor(a.b[i], b.b[i])
	}
	//
	return &Bytes{b:r, s:len(r)}, nil
}

func (o *bytewiseOp)Not(a Bytes)(*Bytes, error){
	//
	r := make([]byte, a.s)
	for i := 0; i < a.s; i++ {
		r[i] = o.bitwiseOp.Not(a.b[i])
	}
	//
	return &Bytes{b:r, s:len(r)}, nil
}

func (o *bytewiseOp)ShiftLeft(a Bytes, n uint)(*Bytes, error){
	x := 0
	r := make([]byte, a.s)
	for i := a.s-1; i > -1; i-- {
		//calc
		y := int(a.b[i])
		p := x + y*int(math.Pow(2.0, float64(n)))
		//recalc
		x = p/256
		r[i] = byte(p - (x*256))
	}
	//
	return &Bytes{b:r, s:len(r)}, nil
}

func (o *bytewiseOp)ShiftRight (a Bytes, n uint)(*Bytes, error){
	x := 0
	r := make([]byte, a.s)
	for i := 0; i < a.s; i++ {
		//calc
		y := x*256 + int(a.b[i])
		p := y/int(math.Pow(2.0, float64(n)))
		//recalc
		x = y - p*int(math.Pow(2.0, float64(n)))
		r[i] = byte(p)
	}
	//
	return &Bytes{b:r, s:len(r)}, nil
}

//===================================================
// Public Operations
//===================================================

func NewBitOperation() BitOperation {
	return &bitwiseOp{}
}

func NewBytes(bs []byte) Bytes {
	return Bytes{bs, len(bs)}
}

func ParseFromBits(bs string) (Bytes, error) {
	//replace spaces
	nbs := strings.Replace(bs, " ", "", -1)
	//check match 01101...
	mtc, err := regexp.MatchString("^[01]*$", nbs)
	if err != nil {
		return Bytes{}, ErrInMatchString
	}
	if mtc == false {
		return Bytes{}, ErrBadFormat
	}
	//check len % 8 == 0
	diff := len(nbs) % 8
	if diff != 0 {
		nbs = strings.Repeat("0", 8-diff) + nbs
	}
	//split in 8 bits
	c := len(nbs)/8
	hexs := make([]string, c)
	for i:=0; i < c; i++{
		hexs[i] = nbs[i*8:(i*8)+8]
	}
	//parse string (base-2) to int
	byts := make([]byte, c)
	for i:=0; i < c; i++{
		v, _ := strconv.ParseInt(hexs[i], 2, 10)
		byts[i] = byte(v)
	}
	//result
	return Bytes{byts, len(byts)}, nil
}

func NewByteOperation() ByteOperation {
	return &bytewiseOp{&bitwiseOp{}}
}

