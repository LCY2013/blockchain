// Code geneorderd by goctl. DO NOT EDIT.
// Source: register.proto

package server

import (
	"context"
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange/internal/logic"
	"github.com/LCY2013/blockchain/exchange/ff-coin/exchange/internal/svc"
	"github.com/LCY2013/blockchain/exchange/ff-coin/grpc-common/exchange/types/order"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServer
}

func (e *OrderServer) FindOrderHistory(ctx context.Context, req *order.OrderReq) (*order.OrderRes, error) {
	l := logic.NewExchangeOrderLogic(ctx, e.svcCtx)
	return l.FindOrderHistory(req)
}

func (e *OrderServer) FindOrderCurrent(ctx context.Context, req *order.OrderReq) (*order.OrderRes, error) {
	l := logic.NewExchangeOrderLogic(ctx, e.svcCtx)
	return l.FindOrderCurrent(req)
}
func (e *OrderServer) Add(ctx context.Context, req *order.OrderReq) (*order.AddOrderRes, error) {
	l := logic.NewExchangeOrderLogic(ctx, e.svcCtx)
	return l.Add(req)
}
func (e *OrderServer) FindByOrderId(ctx context.Context, req *order.OrderReq) (*order.ExchangeOrderOrigin, error) {
	l := logic.NewExchangeOrderLogic(ctx, e.svcCtx)
	return l.FindByOrderId(req)
}
func (e *OrderServer) CancelOrder(ctx context.Context, req *order.OrderReq) (*order.CancelOrderRes, error) {
	l := logic.NewExchangeOrderLogic(ctx, e.svcCtx)
	return l.CancelOrder(req)
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}
