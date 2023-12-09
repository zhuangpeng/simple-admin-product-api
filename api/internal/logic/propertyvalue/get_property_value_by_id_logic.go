package propertyvalue

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
