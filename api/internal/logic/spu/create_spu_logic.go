package spu

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
