package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo/db"
	"todo/routes"
)


func main(){
	ch:=make(chan os.Signal,1)
	signal.Notify(ch,os.Interrupt,syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	connString:=os.Getenv("MONGO_URL")

	slog.Info("Conecting to the db")

	_,err:=db.Init(connString,"ToDo")

	if err!=nil{
		fmt.Println(err)
	}

	slog.Info("Creating routes!!")

	r:=routes.Routes()

	server := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	go func(){
		slog.Info("Starting HTTP server on port 3000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed:", err)
		}
	}()

	sig:=<-ch
	slog.Info("Received signal:", sig)

	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 5*time.Second)
	defer shutdownCancel()

	select{
	case <-shutdownCtx.Done():
		slog.Info("Shutdown completed successfully")
	}
}