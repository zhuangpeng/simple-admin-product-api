package brand

import (
	"context"

	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBrandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBrandLogic {
	return &DeleteBrandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *DeleteBrandLogic) DeleteBrand(req *types.IDReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
