package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPropertyPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyPageLogic {
	return &GetPropertyPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPropertyPageLogic) GetPropertyPage(req *types.PropertyPageReq) (resp *types.PropertyListResp, err error) {
	propertyPageReq := new(productclient.PropertyPageReq)
	copier.Copy(propertyPageReq, req)
	data, err := l.svcCtx.ProductRpc.GetPropertyPage(l.ctx, propertyPageReq)
	if err != nil {
		return nil, err
	}
	propertyListResp := make([]types.PropertyInfo, len(data.Data))
	for i, v := range data.Data {
		propertyResp := new(types.PropertyInfo)
		copier.Copy(propertyResp, v)
		propertyListResp[i] = *propertyResp
	}
	return &types.PropertyListResp{Data: propertyListResp, BaseListInfo: types.BaseListInfo{Total: data.Total}}, nil
}
