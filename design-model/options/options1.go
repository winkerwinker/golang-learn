package options

//使用一个builder类来做包装
type ServerBuilder struct {
	Server
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {

	sb.Server.Addr = addr
	sb.Server.Port = port
	return sb
}
func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	sb.Server.Conf.Protocol = protocol
	return sb
}

//builder
//
//type Config struct {
//	Protocol string
//	Timeout  time.Duration
//	Maxconns int
//	TLS      *tls.Config
//}
//
//type Server struct {
//	Addr string
//	Port int
//	Conf *Config
//}
func (sb *ServerBuilder) Build() (Server) {
	return  sb.Server
}

