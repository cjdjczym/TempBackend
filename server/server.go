package server

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
)

const (
	serverAddr = "127.0.0.1:3305"
)

type Server struct {
	engine   *gin.Engine
	listener net.Listener
	exitC    chan struct{}
}

func Init() (*Server, error) {
	s := &Server{exitC: make(chan struct{})}
	s.engine = gin.New()
	l, err := net.Listen("tcp", serverAddr)
	if err != nil {
		println("listen tcp failed, err: " + err.Error())
		return nil, err
	}
	println("listen success | " + l.Addr().String())
	s.listener = l
	api := s.engine.Group("/api")
	api.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json; multipart/form-data; charset=utf-8")
	})
	api.POST("/user/daily", s.PostUserDaily)
	api.GET("/user/stats", s.GetUserStats)
	api.GET("/manager/daily", s.GetManageDaily)
	api.GET("/manager/moon", s.GetManageMoon)
	api.GET("/manager/all", s.GetManageAll)
	return s, nil
}

func (s *Server) Run() {
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			println("close listener failed, err: " + err.Error())
		}
	}(s.listener)

	errC := make(chan error)

	go func(l net.Listener) {
		h := http.NewServeMux()
		h.Handle("/", s.engine)
		hs := &http.Server{Handler: h}
		errC <- hs.Serve(l)
	}(s.listener)

	select {
	case <-s.exitC:
		println("server exit.")
		return
	case err := <-errC:
		print("server failed, err: " + err.Error())
		return
	}
}

func (s *Server) Close() {
	s.exitC <- struct{}{}
	return
}
