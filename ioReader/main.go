package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	p := []byte("Hello World")
	payload := bytes.NewReader(p)

	b, err := io.ReadAll(payload)
	if err != nil {
		fmt.Println("Error: ", err)
		return 
	}

	hash := sha1.Sum(b)

	fmt.Println("Hex Encoded: ", hex.EncodeToString(hash[:]))


	c, err := io.ReadAll(payload)
	if err != nil {
		fmt.Println("Error: ", err)
		return 
	}
	// how to compose interface in golang
	// no output. byte data has been consumed
	fmt.Println("String: ", string(c))

	hr := NewHashReader(p)
	soln(hr)

}


type HashReader interface {
	io.Reader
	// hash string

}

type hasReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hasReader {
	return &hasReader{
		Reader: bytes.NewReader(b),
		buf: bytes.NewBuffer(b),

	}
}


func (h *hasReader) hash() string {
	hash := sha1.Sum(h.buf.Bytes())
	return  hex.EncodeToString(hash[:])
}

func soln(r io.Reader) {
	hash := r.(*hasReader).hash()
	fmt.Println("Hashed: ", hash)

	c, err := io.ReadAll(r.(*hasReader).Reader)
	if err != nil {
		fmt.Println("Error (soln): ", err)
		return 
	}
	// how to compose interface in golang
	// no output. byte data has been consumed
	fmt.Println("String (soln): ", string(c))

}