package propertyvalue

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyValuePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPropertyValuePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyValuePageLogic {
	return &GetPropertyValuePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPropertyValuePageLogic) GetPropertyValuePage(req *types.PropertyValuePageReq) (resp *types.PropertyValueListResp, err error) {
	valuePageReq := new(productclient.PropertyValuePageReq)
	copier.Copy(valuePageReq, req)
	data, err := l.svcCtx.ProductRpc.GetPropertyValuePage(l.ctx, valuePageReq)
	if err != nil {
		return nil, err
	}
	valueListResp := make([]types.PropertyValueInfo, len(data.Data))
	for i, v := range data.Data {
		propertyResp := new(types.PropertyValueInfo)
		copier.Copy(propertyResp, v)
		valueListResp[i] = *propertyResp
	}
	return &types.PropertyValueListResp{Data: valueListResp, BaseListInfo: types.BaseListInfo{Total: data.Total}}, nil
}
