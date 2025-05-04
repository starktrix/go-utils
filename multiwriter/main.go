package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

// anthonygg
// https://www.youtube.com/watch?v=g9tv-M-VCpU

type Conn struct {
	io.Writer
}

type Server struct {
	peers map[*Conn]bool
}

func NewConn() *Conn {
	return &Conn{
		Writer: new(bytes.Buffer),
	}
}

func (c *Conn) Write(p []byte) (int, error) {
	fmt.Println("wrting to the underlying connection: ", string(p))
	return c.Writer.Write(p)
}

func NewServer() *Server {
	s := &Server{
		peers: make(map[*Conn]bool),
	}

	for i := 0; i < 10; i++ {
		s.peers[NewConn()] = true
	}

	return s

}

func (s *Server) broadcast(msg []byte) error {
	// for peer := range s.peers {
	// 	if _, err := peer.Write(msg);

	// }
	peers := []io.Writer{}
	for peer := range s.peers {
		peers = append(peers, peer)
	}

	mw := io.MultiWriter(peers...)
	// io.MultiReader()
	_, err := mw.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	
	return nil
}

func main() {
	s := NewServer()
	s.broadcast([]byte("foo"))
}
