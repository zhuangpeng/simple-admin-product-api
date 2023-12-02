package spu

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSpuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSpuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSpuLogic {
	return &UpdateSpuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateSpuLogic) UpdateSpu(req *types.SpuUpdateReq) (resp *types.BaseMsgResp, err error) {
	valueUpdateReq := new(productclient.SpuUpdateReqVO)
	copier.Copy(valueUpdateReq, req)
	data, err := l.svcCtx.ProductRpc.UpdateSpu(l.ctx, valueUpdateReq)
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
