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
[https://stackoverflow.com/questions/28997600/golang-bitwise-operation-on-very-long-binary-bit-string-representation](Stackoverflow):
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


