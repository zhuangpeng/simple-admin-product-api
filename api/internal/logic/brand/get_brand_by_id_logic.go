package brand

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBrandByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBrandByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBrandByIdLogic {
	return &GetBrandByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetBrandByIdLogic) GetBrandById(req *types.IDAtPathReq) (resp *types.BrandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
