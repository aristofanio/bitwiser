package bitwiser

import (
	"testing"
)


func TestSimpleBitwiseOp_And(t *testing.T) {
	op := NewBitOperation()
	//
	inOut := [][]byte{
		{0x10, 0x10, 0x10},
		{0x11, 0x10, 0x10},
	}
	//
	for k, v := range  inOut {
		if op.And(v[0], v[1]) != v[2] {
			t.Errorf("Test#%d: %s with error", k, "op.And")
		}
	}

}

func TestSimpleBitwiseOp_Or(t *testing.T) {
	op := NewBitOperation()
	//
	inOut := [][]byte{
		{0x10, 0x10, 0x10},
		{0x11, 0x10, 0x11},
	}
	//
	for k, v := range  inOut {
		if op.Or(v[0], v[1]) != v[2] {
			t.Errorf("Test#%d: %s with error", k, "op.Or")
		}
	}

}

func TestSimpleBitwiseOp_Xor(t *testing.T) {
	op := NewBitOperation()
	//
	inOut := [][]byte{
		{0x10, 0x10, 0x00},
		{0x11, 0x10, 0x01},
	}
	//
	for k, v := range  inOut {
		if op.Xor(v[0], v[1]) != v[2] {
			t.Errorf("Test#%d: %s with error", k, "op.Xor")
		}
	}

}

func TestSimpleBitwiseOp_Not(t *testing.T) {
	op := NewBitOperation()
	//
	inOut := [][]byte{
		{0x10, 0xEF},
		{0x11, 0xEE},
	}
	//
	for k, v := range  inOut {
		if op.Not(v[0]) != v[1] {
			t.Errorf("Test#%d: %s with error", k, "op.Not")
		}
	}
}

func TestSimpleBitwiseOp_ShiftLeft(t *testing.T) {
	op := NewBitOperation()
	//
	inOut := []struct {
		a byte
		b uint
		c byte
	}{
		{0x10, 2, 0x40}, //0001 0000 --> (2) --> 0100 0000
		{0x11, 1, 0x22}, //0001 0001 --> (1) --> 0010 0010
		{0x11, 10, 0x00}, //0001 0001 --> (1) --> 0000 0000
	}
	//
	for k, v := range  inOut {
		if op.ShiftLeft(v.a, v.b) != v.c {
			t.Errorf("Test#%d: %s with error", k, "op.ShiftLeft")
		}
	}
}

func TestSimpleBitwiseOp_ShiftRight(t *testing.T) {
	op := NewBitOperation()
	//
	inOut := []struct {
		a byte
		b uint
		c byte
	}{
		{0x10, 2, 0x04}, //0001 0000 --> (2) --> 0000 0100
		{0x11, 1, 0x08}, //0001 0001 --> (1) --> 0000 1000
	}
	//
	for k, v := range  inOut {
		if op.ShiftRight(v.a, v.b) != v.c {
			t.Errorf("Test#%d: %s with error", k, "op.ShiftRight")
		}
	}
}

func TestBytewiseOp_ShiftLeft(t *testing.T) {
	op := NewByteOperation()
	//
	inOut := []struct {
		a Bytes
		b uint
		c Bytes
	}{
		{
			Bytes{[]byte{0x01}, 1}, 1,
			Bytes{[]byte{0x02}, 1}}, //0000 0001 --> (1) --> 0000 0010
		{
			Bytes{[]byte{0x00,0xF1}, 2}, 1,
			Bytes{[]byte{0x01, 0xe2}, 2}}, //0000 0000 1111 0001 --> (1) --> 0000 0001 1110 0010
		{
			Bytes{[]byte{0xf0,0xf9}, 2}, 1,
			Bytes{[]byte{0xe1, 0xf2}, 2}}, //1111 0000 1111 1001 --> (1) --> 1110 0001 1111 0010
		{
			Bytes{[]byte{0x2d,0xe9,0x2f,0xfe}, 4}, 2,
			Bytes{[]byte{0xb7,0xa4,0xbf,0xf8}, 4}}, //0x2DE92FFE --> (2) --> 0xB7A4BFF8
		{
			Bytes{[]byte{0xff,0xe9,0x2f,0xfe}, 4}, 3,
			Bytes{[]byte{0xff,0x49,0x7f,0xf0}, 4}}, //0xffE92FFE --> (2) --> 0xFF497FF0
	}
	//
	for k, v := range  inOut {
		//
		r, e := op.ShiftLeft(v.a, v.b)
		if e != nil {
			t.Errorf("Test#%d: %s with error %s", k, "op.ShiftLeft", e.Error())
		}
		//
		if r.ToInt() != v.c.ToInt() {
			t.Errorf("Test#%d: %s with error. Expected: 0x%x. Result: 0x%x", k, "op.ShiftLeft", v.c.b, r.b)
		} else {
			t.Logf("Test#%d: %s with success. Expected: 0x%x. Result: 0x%x", k, "op.ShiftLeft", v.c.b, r.b)
		}
	}
}


func TestBytewiseOp_ShiftRight(t *testing.T) {
	op := NewByteOperation()
	//
	inOut := []struct {
		a Bytes
		b uint
		c Bytes
	}{
		{
			Bytes{[]byte{0x01}, 1}, 1,
			Bytes{[]byte{0x00}, 1}}, //0000 0001 --> (1) --> 0000 0010
		{
			Bytes{[]byte{0x00,0xF1}, 2}, 1,
			Bytes{[]byte{0x00, 0x78}, 2}}, //0000 0000 1111 0001 --> (1) --> 0000 0001 1110 0010
		{
			Bytes{[]byte{0xf0,0xf9}, 2}, 1,
			Bytes{[]byte{0x78, 0x7c}, 2}}, //1111 0000 1111 1001 --> (1) --> 1110 0001 1111 0010
		{
			Bytes{[]byte{0x2d,0xe9,0x2f,0xfe}, 4}, 2,
			Bytes{[]byte{0x0b,0x7a,0x4b,0xff}, 4}}, //0x2DE92FFE --> (2) --> 0xB7A4BFF
		{
			Bytes{[]byte{0xff,0xe9,0x2f,0xfe}, 4}, 3,
			Bytes{[]byte{0x1f,0xfd,0x25,0xff}, 4}}, //0xffE92FFE --> (3) --> 0x1FFD25FF
	}
	//
	for k, v := range  inOut {
		//
		r, e := op.ShiftRight(v.a, v.b)
		if e != nil {
			t.Errorf("Test#%d: %s with error %s", k, "op.ShiftLeft", e.Error())
		}
		//
		if r.ToInt() != v.c.ToInt() {
			t.Errorf("Test#%d: %s with error. Expected: 0x%x. Result: 0x%x", k, "op.ShiftLeft", v.c.b, r.b)
		} else {
			t.Logf("Test#%d: %s with success. Expected: 0x%x. Result: 0x%x", k, "op.ShiftLeft", v.c.b, r.b)
		}
	}
}
