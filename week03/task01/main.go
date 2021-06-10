package main

/*
基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
*/
import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello world"))
	})
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	eg.Go(func() error {
		err := server.ListenAndServe()
		return fmt.Errorf("server.ListenAndServe err:%w", err) //%w, wrap err with msg
	})

	//shutdown http server
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel() //stop timer
			err := server.Shutdown(timeoutCtx)
			fmt.Println("server.Shutdown err:", err)
			return ctx.Err()
		}
	})

	//listen signal
	eg.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, os.Kill)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case s := <-quit:
			fmt.Printf("get signal:%v and quit\n", s)
			return fmt.Errorf("get signal:%v and quit", s)
		}
	})

	//wait all task quit
	fmt.Println(eg.Wait().Error())
}
