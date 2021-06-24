package server

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	v1 "github.com/jursonmo/geektask/week04/blog/api/v1"
	"github.com/jursonmo/geektask/week04/blog/internal/service"
)

type MyHttpServer struct {
	server *http.Server
}

func NewHttpServer(as *service.ArticleService) *MyHttpServer {
	//mux := http.NewServeMux()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/article/{id}", func(rw http.ResponseWriter, r *http.Request) {
		ar := v1.ArticleReq{}
		vars := mux.Vars(r)
		if id, ok := vars["id"]; !ok {
			ar.Id, _ = strconv.Atoi(id)
		}
		article := as.GetArticle(ar.Id)
		data, _ := json.Marshal(article)
		rw.Write(data)
	})
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
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
