package server

import (
	"bytes"
	"sync"

	"github.com/google/uuid"
	"github.com/iobrother/ztimer"
	"github.com/panjf2000/gnet"
	"github.com/yunzhong/gateway/protocol"
)

type Connection struct {
	ID        string
	Status    int
	TimerTask *ztimer.TimerTask
	DeviceId  string
	Conn      gnet.Conn
	Version   int
	Uin       string
	Platform  string
	Server    string
}

func (c *Connection) Write(data []byte) error {
	return c.Conn.AsyncWrite(data)
}

func (c *Connection) WritePacket(p *protocol.Packet) error {
	buf := &bytes.Buffer{}
	if err := p.Write(buf); err != nil {
		return err
	}
	return c.Conn.AsyncWrite(buf.Bytes())
}

func (c *Connection) Close() {
	if c.Conn != nil {
		c.Conn.Close()
	}
}

type Options struct {
	Id      string
	TcpAddr string
	// WsAddr  string
}

func NewOptions(opts ...Option) Options {
	options := Options{
		Id: uuid.New().String(),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

type Option func(*Options)

func TcpAddr(addr string) Option {
	return func(o *Options) {
		o.TcpAddr = addr
	}
}

// func WsAddr(addr string) Option {
// 	return func(o *Options) {
// 		o.WsAddr = addr
// 	}
// }

type ConnManager struct {
	sync.Mutex
	conns map[string]*Connection
}

func NewConnManager() *ConnManager {
	cm := new(ConnManager)
	cm.conns = make(map[string]*Connection)
	return cm
}

func (cm *ConnManager) Add(c *Connection) {
	cm.Lock()
	cm.conns[c.ID] = c
	cm.Unlock()
}

func (cm *ConnManager) Get(id string) *Connection {
	cm.Lock()
	defer cm.Unlock()
	c := cm.conns[id]
	return c
}

func (cm *ConnManager) Remove(c *Connection) {
	cm.Lock()
	defer cm.Unlock()
	delete(cm.conns, c.ID)
	return
}
