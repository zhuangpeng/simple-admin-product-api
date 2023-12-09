package spu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/api/internal/logic/spu"
	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
)

// swagger:route post /spu/simple-list spu GetSpuSimpleList
//
// Get spu list | 获取Spu列表
//
// Get spu list | 获取Spu列表
//
// Responses:
//  200: SpuSimpleListResp

func GetSpuSimpleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := spu.NewGetSpuSimpleListLogic(r.Context(), svcCtx)
		resp, err := l.GetSpuSimpleList()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
