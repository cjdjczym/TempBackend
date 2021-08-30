package main

import (
	"TempBackend/model"
	"TempBackend/server"
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
)
// @title TemperatureBackend API
// @version 1.0
// @description author github@cjdjczym.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 47.100.236.6:3305
// @BasePath "/api"
func main() {
	configFilePath := flag.String("config", "etc/config.yaml", "temperature backend config file path")
	flag.Parse()
	configData, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		println("read config file failed, err: " + err.Error())
		os.Exit(1)
	}

	cfg, err := model.UnmarshalConfig(configData)
	if err != nil {
		println("parse config file failed, err: " + err.Error())
		os.Exit(1)
	}

	srv, err := server.Init(cfg)
	if err != nil {
		println("init failed, err: " + err.Error())
		os.Exit(1)
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
				break
			}
			println("ignore signal: " + sig.String())
		}
	}()

	srv.Run()
	wg.Wait()
}
