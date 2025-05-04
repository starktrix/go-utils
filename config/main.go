package main

import "fmt"

func main() {
	// s0 := newServer(Opts{})
	s := newServerConfig1()
	s2 := newServerConfig1(withTLS, withMaxConn(100))
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", s2)

	config := NewConfig().
		withTLS(true).
		withMaxConn(50)

	srv := NewServer(config)
	fmt.Printf("%+v\n", srv)

}
