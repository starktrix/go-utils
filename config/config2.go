package main


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

// 2b  default config
func NewServer2() *Servers {
	return &Servers{
		config: NewConfig(),
	}
}

// specified config
func NewServer2WithConfig(config Config) *Servers {
	return &Servers{
		config: config,
	}
}