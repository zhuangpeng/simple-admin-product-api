package brand

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

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
	data, err := l.svcCtx.ProductRpc.GetBrandList(l.ctx, &productclient.BrandListReq{Name: req.Name})
	if err != nil {
		return nil, err
	}
	brandListResp := make([]types.BrandResp, len(data.Data))
	for i, v := range data.Data {
		brandResp := new(types.BrandResp)
		copier.Copy(brandResp, v)
		brandListResp[i] = *brandResp
	}
	return &types.BrandListInfoResp{Data: brandListResp}, nil
}
