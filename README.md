# Bitwiser
Bitwiser is an simple library for operations bit-to-bit (bitwise) and byte-to-byte (here is call bytewise).

## Objective
Wrap operations bitwise and support a help for n-bytes operations.

## Structs and Interfaces

* BitOperation interface - allows bitwise operations using literal name (AND, OR, XOR, NOT, SHIFT_LEFT, SHIFT_RIGHT);
* ByteOperation interface - allows n-bytes operations using literal name (idem);
* Bytes struct - wrap an bytes array that is used in n-bytes operations.

## Structures

```
type BitOperation interface{
	And(byte, byte)byte
	Or(byte, byte)byte
	Xor(byte, byte)byte
	Not(byte)byte
	ShiftLeft (byte, uint)byte
	ShiftRight (byte, uint)byte
}

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
```

## Download
Download code source by command:
```
$ go get -u github.com/aristofanio/bitwiser
```

## Use
Create an BitOperation or an ByteOperation using the follows funcs:
```Go

bitOp := bitwiser.NewBitOperation()

//or

bytOp := bitwiser.NewByteOperation()

```

After just call methods and the params in byte or bytes(bitwiser.Bytes) according to your need.


## Examples

AND Bitwise operation:
```Go
package main

import (
  "fmt"
  "github.com/aristofanio/bitwiser"
)

func main(){
  op := bitwiser.NewBitOperation()
  rs := op.And(0x10, 0x10)
  println(fmt.Sprintf("0x%x", rs)) //output: 0x10
}

```

XOR and AND Bitwise operation:
```Go
package main

import (
  "fmt"
  "github.com/aristofanio/bitwiser"
)

func main(){
  op := bitwiser.NewBitOperation()
  rs := op.Xor(0x11, op.And(0x10, 0x10))
  println(fmt.Sprintf("0x%x", rs)) //output: 0x01 or 0x1
}
```

XOR and AND 'Bytewise' operation for long bytes like descript in
[Stackoverflow](https://stackoverflow.com/questions/28997600/golang-bitwise-operation-on-very-long-binary-bit-string-representation):
```Go
package main

import (
  "fmt"
  "github.com/aristofanio/bitwiser"
)

func main() {
  //
  x := []byte{0x01, 0x11, 0x00} // 0x011100
  y := []byte{0x00, 0x00, 0x11} // 0x000011
  //
  op := bitwiser.NewByteOperation()
  rs, _ := op.Or(bitwiser.NewBytes(x), bitwiser.NewBytes(y))
  println(fmt.Sprintf("0x%x", rs.Array()))  //output: 0x011111
  println(rs.ToString())                    //output: 0x011111
  println(rs.ToInt())                       //output: 69905 (base 10)
}
```

Parse Bytes from string (of bits)
```Go
package main

import (
	"github.com/aristofanio/bitwiser"
)

func main() {
	//
	b0, _ := bitwiser.ParseFromBits("0  011100")     //auto-complete for 8-multiple --> 0001 1100
	b1, _ := bitwiser.ParseFromBits("0001 1100")     //remove spaces                --> 0001 1100
	b2, _ := bitwiser.ParseFromBits("1 001 1100")    //remove spaces and fix size   --> 1001 1100
	b3, _ := bitwiser.ParseFromBits("0001001 1100")  //remove spaces and fix size   --> 0000 1001 1100
	b4, _ := bitwiser.ParseFromBits("11111001 1100") //remove spaces and fix size   --> 1111 1001 1100
	b5, _ := bitwiser.ParseFromBits("1101001 1100")  //remove spaces and fix size   --> 0110 1001 1100
	//
	println(b0.ToString()) //output: 0x1c   (len(array) = 1byte)
	println(b1.ToString()) //output: 0x1c   (len(array) = 1byte)
	println(b2.ToString()) //output: 0x9c   (len(array) = 1byte)
	println(b3.ToString()) //output: 0x009c (len(array) = 2bytes)
	println(b4.ToString()) //output: 0x0f9c (len(array) = 2bytes)
	println(b5.ToString()) //output: 0x069c (len(array) = 2bytes)
}
```


