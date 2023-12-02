package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePropertyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePropertyLogic {
	return &CreatePropertyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreatePropertyLogic) CreateProperty(req *types.PropertyBase) (resp *types.BaseMsgResp, err error) {
	propertyCreateReq := new(productclient.PropertyCreateReq)
	copier.Copy(propertyCreateReq, req)
	data, err := l.svcCtx.ProductRpc.CreateProperty(l.ctx, propertyCreateReq)
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
