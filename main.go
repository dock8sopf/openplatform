package main

import (
	"github.com/karldoenitz/Tigo/TigoWeb"
	"opengateway/handlers"
)

var urlRouter = []TigoWeb.Router{
	{"/ping", &handlers.PingHandler{}, nil},
	{"/", &handlers.EngineHandler{}, nil},
}

func main() {
	application := TigoWeb.Application{
		UrlRouters: urlRouter,
		ConfigPath: "./serverConfig.yaml",
	}
	application.Run()
}
