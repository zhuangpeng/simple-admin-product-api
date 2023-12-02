package brand

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/internal/logic/brand"
	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"
)

// swagger:route put /brand/update brand UpdateBrand
//
// Update brand information | 更新Brand
//
// Update brand information | 更新Brand
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: BrandUpdateReq
//
// Responses:
//  200: BaseMsgResp

func UpdateBrandHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BrandUpdateReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := brand.NewUpdateBrandLogic(r.Context(), svcCtx)
		resp, err := l.UpdateBrand(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
