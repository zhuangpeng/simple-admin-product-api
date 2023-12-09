package brand

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/api/internal/logic/brand"
	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"
)

// swagger:route get /brand/page brand GetBrandPage
//
// Get brand page | 获取Brand列表
//
// Get brand page | 获取Brand列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: BrandPageReq
//
// Responses:
//  200: BrandListInfoResp

func GetBrandPageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BrandPageReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := brand.NewGetBrandPageLogic(r.Context(), svcCtx)
		resp, err := l.GetBrandPage(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
