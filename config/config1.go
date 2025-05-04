package main

// 1. traditional approach
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

type Server struct {
	Opts
}

// 1a
func newServer(opts Opts) *Server {
	return &Server{
		Opts: opts,
	}
}

// 1b
func withTLS(opts *Opts) {
	opts.tls = true
}

func withMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 5,
		id:      "Default",
		tls:     false,
	}
}

func newServerConfig1(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}
