package main

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kahuri1/birthday-notifier/internal/server"
	srv "github.com/kahuri1/birthday-notifier/pkg/server"
)

func main() {
	s := server.NewServer()

	r := gin.Default()

	srv.RegisterHandlersWithOptions(r, s, srv.GinServerOptions{
		BaseURL: "/api/web/v1",
	})

	httpsrv := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", "8080"),
	}

	if err := httpsrv.ListenAndServe(); err != nil {
		panic(err)
	}
}
