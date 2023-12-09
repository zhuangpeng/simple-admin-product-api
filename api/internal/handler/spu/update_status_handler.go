package spu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/api/internal/logic/spu"
	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"
)

// swagger:route put /spu/update-status spu UpdateStatus
//
// Update spu information | 更新Spu
//
// Update spu information | 更新Spu
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SpuUpdateStatusReq
//
// Responses:
//  200: BaseMsgResp

func UpdateStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SpuUpdateStatusReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := spu.NewUpdateStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateStatus(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
