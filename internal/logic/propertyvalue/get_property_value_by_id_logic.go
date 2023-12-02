package propertyvalue

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyValueByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPropertyValueByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyValueByIdLogic {
	return &GetPropertyValueByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPropertyValueByIdLogic) GetPropertyValueById(req *types.IDReq) (resp *types.PropertyValueInfo, err error) {
	data, err := l.svcCtx.ProductRpc.GetPropertyValueById(l.ctx, &productclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	value := new(types.PropertyValueInfo)
	copier.Copy(value, data)
	return value, nil
}
