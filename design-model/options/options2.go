package options

import "time"

// 真正的optional 实现

type Option func(server *Server)

func TimeOut(timeout string) Option {
	return func(server *Server) {
		server.Conf.Timeout = 100 * time.Second
	}
}

func MaxCon(maxcon int) Option {
	return func(server *Server) {
		server.Conf.Maxconns = maxcon
	}
}

func New(option []Option) (server *Server) {

	for _, function := range option {
		function(server)
	}

	return
}
