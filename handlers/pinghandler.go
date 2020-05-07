package handlers

import "github.com/karldoenitz/Tigo/TigoWeb"

type PingHandler struct {
	TigoWeb.BaseHandler
}

func (p *PingHandler) Get() {
	p.ResponseAsText("pong~")
}
