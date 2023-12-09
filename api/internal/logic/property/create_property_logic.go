package property

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
