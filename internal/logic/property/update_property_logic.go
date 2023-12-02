package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePropertyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePropertyLogic {
	return &UpdatePropertyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdatePropertyLogic) UpdateProperty(req *types.PropertyUpdateReq) (resp *types.BaseMsgResp, err error) {
	propertyUpdateReq := new(productclient.PropertyUpdateReq)
	copier.Copy(propertyUpdateReq, req)
	data, err := l.svcCtx.ProductRpc.UpdateProperty(l.ctx, propertyUpdateReq)
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
