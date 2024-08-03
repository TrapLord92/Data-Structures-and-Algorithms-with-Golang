// main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing fmt package
import (
	"fmt"
	//	"strconv"
	"crypto/sha1"
	"hash"
)

/*Hashing
Hash functions were introduced in Chapter 4, Non-Linear Data Structures. Hash
implementation in Go has crc32 and sha256 implementations. An implementation of a
hashing algorithm with multiple values using an XOR transformation is shown in the
following code snippet. */

/*The CreateHash function takes a byte array, byteStr, as a
parameter and returns the sha256 checksum of the byte array:*/
//CreateHash method
func CreateHash(byteStr []byte) []byte {
	var hashVal hash.Hash
	hashVal = sha1.New()
	hashVal.Write(byteStr)

	var bytes []byte

	bytes = hashVal.Sum(nil)
	return bytes
}

/*The CreateHashMutliple method
The CreateHashMutliple method takes the byteStr1 and byteStr2 byte arrays as
parameters and returns the XOR-transformed bytes value, as follows:*/
// Create hash for Multiple Values method
func CreateHashMultiple(byteStr1 []byte, byteStr2 []byte) []byte {
	return xor(CreateHash(byteStr1), CreateHash(byteStr2))
}

// XOR method
/*The XOR method
The xor method takes the byteStr1 and byteStr2 byte arrays as parameters and returns
the XOR-transformation result, as follows:*/

func xor(byteStr1 []byte, byteStr2 []byte) []byte {
	var xorbytes []byte
	xorbytes = make([]byte, len(byteStr1))
	var i int
	for i = 0; i < len(byteStr1); i++ {
		xorbytes[i] = byteStr1[i] ^ byteStr2[i]
	}
	return xorbytes
}

// main method
/*The main method
The main method invokes the createHashMutliple method, passing Check and
Hash as string parameters, and prints the hash value of the strings, as follows:*/
func main() {

	var bytes []byte
	bytes = CreateHashMultiple([]byte("Check"), []byte("Hash"))

	fmt.Printf("%x\n", bytes)
}
