package spu

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSpuDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSpuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSpuDetailLogic {
	return &GetSpuDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetSpuDetailLogic) GetSpuDetail(req *types.IDReq) (resp *types.SpuDetailResp, err error) {
	spu, err := l.svcCtx.ProductRpc.GetSpuById(l.ctx, &productclient.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	skus, err := l.svcCtx.ProductRpc.GetSkuListBySpuId(l.ctx, &productclient.IDReq{Id: req.Id})
	spuResp := new(types.SpuInfo)
	copier.Copy(spuResp, spu)
	SkuResp := make([]types.SkuInfo, len(skus.Data))
	for i, s := range skus.Data {
		sku := new(types.SkuInfo)
		copier.Copy(sku, s)
		SkuResp[i] = *sku
	}
	return &types.SpuDetailResp{
		IDReq:   types.IDReq{Id: *spu.Id},
		SpuInfo: *spuResp,
		Skus:    SkuResp,
	}, nil
}
