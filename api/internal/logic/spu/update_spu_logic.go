package spu

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
