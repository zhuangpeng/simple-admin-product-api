package spu

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpuCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSpuCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpuCountLogic {
	return &GetSpuCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetSpuCountLogic) GetSpuCount() (resp *types.SpuCountResp, err error) {
	// todo: add your logic here and delete this line

	return
}
