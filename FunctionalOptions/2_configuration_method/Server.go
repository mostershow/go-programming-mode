package __configuration_method

import (
	"crypto/tls"
	"fmt"
	"time"
)

//使用一个配置对象,把那些非必输的选项都移到一个结构体里
type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

// 简化后的server
type Server struct {
	Addr string
	Port int
	Conf *Config
}

func NewServer(addr string, port int, conf *Config) (*Server, error) {
	//...
	return nil, nil
}

func main() {
	//Using the default configuratrion
	srv1, _ := NewServer("localhost", 9000, nil)
	conf := &Config{}
	// 在使用前需要构造 Config 对象 上省略详细构造
	srv2, _ := NewServer("locahost", 9000, conf)

	fmt.Println(srv1, srv2)
}
