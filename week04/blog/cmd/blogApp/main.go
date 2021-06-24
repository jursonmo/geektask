package main

import (
	"os"

	kratos "github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jursonmo/geektask/week04/blog/internal/server"
)

var Name = "blogApp"
var Version = "v1.0.0"

func newApp(logger log.Logger, hs *server.MyHttpServer, ss *server.SignalServer) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Signal(os.Interrupt, os.Kill),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			ss,
		),
	)
}

func main() {
	//logger := NewZapLogger()
	logger := log.NewStdLogger(os.Stdout)
	//log := log.NewHelper(logger)
	app, err := initApp(logger)
	if err != nil {
		panic(err)
	}
	err = app.Run()
	if err != nil {
		panic(err)
	}
}
