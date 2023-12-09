package brand

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBrandListLogic {
	return &GetBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetBrandListLogic) GetBrandList(req *types.BrandListReq) (resp *types.BrandListInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
