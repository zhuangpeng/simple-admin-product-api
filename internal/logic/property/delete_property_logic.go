package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

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
	data, err := l.svcCtx.ProductRpc.DeleteProperty(l.ctx, &productclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{
		Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg),
	}, nil
}
