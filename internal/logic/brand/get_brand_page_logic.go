package brand

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

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
	brandPageReq := new(productclient.BrandPageReq)
	copier.Copy(brandPageReq, req)
	data, err := l.svcCtx.ProductRpc.GetBrandPage(l.ctx, brandPageReq)
	if err != nil {
		return nil, err
	}
	brandListResp := make([]types.BrandResp, len(data.Data))
	for i, v := range data.Data {
		brandResp := new(types.BrandResp)
		copier.Copy(brandResp, v)
		brandListResp[i] = *brandResp
	}
	return &types.BrandListInfoResp{Data: brandListResp, BaseListInfo: types.BaseListInfo{Total: data.Total}}, nil
}
