package main

import (
	"context"
	"fmt"
	"github.com/astlaure/httpmirror/internal/app"
	"github.com/astlaure/httpmirror/internal/core"
	"github.com/astlaure/httpmirror/internal/proxy"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	core.ReadConfigFile()
	core.Connect()

	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	appServer := app.StartApp()
	proxyServer := proxy.CreateProxy()

	fmt.Println("Starting App Server")
	go func() {
		appServer.ListenAndServe()
	}()

	fmt.Println("Starting Proxy Server")
	go func() {
		proxyServer.ListenAndServe()
	}()

	defer func() {
		if err := appServer.Shutdown(ctx); err != nil {
			fmt.Println("error when shutting down the main server: ", err)
		}
		if err := proxyServer.Shutdown(ctx); err != nil {
			fmt.Println("error when shutting down the admin server: ", err)
		}
	}()

	sig := <-sigs
	fmt.Println(sig)

	cancel()

	fmt.Println("service has shutdown")
}
