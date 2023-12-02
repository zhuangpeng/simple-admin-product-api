package spu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/internal/logic/spu"
	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"
)

// swagger:route put /spu/update spu UpdateSpu
//
// Update spu information | 更新Spu
//
// Update spu information | 更新Spu
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SpuUpdateReq
//
// Responses:
//  200: BaseMsgResp

func UpdateSpuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SpuUpdateReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := spu.NewUpdateSpuLogic(r.Context(), svcCtx)
		resp, err := l.UpdateSpu(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
