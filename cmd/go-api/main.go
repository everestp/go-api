package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/everestp/go-api/internal/config"
	"github.com/everestp/go-api/internal/http/handlers/student"
)

func main() {

	//load config
	cfg := config.MustLoad()

	//database setup

	//setup router
 router := http.NewServeMux()
router.HandleFunc("POST /api/students", student.New())


	// setuop server
	 server := http.Server{
		Addr: cfg.HTTPServer.Addr,
		Handler: router,
	}
	slog.Info("Server started %s" ,slog.String("Address", cfg.HTTPServer.Addr))
	fmt.Printf("server started %s",cfg.HTTPServer.Addr)

	done := make(chan os.Signal , 1)

	signal.Notify(done , os.Interrupt , syscall.SIGINT ,syscall.SIGTERM)

	go func ()  {
		
		err := server.ListenAndServe()
		if err != nil {
		   log.Fatal("Faied to starte server")
		}
	}()

<- done 
slog.Info("Shutting down the server")

 ctx ,cancel:= context.WithTimeout(context.Background() ,5 * time.Second )
 defer cancel()
 err := server.Shutdown(ctx)
 if err != nil{
	slog.Error("failed to shutdown server",slog.String("error",err.Error()))
 }
slog.Info("Server shutdown Sucessfully")
}