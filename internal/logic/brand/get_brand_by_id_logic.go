package brand

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

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
	data, err := l.svcCtx.ProductRpc.GetBrandById(l.ctx, &productclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	brand := new(types.BrandResp)
	copier.Copy(brand, data)
	return brand, nil
}
