package property

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPropertyAndValueListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPropertyAndValueListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPropertyAndValueListLogic {
	return &GetPropertyAndValueListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPropertyAndValueListLogic) GetPropertyAndValueList(req *types.PropertyListReq) (resp *types.PropertyAndValueListResp, err error) {
	// todo: add your logic here and delete this line

	return
}