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
	mux    *mux.Router
}

func (hs *MyHttpServer) RegisterApi(as *service.ArticleService) {
	r := hs.mux
	//register middleware
	//r.Use()
	r.HandleFunc("/api/v1/article/{id}", func(rw http.ResponseWriter, req *http.Request) {
		ar := v1.ArticleReq{}
		vars := mux.Vars(req)
		if id, ok := vars["id"]; !ok {
			i, _ := strconv.Atoi(id)
			ar.Id = int64(i)
		}
		article := as.GetArticle(ar.Id)
		data, _ := json.Marshal(article)
		rw.Write(data)
	}).Methods("GET")
	r.HandleFunc("/api/v1/article/", func(rw http.ResponseWriter, req *http.Request) {
		acr := &v1.ArticleCreateReq{}
		if req.Header.Get("Content-Type") != "application/json" {
			return
		}
		buf := make([]byte, req.ContentLength)
		req.Body.Read(buf)
		err := json.Unmarshal(buf, &acr)
		if err != nil {
			return
		}
		article := as.CreateArticle(acr)
		data, _ := json.Marshal(article)
		rw.Write(data)
	}).Methods("POST")
}

func NewHttpServer(as *service.ArticleService) *MyHttpServer {
	//mux := http.NewServeMux()
	r := mux.NewRouter()
	/*
		r.HandleFunc("/api/v1/article/{id}", func(rw http.ResponseWriter, r *http.Request) {
			ar := v1.ArticleReq{}
			vars := mux.Vars(r)
			if id, ok := vars["id"]; !ok {
				i, _ := strconv.Atoi(id)
				ar.Id = int64(i)
			}
			article := as.GetArticle(ar.Id)
			data, _ := json.Marshal(article)
			rw.Write(data)
		})
	*/
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}
	hs := &MyHttpServer{
		server: server,
		mux:    r,
	}
	hs.RegisterApi(as)
	return hs
}

func (hs *MyHttpServer) Start(ctx context.Context) error {
	return hs.server.ListenAndServe()
}

func (hs *MyHttpServer) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return hs.server.Shutdown(ctx)
}
