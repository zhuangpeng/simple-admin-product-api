package propertyvalue

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePropertyValueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePropertyValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePropertyValueLogic {
	return &CreatePropertyValueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreatePropertyValueLogic) CreatePropertyValue(req *types.PropertyValueBase) (resp *types.BaseMsgResp, err error) {
	propertyValueCreateReq := new(productclient.PropertyValueCreateReq)
	copier.Copy(propertyValueCreateReq, req)
	data, err := l.svcCtx.ProductRpc.CreatePropertyValue(l.ctx, propertyValueCreateReq)
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
