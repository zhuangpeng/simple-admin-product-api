package spu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/internal/logic/spu"
	"github.com/agui-coder/simple-admin-product-api/internal/svc"
)

// swagger:route post /spu/count spu GetSpuCount
//
// Get spu by ID | 通过ID获取Spu
//
// Get spu by ID | 通过ID获取Spu
//
// Responses:
//  200: SpuCountResp

func GetSpuCountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := spu.NewGetSpuCountLogic(r.Context(), svcCtx)
		resp, err := l.GetSpuCount()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
