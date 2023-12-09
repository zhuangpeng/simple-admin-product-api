package property

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPropertyPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyPageLogic {
	return &GetPropertyPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPropertyPageLogic) GetPropertyPage(req *types.PropertyPageReq) (resp *types.PropertyListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
