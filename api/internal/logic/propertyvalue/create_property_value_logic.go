package propertyvalue

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
