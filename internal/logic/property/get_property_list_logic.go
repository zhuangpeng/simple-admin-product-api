package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPropertyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyListLogic {
	return &GetPropertyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPropertyListLogic) GetPropertyList(req *types.PropertyListReq) (resp *types.PropertyListResp, err error) {
	data, err := l.svcCtx.ProductRpc.GetPropertyListByName(l.ctx,
		&productclient.PropertyListByNameReq{Name: req.Name})
	if err != nil {
		return nil, err
	}
	logx.Info(data)
	propertyListResp := make([]types.PropertyInfo, len(data.Data))
	for i, v := range data.Data {
		propertyResp := new(types.PropertyInfo)
		copier.Copy(propertyResp, v)
		propertyListResp[i] = *propertyResp
	}
	return &types.PropertyListResp{Data: propertyListResp}, nil
}
