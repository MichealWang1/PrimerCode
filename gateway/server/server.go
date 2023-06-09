package server

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/iobrother/ztimer"
	"github.com/panjf2000/gnet/pool/goroutine"
	"github.com/yunzhong/gateway/protocol"
	"github.com/zmicro-team/zmicro/core/log"
)

const (
	WsUpgrading = 0
	AuthPending = 1
	Authed      = 2
)

type CmdFunc func(c *Connection, p *protocol.Packet) (err error)

type Server struct {
	opts        Options
	tcpServer   *TcpServer
	wsServer    *WsServer
	timer       *ztimer.Timer
	connManager *ConnManager
	workerPool  *goroutine.Pool
	mapCmdFunc  map[protocol.CmdId]CmdFunc
}

func NewServer(opts ...Option) *Server {
	s := new(Server)
	s.opts = NewOptions(opts...)
	s.connManager = NewConnManager()
	s.workerPool = goroutine.Default()
	if s.opts.TcpAddr != "" {
		s.tcpServer = NewTcpServer(s, s.opts.TcpAddr)
	}
	if s.opts.WsAddr != "" {
		s.wsServer = NewWsServer(s, s.opts.WsAddr)
	}
	s.timer = ztimer.NewTimer(100*time.Millisecond, 20)
	s.registerCmdFunc()
	return s
}

func (s *Server) registerCmdFunc() {
	s.mapCmdFunc = make(map[protocol.CmdId]CmdFunc)
	s.mapCmdFunc[protocol.CmdId_Cmd_Head] = s.handleNoop
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

func (s *Server) GetWsServer() *WsServer {
	return s.wsServer
}

func (s *Server) GetTimer() *ztimer.Timer {
	return s.timer
}

func (s *Server) Start() error {
	go func() {
		if err := s.consumePush(); err != nil {
			log.Error(err)
		}
	}()
	go func() {
		s.timer.Start()
	}()
	go func() {
		if s.tcpServer != nil {
			if err := s.tcpServer.Start(); err != nil {
				log.Error(err)
			}
		}
	}()
	go func() {
		if s.wsServer != nil {
			if err := s.wsServer.Start(); err != nil {
				log.Error(err)
			}
		}
	}()
	return nil
}

func (s *Server) Stop() error {
	var lastError error
	if s.tcpServer != nil {
		if err := s.tcpServer.Stop(); err != nil {
			lastError = err
		}
	}
	if s.wsServer != nil {
		if err := s.wsServer.Stop(); err != nil {
			lastError = err
		}
	}
	return lastError
}

func (s *Server) consumeKick() error {
	return nil
}

func (s *Server) consumePush() error {
	return nil
}

func (s *Server) OnOpen(c *Connection) {
	// 10秒钟之内没有认证成功，关闭连接
	// c.TimerTask = s.GetTimer().AfterFunc(time.Second*10, func() {
	// 	log.Info("auth timeout...")
	// 	c.Close()
	// })
}

func (s *Server) OnClose(c *Connection) {
	log.Infof("client=%s close", c.ID)
	if c.ID == "" {
		return
	}
	_ = s.workerPool.Submit(func() {
		if c != nil {

		}
	})
	s.GetConnManager().Remove(c)
}

func (s *Server) OnMessage(data []byte, c *Connection) {
	_ = s.workerPool.Submit(func() {
		p := &protocol.Packet{}
		if err := p.Read(data); err != nil {
			log.Error(err)
			c.Close()
			return
		}
		s.handleProto(c, p)
	})

}

func (s *Server) handleLogin(c *Connection, p *protocol.Packet) (err error) {
	fmt.Println(" 发来登录消息 ")
	// p.Body = []byte("hello")
	// c.WritePacket(p)
	s.responseError(c, p, 110)
	return
}

func (s *Server) handleProto(c *Connection, p *protocol.Packet) (err error) {
	log.Infof("cmd = %d", p.Cmd)
	cmd := protocol.CmdId(p.Cmd)
	if cmd > 0 {
		err = s.handleLogin(c, p)
	}

	return
}

func (s *Server) handleNoop(c *Connection, p *protocol.Packet) (err error) {
	log.Infof("client %s noop", c.ID)
	c.WritePacket(p)
	return
}

func (s *Server) responseError(c *Connection, p *protocol.Packet, err int32) {
	fmt.Println(" 发生错误码 ")
	rsp := &protocol.Error{}
	rsp.Code = err
	rsp.Message = "111111"
	b, _ := proto.Marshal(rsp)
	p.BodyLen = uint32(len(b))
	p.Body = b
	_ = c.WritePacket(p)
}

func (s *Server) responseMessage(c *Connection, p *protocol.Packet, m proto.Message) {
	b, _ := proto.Marshal(m)
	p.BodyLen = uint32(len(b))
	p.Body = b
	_ = c.WritePacket(p)
}
