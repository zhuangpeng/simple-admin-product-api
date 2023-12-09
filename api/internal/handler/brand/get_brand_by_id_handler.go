package brand

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/api/internal/logic/brand"
	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"
)

// swagger:route get /brand/{id} brand GetBrandById
//
// Get brand by ID | 通过ID获取Brand
//
// Get brand by ID | 通过ID获取Brand
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDAtPathReq
//
// Responses:
//  200: BrandResp

func GetBrandByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDAtPathReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := brand.NewGetBrandByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetBrandById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
