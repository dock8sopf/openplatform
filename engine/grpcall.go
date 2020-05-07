package engine

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"opengateway/global"
	"opengateway/grpcall"
	"strings"
)

// Call 调用grpc服务
//  - project 项目名
//  - service 服务名
//  - function 方法名
//  - proto pb文件名
//  - protoContent pb文件内容
//  - iport ip和端口
//  - sendBody 发送的内容
//  - gclient grpc的客户端
func Call(project, service, function, proto, protoContent, iport, sendBody string, gclient *grpc.ClientConn) (result *grpcall.ResultModel, err error) {
	grpcall.SetProtoContents(project, proto, protoContent)
	grpcall.SetMode(grpcall.ProtoContentMode)
	err = grpcall.InitDescSource()
	if err != nil {
		return
	}

	handler := DefaultEventHandler{}

	grpcEnter, err := grpcall.New(
		grpcall.SetHookHandler(&handler),
	)
	if err != nil {
		return
	}
	err = grpcEnter.Init()
	if err != nil {
		return
	}
	res, err := grpcEnter.Call(iport, service, function, sendBody, gclient)
	if err != nil && strings.Contains(err.Error(), "does not include a method named") {
		return res, global.MethodFoundError
	}
	return res, err
}

type DefaultEventHandler struct {
	sendChan chan []byte
}

func (h *DefaultEventHandler) OnReceiveData(md metadata.MD, resp string, respErr error) {
}

func (h *DefaultEventHandler) OnReceiveTrailers(stat *status.Status, md metadata.MD) {
}
