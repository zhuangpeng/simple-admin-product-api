package propertyvalue

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyValuePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPropertyValuePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyValuePageLogic {
	return &GetPropertyValuePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPropertyValuePageLogic) GetPropertyValuePage(req *types.PropertyValuePageReq) (resp *types.PropertyValueListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
