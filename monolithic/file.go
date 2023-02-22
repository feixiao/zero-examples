package main

import (
	"context"
	"flag"
	"fmt"
	"monolithic/internal/logic"
	"net/http"

	"monolithic/internal/config"
	"monolithic/internal/handler"
	"monolithic/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/file-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	upload := logic.NewUploadLogic(context.Background(), ctx)

	server.AddRoutes(
		[]rest.Route{
			{
				Method: http.MethodPost,
				Path:   "/upload",
				Handler: func(w http.ResponseWriter, r *http.Request) {
					upload.Upload(ctx, w, r)
				},
			},
		},
	)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
