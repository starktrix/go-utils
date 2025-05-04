package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

type FilesizeWriter struct {
	written int
}

const MAX_FILE_SIZE = 513

func (w *FilesizeWriter) Write(b []byte) (int, error) {
	log.Println("Written before: ", w.written)
	w.written += len(b)
	log.Println("Written  after: ", w.written)
	if w.written > MAX_FILE_SIZE {
		log.Println("Limit exceeded: ", w.written)
		return 0, fmt.Errorf("max filesize (%d) exceeded", MAX_FILE_SIZE)
	}
	return len(b), nil
}

func main() {
	src := new(bytes.Buffer)
	for i := range 100 {
		content := fmt.Sprintf("Hello World: (%d)\n", i)
		src.Write([]byte(content))
	}
	dst := new(bytes.Buffer)
	fw := &FilesizeWriter{} // ensure its a pointer otherwise, w.writen is overriden each call
	// io.Copy(dstWriter, srcReader)
	// io.Copy() simpy copies from src to dst with no control
	// using tee reader, it returns a reader `tee` that reads
	// from src to fw (custom writer with control)
	tee := io.TeeReader(src, fw)
	// this then copies from tee(src->fw) -> dst
	n, err := io.Copy(dst, tee) //copy seems to be 512 bytes
	fmt.Printf("Size copied from tee to dst: (%d)\n", n)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println((dst.String()))
}

