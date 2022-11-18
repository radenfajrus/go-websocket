package main

import (
	"app/server"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	httpConfig := &server.HttpConfig{
		Port: "3000",
	}

	server.Init(httpConfig)

	http := server.HttpServer
	wss := server.WebsocketServer

	go func() {
		http.Start()
	}()
	go func() {
		wss.Start()
	}()

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt)
	signal.Notify(shutdownChan, os.Kill)

	sig := <-shutdownChan
	fmt.Println("Shutting down : ", sig)
	wss.Shutdown()
	http.Shutdown()
}
