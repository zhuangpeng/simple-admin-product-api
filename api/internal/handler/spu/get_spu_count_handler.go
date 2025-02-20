package spu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/api/internal/logic/spu"
	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"
)

// swagger:route post /spu/count spu GetSpuCount
//
// Get spu by ID | 通过ID获取Spu
//
// Get spu by ID | 通过ID获取Spu
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: SpuCountResp

func GetSpuCountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := spu.NewGetSpuCountLogic(r.Context(), svcCtx)
		resp, err := l.GetSpuCount(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
