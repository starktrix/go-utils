package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWSOrderbook(ws *websocket.Conn) {
	fmt.Println("new incoming ws connection from client to orderbook feed: ", ws.RemoteAddr)


	for {
		payload := fmt.Sprintf("orderbook data -> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 2)
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming ws connection")

	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		_, err := ws.Read(buf)
		if err != nil {
			if err == io.ErrClosedPipe {
				break
			}

			fmt.Println("Error: ", err)
			continue //continue to prevent dropping connection
		}
		msg := buf[:]
		fmt.Println(string(msg))
		// ws.Write([]byte("Thank you!!!!"))

		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error: ", err)
			}
		}(ws)
	}
}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/orderbook", websocket.Handler(server.handleWSOrderbook))
	http.ListenAndServe(":8080", nil)
}
