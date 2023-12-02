package brand

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBrandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBrandLogic {
	return &CreateBrandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateBrandLogic) CreateBrand(req *types.BrandCreateReq) (resp *types.BaseMsgResp, err error) {
	brandCreate := new(productclient.BrandCreateReq)
	copier.Copy(brandCreate, req)
	data, err := l.svcCtx.ProductRpc.CreateBrand(l.ctx, brandCreate)
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
