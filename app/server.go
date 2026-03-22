package app

import (
	"dayz-server-tools/logger"
	"dayz-server-tools/server"
	"sync"
)

var (
	_server *Server
	once    sync.Once
)

type Server struct {
}

func GetServer() *Server {
	once.Do(func() {
		_server = &Server{}
	})
	return _server
}

func (s *Server) TestServerStart() {
	logger.Info("测试服务器启动")
	server.StartServer()
}
