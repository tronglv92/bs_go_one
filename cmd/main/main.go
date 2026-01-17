package main

import (
	"flag"
	"fmt"
	"go_one/helper/server"
	"go_one/internal/config"
	"go_one/internal/handler"
	"go_one/internal/registry"

	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/app.yaml", "the config file")

func main() {
	c := config.Load(configFile)

	svcGroup := service.NewServiceGroup()
	svcGroup.Add(server.NewHttpServer(c.Server,
		handler.NewRestHandler(registry.NewServiceContext(c)),
	))
	defer svcGroup.Stop()
	fmt.Printf("Starting server at %s:%d...\n", c.Server.Http.Host, c.Server.Http.Port)
	svcGroup.Start()
}
