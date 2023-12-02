package brand

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

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
	data, err := l.svcCtx.ProductRpc.DeleteBrand(l.ctx, &productclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{
		Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg),
	}, nil
}
