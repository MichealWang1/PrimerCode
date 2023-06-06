package driver

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/yunzhong/gateway/driver/server"
	"github.com/zmicro-team/zmicro/core/config"
	"github.com/zmicro-team/zmicro/core/util/env"
)

type BeforeFunc func() error

type SelectOptions struct {
	Before BeforeFunc
}

func newSelectOptions(opts ...SelOption) SelectOptions {
	options := SelectOptions{}

	for _, o := range opts {
		o(&options)
	}

	return options
}

type SelOption func(*SelectOptions)

func Before(f BeforeFunc) SelOption {
	return func(o *SelectOptions) {
		o.Before = f
	}
}

// 读取配置文件
var cfgFile string

// 配置文件内容
type zconfig struct {
	App struct {
		Mode string
		Name string
	}
	Server struct {
		TcpAddr string
		WsAddr  string
	}
}

func init() {
	flag.StringVar(&cfgFile, "config", "config.yaml", "config file")
}

type Driver struct {
	opts   SelectOptions
	zc     *zconfig
	server *server.Server
}

func NewDriver(sop SelOption) *Driver {
	seloptions := newSelectOptions(sop)
	flag.Parse()
	_, err := os.Stat(cfgFile)
	if os.IsNotExist(err) {
		logrus.Fatal("没有配置文件")
	}
	c := config.New(config.Path(cfgFile))
	config.ResetDefault(c)
	zc := &zconfig{}
	if err = config.Unmarshal(zc); err != nil {
		logrus.Fatal(err.Error())
	}
	if zc.App.Name == "" {
		logrus.Fatal("配置项app.name不能为空")
	}
	env.Set(zc.App.Mode)

	driver := &Driver{
		opts: seloptions,
		zc:   zc,
	}
	driver.server = server.NewServer(
		server.TcpAddr(zc.Server.TcpAddr),
	)
	return driver
}

func (driver *Driver) Run() error {
	if driver.opts.Before != nil {
		if err := driver.opts.Before(); err != nil {
			return err
		}
	}
	if err := driver.server.Start(); err != nil {
		return err
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)
	logrus.Infof("启动读取信息 received signal %s", <-ch)
	err := driver.server.Stop()
	return err
}
