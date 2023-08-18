// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"lebron/apps/product/rpc/internal/logic"
	"lebron/apps/product/rpc/internal/svc"
	"lebron/apps/product/rpc/rpc"
)

type RpcServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedRpcServer
}

func NewRpcServer(svcCtx *svc.ServiceContext) *RpcServer {
	return &RpcServer{
		svcCtx: svcCtx,
	}
}

func (s *RpcServer) Ping(ctx context.Context, in *rpc.Request) (*rpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
