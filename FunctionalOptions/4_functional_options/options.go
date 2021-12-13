package __functional_options

import (
	"crypto/tls"
	"time"
)

// 先定义一个函数类型
type Option func(*Server)

// 传入一个参数，然后返回一个函数，返回的这个函数会设置自己的 Server 参数
func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}
func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}
func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}
