package spu

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
