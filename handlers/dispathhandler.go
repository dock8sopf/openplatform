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
		e.ResponseAsText("service not found")
		return
	}
	if err != nil {
		e.ResponseAsJson(global.RespFailed(err.Error()))
		return
	}

	resp, err := engine.Call(project, service, function, serviceInfo.Proto, serviceInfo.Content, serviceInfo.Iport, e.getParamToJsonFmt(), nil)
	if err == global.MethodFoundError {
		e.ResponseWriter.WriteHeader(http.StatusNotFound)
		e.ResponseAsText("method not found")
		return
	}
	if err != nil {
		e.ResponseAsJson(global.RespFailed(err.Error()))
		return
	}
	result := map[string]interface{}{}
	_ = json.Unmarshal([]byte(resp.Data), &result)
	e.ResponseAsJson(global.RespSuccess(result))
}

func (e *EngineHandler) Get() {
	e.process()
}

func (e *EngineHandler) Post() {
	e.process()
}

func (e *EngineHandler) Put() {
	e.process()
}

func (e *EngineHandler) Delete() {
	e.process()
}

func (e EngineHandler) getParamToJsonFmt() string {
	body := e.GetBody()
	if len(body) > 0 {
		return string(body)
	}
	urlParams := e.Request.URL.RawQuery
	jr := map[string]string{}
	for _, urlParam := range strings.Split(urlParams, "&") {
		uv := strings.Split(urlParam, "=")
		if len(uv) >= 2 {
			jr[uv[0]] = uv[1]
		} else {
			jr[uv[0]] = ""
		}
	}
	paramJson, err := json.Marshal(jr)
	if err == nil {
		return string(paramJson)
	}
	return "{}"
}
