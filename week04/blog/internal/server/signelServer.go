package server

//kratos 本身就可以侦听信号退出，为了练手，这里自己写个侦听信号退出的服务
import (
	"context"
	"os"
	"os/signal"

	//"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

type SignalServer struct {
	logger log.Logger
	ch     chan os.Signal
}

func NewSignalServer(logger log.Logger) *SignalServer {
	return &SignalServer{
		logger: logger,
		ch:     make(chan os.Signal, 1),
	}
}

func (ss *SignalServer) Start(ctx context.Context) error {
	log := log.NewHelper(ss.logger)
	log.Info("my SignalServer start")
	//signal.NotifyContext(ctx, os.Interrupt, os.Kill)//go 1.16
	signal.Notify(ss.ch, os.Interrupt, os.Kill)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case sig := <-ss.ch:
		return errors.Errorf("receive quit signal:%v", sig)
	}
	return errors.New("my SignalServer quit")
}

func (ss *SignalServer) Stop(ctx context.Context) error {
	log := log.NewHelper(ss.logger)
	log.Infof("my SignalServer stop, ctx.Err:%v", ctx.Err())
	return nil
}
