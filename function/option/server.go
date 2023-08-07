package option

import (
	"crypto/tls"
	"time"
)

// 初版本
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

// NewDefaultServer 默认初始化函数
func NewDefaultServer(addr string, port int) (*Server, error) {
	return &Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  time.Second * 30,
		MaxConns: 100,
		TLS:      nil,
	}, nil
}

// NewTLSServer 带TLS的初始化函数
func NewTLSServer(addr string, port int, tls *tls.Config) (*Server, error) {
	return &Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  time.Second * 30,
		MaxConns: 100,
		TLS:      tls,
	}, nil
}

// NewServerWithTimeout 带超时时间的初始化函数
func NewServerWithTimeout(addr string, port int, timeout time.Duration) (*Server, error) {
	return &Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  timeout,
		MaxConns: 100,
		TLS:      nil,
	}, nil
}

// NewTLSServerWithMaxConnAndTimeout 支持TLS和最大连接数和超时时间的初始化函数
func NewTLSServerWithMaxConnAndTimeout(addr string, port int, maxconns int, timeout time.Duration, tls *tls.Config) (*Server, error) {
	return &Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  timeout,
		MaxConns: maxconns,
		TLS:      tls,
	}, nil
}

// Option版本
type Option func(*Server)

func WithProtocol(protocol string) Option {
	return func(s *Server) {
		s.Protocol = protocol
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithMaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}

func WihtTLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...Option) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  time.Second * 30,
		MaxConns: 100,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}
	//...
	return &srv, nil
}
