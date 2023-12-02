package category

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

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
	data, err := l.svcCtx.ProductRpc.GetCategoryById(l.ctx, &productclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	category := new(types.CategoryResp)
	copier.Copy(category, data)
	return category, nil
}
