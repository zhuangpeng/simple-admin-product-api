package property

import (
	"context"
	"github.com/agui-coder/simple-admin-product-rpc/common/convert"
	"github.com/agui-coder/simple-admin-product-rpc/productclient"
	"github.com/jinzhu/copier"

	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"

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
	data, err := l.svcCtx.ProductRpc.GetPropertyListByName(l.ctx,
		&productclient.PropertyListByNameReq{Name: req.Name})
	if err != nil {
		return nil, err
	}
	if len(data.Data) == 0 {
		return &types.PropertyAndValueListResp{PropertyAndValue: []*types.PropertyAndValue{}}, nil
	}
	data.Data = convert.SliceUniq(data.Data, func(t *productclient.PropertyInfo) *uint64 {
		return t.Id
	})
	ids := convert.SliceUpdateElement(data.Data, func(t *productclient.PropertyInfo, _ int) uint64 {
		return *t.Id
	})
	valuePage, err := l.svcCtx.ProductRpc.GetPropertyValueListByPropertyId(l.ctx, &productclient.IDsReq{Ids: ids})
	if err != nil {
		return nil, err
	}
	valueMap := convert.MultiMap(valuePage.Data, func(t *productclient.PropertyValueInfo) uint64 {
		return *t.PropertyId
	})
	listResp := convert.SliceUpdateElement(data.Data, func(t *productclient.PropertyInfo, _ int) *types.PropertyAndValue {
		respVO := new(types.PropertyAndValue)
		copier.Copy(respVO, t)
		if len(valuePage.Data) == 0 {
			respVO.Value = []*types.Value{}
		} else {
			respVO.Value = make([]*types.Value, len(valueMap[*t.Id]))
			for i, valueInfo := range valueMap[*t.Id] {
				respVO.Value[i] = &types.Value{
					Id:   valueInfo.Id,
					Name: valueInfo.Name,
				}
			}
		}
		return respVO
	})
	return &types.PropertyAndValueListResp{PropertyAndValue: listResp}, nil
}
