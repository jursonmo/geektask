package server

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jursonmo/geektask/week04/blog/internal/service"
)

type MyHttpServer struct {
	server *http.Server
}

func NewHttpServer(as *service.ArticleService) *MyHttpServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/blog/v1/GetArticle", func(rw http.ResponseWriter, r *http.Request) {
		article := as.GetArticle()
		data, _ := json.Marshal(article)
		rw.Write(data)
	})
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	return &MyHttpServer{
		server: server,
	}
}

func (hs *MyHttpServer) Start(ctx context.Context) error {
	return hs.server.ListenAndServe()
}

func (hs *MyHttpServer) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return hs.server.Shutdown(ctx)
}
