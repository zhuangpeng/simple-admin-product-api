package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPropertyByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyByIdLogic {
	return &GetPropertyByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPropertyByIdLogic) GetPropertyById(req *types.IDReq) (resp *types.PropertyInfo, err error) {
	data, err := l.svcCtx.ProductRpc.GetPropertyById(l.ctx, &productclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	category := new(types.PropertyInfo)
	copier.Copy(category, data)
	return category, nil
}
