package spu

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStatusLogic {
	return &UpdateStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateStatusLogic) UpdateStatus(req *types.SpuUpdateStatusReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.ProductRpc.UpdateStatus(l.ctx, &productclient.SpuUpdateStatusReq{Id: req.Id, Status: req.Status})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
