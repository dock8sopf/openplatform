package handlers

import (
	"encoding/json"
	"github.com/karldoenitz/Tigo/TigoWeb"
	"net/http"
	"opengateway/engine"
	"opengateway/global"
	"strings"
)

type EngineHandler struct {
	TigoWeb.BaseHandler
}

func (e *EngineHandler) process() {
	url := e.Request.URL.Path
	serviceParams := strings.Split(url, "/")
	if len(serviceParams) < 4 {
		e.ResponseAsJson(global.RespNotFound())
		return
	}
	project := serviceParams[1]
	service := serviceParams[2]
	function := serviceParams[3]
	serviceInfo, err := engine.DiscoverService(url, project, service)
	if err == global.ServiceNotFoundError {
		e.ResponseWriter.WriteHeader(http.StatusNotFound)
		e.ResponseAsText("")
		return
	}
	if err != nil {
		e.ResponseAsJson(global.RespFailed(err.Error()))
		return
	}
	//conn, err := engine.GetGrpcConnection(serviceInfo.Iport)
	//if err != nil {
	//	fmt.Printf("get connection failed: %s\n", err.Error())
	//}

	resp, err := engine.Call(project, service, function, serviceInfo.Proto, serviceInfo.Content, serviceInfo.Iport, string(e.GetBody()), nil)
	if err == global.MethodFoundError {
		e.ResponseWriter.WriteHeader(http.StatusNotFound)
		e.ResponseAsText("")
		return
	}
	if err != nil {
		e.ResponseAsJson(global.RespFailed(err.Error()))
		return
	}
	result := map[string]interface{}{}
	json.Unmarshal([]byte(resp.Data), &result)
	e.ResponseAsJson(global.RespSuccess(result))
}

func (e *EngineHandler) Get() {
	e.process()
}

func (e *EngineHandler) Post() {
	e.process()
}
