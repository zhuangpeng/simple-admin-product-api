package propertyvalue

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
