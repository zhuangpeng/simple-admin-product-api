package property

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePropertyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePropertyLogic {
	return &DeletePropertyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeletePropertyLogic) DeleteProperty(req *types.IDReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
