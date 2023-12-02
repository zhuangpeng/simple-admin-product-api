package brand

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/internal/logic/brand"
	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"
)

// swagger:route get /brand/list brand GetBrandList
//
// Get brand list | 获取品牌精简信息列表
//
// Get brand list | 获取品牌精简信息列表
//
// Parameters:
//  + name: name
//    require: false
//    in: path
//    type: string
//
// Responses:
//  200: BrandListInfoResp

func GetBrandListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BrandListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := brand.NewGetBrandListLogic(r.Context(), svcCtx)
		resp, err := l.GetBrandList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
