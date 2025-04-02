package main

import "fmt"

// 1. 
// type Server struct {
// 	maxConn		int
// 	id			string
// 	tls			bool
// }

// func newServer(maxConn int, id string, tls bool) *Server {}

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func withTLS(opts *Opts) {
	opts.tls = true
}
func withMaxConn(n int) OptFunc {
	return func (opts *Opts) {
	opts.maxConn = n
}
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 5,
		id: "Default",
		tls: false,
	}
}

type Server struct {
	Opts
}

// func newServer(opts Opts) *Server {
// 	return &Server{
// 		Opts: opts,

// 	}
// }
func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,

	}
}

// 2.
type Servers struct {
	config		Config
}

type Config struct {
	maxConn		int
	id			string
	tls			bool
}

func (c Config) withMaxConn(n int) Config {
	c.maxConn = n
	return c
}
func (c Config) withTLS(status bool) Config {
	c.tls = status
	return c
}

func NewConfig() Config {
	return Config{
		maxConn: 5,
		id: "Dev",
		tls: false,
	}
}

func NewServer(config Config) *Servers {
	return &Servers{
		config: config,
	}
}

// 2b
func NewServer2() *Servers {
	return &Servers{
		config: NewConfig(),
	}
}
func NewServer2WithConfig(config Config) *Servers {
	return &Servers{
		config: config,
	}
}


func main() {
	// s := newServer(Opts{})
	s := newServer()
	s2 := newServer(withTLS, withMaxConn(100))
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", s2)

	config := NewConfig().
	withTLS(true).
	withMaxConn(50)

	srv := NewServer(config)
	fmt.Printf("%+v\n", srv)

	di()
	di2()
	di3()

}
