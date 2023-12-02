package spu

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSpuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSpuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSpuLogic {
	return &CreateSpuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateSpuLogic) CreateSpu(req *types.SpuCreateReq) (resp *types.BaseMsgResp, err error) {
	spuCreateReq := new(productclient.SpuCreateReq)
	copier.Copy(spuCreateReq, req)
	data, err := l.svcCtx.ProductRpc.CreateSpu(l.ctx, spuCreateReq)
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
