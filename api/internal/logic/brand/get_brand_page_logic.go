package brand

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBrandPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBrandPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBrandPageLogic {
	return &GetBrandPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetBrandPageLogic) GetBrandPage(req *types.BrandPageReq) (resp *types.BrandListInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
