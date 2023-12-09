package category

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryByIdLogic {
	return &GetCategoryByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetCategoryByIdLogic) GetCategoryById(req *types.IDAtPathReq) (resp *types.CategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
