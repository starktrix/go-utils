package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type FileServer struct{}

func (fs *FileServer) start() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go fs.readLoop(conn)
	}
}

// func (fs *FileServer) readLoop(conn net.Conn) {
// 	buf := make([]byte, 2048)
// 	for {
// 		n, err := conn.Read(buf)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		file := buf[:n]
// 		fmt.Println(file)
// 		fmt.Printf("recevied %d bytes over the network\n", n)
// 	}
// }
func (fs *FileServer) readLoop(conn net.Conn) {
	buf := new(bytes.Buffer)
	for {
		// n, err := io.Copy(buf, conn)
		var size int64
		binary.Read(conn, binary.LittleEndian, &size)
		n, err := io.CopyN(buf, conn, size) //file size is not known
		if err != nil {
			log.Fatal(err)
		}

		// panic("should panic!!!") //never occurs when using io.Copy
		fmt.Println("#############################################################################################")
		fmt.Println("#############################################################################################")
		fmt.Println("#############################################################################################")
		fmt.Println(buf.Bytes())
		fmt.Println("#############################################################################################")
		fmt.Println("#############################################################################################")
		fmt.Println("#############################################################################################")
		fmt.Printf("recevied %d bytes over the network\n", n)
	}
}

func main() {
	go func() {
		time.Sleep(4 * time.Second)
		sendFile(2000)
		time.Sleep(2 * time.Second)
		sendFile(2000)
		time.Sleep(2 * time.Second)
		sendFile(2000)
		time.Sleep(2 * time.Second)
		sendFile(2000)
		time.Sleep(2 * time.Second)
		sendFile(2000)
	}()
	server := &FileServer{}
	server.start()
}

// func sendFile(size int) error {
// 	file := make([]byte, size)
// 	_, err := io.ReadFull(rand.Reader, file)
// 	if err != nil {
// 		return err
// 	}

// 	conn, err := net.Dial("tcp", ":8000")
// 	if err != nil {
// 		return err
// 	}

//		n, err := conn.Write(file)
//		if err != nil {
//			return err
//		}
//		fmt.Printf("writtend %d bytes over the network\n", n)
//		return nil
//	}
func sendFile(size int) error {
	file := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		return err
	}

	// n, err := io.Copy(conn, bytes.NewReader(file))
	binary.Write(conn, binary.LittleEndian, int64(size)) //send file metadata to server
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))

	if err != nil {
		return err
	}
	fmt.Printf("written %d bytes over the network\n", n)
	return nil
}
