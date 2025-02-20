package property

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPropertyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyListLogic {
	return &GetPropertyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPropertyListLogic) GetPropertyList(req *types.PropertyListReq) (resp *types.PropertyListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
