package spu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/api/internal/logic/spu"
	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"
)

// swagger:route post /spu/page spu GetSpuPage
//
// Get spu list | 获取Spu列表
//
// Get spu list | 获取Spu列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SpuPageReq
//
// Responses:
//  200: SpuPageResp

func GetSpuPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SpuPageReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := spu.NewGetSpuPageLogic(r.Context(), svcCtx)
		resp, err := l.GetSpuPage(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
