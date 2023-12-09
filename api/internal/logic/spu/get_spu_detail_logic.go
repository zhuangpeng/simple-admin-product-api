package spu

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSpuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpuDetailLogic {
	return &GetSpuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetSpuDetailLogic) GetSpuDetail(req *types.IDReq) (resp *types.SpuDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
