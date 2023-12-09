package property

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
