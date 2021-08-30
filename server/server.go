package server

import (
	"TempBackend/metrics"
	"TempBackend/model"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net"
	"net/http"
)

type Server struct {
	engine   *gin.Engine
	listener net.Listener
	cfg      *model.Config
	exitC    chan struct{}
}

func Init(cfg *model.Config) (*Server, error) {
	metrics.RegisterProxyMetrics()

	s := &Server{cfg: cfg, exitC: make(chan struct{})}
	s.engine = gin.New()
	l, err := net.Listen("tcp", cfg.ServerAddr)
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

	api.GET("/metrics", gin.WrapF(promhttp.Handler().ServeHTTP))

	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
