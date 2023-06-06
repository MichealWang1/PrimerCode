package server

import (
	"time"

	"github.com/iobrother/ztimer"
	"github.com/panjf2000/gnet/pkg/pool/goroutine"
	"github.com/sirupsen/logrus"
	"github.com/yunzhong/gateway/protocol"
)

const (
	WsUpgrading = 0
	AuthPending = 1
	Authed      = 2
)

type CmdFunc func(c *Connection, p *protocol.Packet) (err error)

type Server struct {
	opts      Options
	tcpServer *TcpServer
	// wsServer  *WsServer
	timer       *ztimer.Timer
	connManager *ConnManager
	workerPool  *goroutine.Pool
	mapCmdFunc  map[protocol.MSG_ID]CmdFunc
}

func NewServer(opts ...Option) *Server {
	s := new(Server)
	s.opts = NewOptions(opts...)
	s.connManager = NewConnManager()
	s.workerPool = goroutine.Default()

	if s.opts.TcpAddr != "" {
		s.tcpServer = NewTcpServer(s, s.opts.TcpAddr)
	}

	// if s.opts.WsAddr != "" {
	// 	s.wsServer = NewWsServer(s, s.opts.WsAddr)
	// }

	s.timer = ztimer.NewTimer(100*time.Millisecond, 20)

	s.registerCmdFunc()

	return s
}

func (s *Server) registerCmdFunc() {
	s.mapCmdFunc = make(map[protocol.MSG_ID]CmdFunc)
}

func (s *Server) GetConnManager() *ConnManager {
	return s.connManager
}

func (s *Server) GetServerId() string {
	return s.opts.Id
}

func (s *Server) GetTcpServer() *TcpServer {
	return s.tcpServer
}

// func (s *Server) GetWsServer() *WsServer {
// 	return s.wsServer
// }

func (s *Server) GetTimer() *ztimer.Timer {
	return s.timer
}

func (s *Server) consumeKick() error {
	return nil
}

func (s *Server) consumePush() error {

	return nil
}

func (s *Server) Start() error {
	go func() {
		if err := s.consumePush(); err != nil {
			logrus.Error(err)
		}
	}()
	go func() {
		s.timer.Start()
	}()
	go func() {
		if s.tcpServer != nil {
			if err := s.tcpServer.Start(); err != nil {
				logrus.Error(err)
			}
		}
	}()
	// go func() {
	// 	if s.wsServer != nil {
	// 		if err := s.wsServer.Start(); err != nil {
	// 			logrus.Error(err)
	// 		}
	// 	}
	// }()

	return nil
}

func (s *Server) Stop() error {
	var lastError error
	if s.tcpServer != nil {
		if err := s.tcpServer.Stop(); err != nil {
			lastError = err
		}
	}
	// if s.wsServer != nil {
	// 	if err := s.wsServer.Stop(); err != nil {
	// 		lastError = err
	// 	}
	// }
	return lastError
}

func (s *Server) OnOpen(c *Connection) {
	// 10秒钟之内没有认证成功，关闭连接
	c.TimerTask = s.GetTimer().AfterFunc(time.Second*10, func() {
		logrus.Info("auth timeout...")
		c.Close()
	})
}

func (s *Server) OnClose(c *Connection) {
	logrus.Infof("client = %s close", c.Uin)

	if c.ID == "" {
		return
	}

	_ = s.workerPool.Submit(func() {
		if c != nil {
			logrus.Infof("workerPool.Submit = %s close", c.Uin)
		}
	})

	s.GetConnManager().Remove(c)
}

func (s *Server) OnMessage(data []byte, c *Connection) {
	_ = s.workerPool.Submit(func() {
		p := &protocol.Packet{}
		if err := p.Read(data); err != nil {
			logrus.Error(err)
			c.Close()
			return
		}
		// if c.Status == AuthPending {
		// 	cmd := protocol.MSG_ID(p.Cmd)
		// 	if cmd != protocol.CmdId_Cmd_Login {
		// 		logrus.Error("first packet must be cmd_login")
		// 		c.Close()
		// 		return
		// 	}
		// 	if err := s.handleLogin(c, p); err != nil {
		// 		c.Close()
		// 		logrus.Info("AUTH FAILED")
		// 	} else {
		// 		c.Status = Authed
		// 	}
		// } else {
		s.handleProto(c, p)
		// }
	})

}

func (s *Server) handleLogin(c *Connection, p *protocol.Packet) (err error) {

	return
}

func (s *Server) handleLogout(c *Connection, p *protocol.Packet) (err error) {

	return
}

func (s *Server) handleProto(c *Connection, p *protocol.Packet) (err error) {
	logrus.Infof("cmd=%d", p.Cmd)
	cmd := protocol.MSG_ID(p.Cmd)

	if s.mapCmdFunc[cmd] != nil {
		err = s.mapCmdFunc[cmd](c, p)
	}

	return
}

func (s *Server) handleNoop(c *Connection, p *protocol.Packet) (err error) {
	return
}
