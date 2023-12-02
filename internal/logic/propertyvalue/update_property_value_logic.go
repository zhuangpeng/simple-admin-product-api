package propertyvalue

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePropertyValueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePropertyValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePropertyValueLogic {
	return &UpdatePropertyValueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdatePropertyValueLogic) UpdatePropertyValue(req *types.PropertyValueUpdateReq) (resp *types.BaseMsgResp, err error) {
	valueUpdateReq := new(productclient.PropertyValueUpdateReq)
	copier.Copy(valueUpdateReq, req)
	data, err := l.svcCtx.ProductRpc.UpdatePropertyValue(l.ctx, valueUpdateReq)
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
