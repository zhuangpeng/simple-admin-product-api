package propertyvalue

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePropertyValueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePropertyValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePropertyValueLogic {
	return &DeletePropertyValueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeletePropertyValueLogic) DeletePropertyValue(req *types.IDReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
