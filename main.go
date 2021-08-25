package main

import (
	"./server"
	"os"
	"os/signal"
	"sync"
	"syscall"
)


func main() {
	srv, err := server.Init()
	if err != nil {
		println("init failed, err: " + err.Error())
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGPIPE, syscall.SIGUSR1)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			sig := <-c
			if sig == syscall.SIGINT || sig == syscall.SIGTERM || sig == syscall.SIGQUIT {
				println("got signal, quit, signal: " + sig.String())
				srv.Close()
				return
			}
			println("ignore signal: " + sig.String())
		}
	}()
	srv.Run()
	wg.Wait()
}
