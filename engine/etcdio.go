package engine

import (
	"context"
	"encoding/json"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/naming"
	"github.com/karldoenitz/Tigo/logger"
	"google.golang.org/grpc"
	"opengateway/global"
	"time"
)

type ServerInfo struct {
	Iport   string `json:"iport"`
	Proto   string `json:"proto"`
	Content string `json:"content"`
}

// RegisterToEtcd 向etcd注册参数
//  - url 请求的url
//  - project 项目名
//  - service 服务名
func DiscoverService(url string, project string, service string) (*ServerInfo, error) {
	etcdIport := global.GetEtcdIport()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdIport},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logger.Error.Printf("创建etcd客户端失败，etcd(%s) => (%s)", etcdIport, err.Error())
		return nil, err
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := cli.Get(ctx, project+"-"+service)
	if err != nil {
		logger.Error.Printf("从etcd(%s)获取值失败 => (%s)", etcdIport, err.Error())
		return nil, err
	}
	info := &ServerInfo{}
	if len(resp.Kvs) <= 0 {
		return nil, global.ServiceNotFoundError
	}
	for _, v := range resp.Kvs {
		if e := json.Unmarshal(v.Value, info); e != nil {
			return info, e
		}
		break
	}
	return info, nil
}

func GetGrpcConnection(serviceAddress string) (gclient *grpc.ClientConn, err error) {
	cli, cerr := clientv3.NewFromURL("127.0.0.1:12379")
	if cerr != nil {
		return nil, cerr
	}
	r := &naming.GRPCResolver{Client: cli}
	b := grpc.RoundRobin(r)
	conn, gerr := grpc.Dial(serviceAddress, grpc.WithInsecure(), grpc.WithBalancer(b))
	if gerr != nil {
		return nil, gerr
	}
	return conn, nil
}
