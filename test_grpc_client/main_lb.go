package main

import (
	"context"
	"flag"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"

	"github.com/sirupsen/logrus"
	grpclb "github.com/wwcd/grpc-lb/etcdv3"
	pb "gitee.com/y0/nvwa-sys/pb/admin"
)

var (
	svc = flag.String("service", "admin_adminsrv", "service name")
	reg = flag.String("reg", "http://localhost:2379", "register etcd address")
)

func main() {
	flag.Parse()
	r := grpclb.NewResolver(*reg, *svc)
	resolver.Register(r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	/// authority
	conn, err := grpc.DialContext(ctx, r.Scheme()+"://authority/"+*svc, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithBlock())
	cancel()
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(10 * time.Millisecond)
	index := 0
	for t := range ticker.C {
		index++
		client := pb.NewAdminSrvClient(conn)
		resp, err := client.AdminUser(context.Background(), &pb.AdminUserRequest{
			In:  "11",
			In1: "22",
		})
		if err == nil {
			logrus.Infof("收到服务响应: %s %v\n", resp.Out, time.Now().UnixNano()-t.UnixNano())
		} else {
			logrus.Infof("服务不可用")
		}
	}
}
