package spu

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpuSimpleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSpuSimpleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpuSimpleListLogic {
	return &GetSpuSimpleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetSpuSimpleListLogic) GetSpuSimpleList() (resp *types.SpuSimpleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
