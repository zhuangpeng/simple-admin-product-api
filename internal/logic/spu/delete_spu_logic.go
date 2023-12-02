package spu

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSpuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSpuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSpuLogic {
	return &DeleteSpuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteSpuLogic) DeleteSpu(req *types.IDReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.ProductRpc.DeleteSpu(l.ctx, &productclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{
		Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg),
	}, nil
}
