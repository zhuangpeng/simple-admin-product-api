package category

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryListLogic {
	return &GetCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetCategoryListLogic) GetCategoryList(req *types.CategoryListReq) (resp *types.CategoryListResp, err error) {
	data, err := l.svcCtx.ProductRpc.GetEnableCategoryList(l.ctx,
		&productclient.CategoryListReq{Name: req.Name, Status: req.Status, ParentId: req.ParentId})
	if err != nil {
		return nil, err
	}
	logx.Info(data)
	categoryListResp := make([]types.CategoryResp, len(data.Data))
	for i, v := range data.Data {
		categoryResp := new(types.CategoryResp)
		copier.Copy(categoryResp, v)
		categoryListResp[i] = *categoryResp
	}
	return &types.CategoryListResp{Data: categoryListResp}, nil
}
