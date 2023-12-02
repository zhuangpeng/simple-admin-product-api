package brand

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/internal/logic/brand"
	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"
)

// swagger:route get /brand/page brand GetBrandPage
//
// Get brand list | 获取Brand列表
//
// Get brand list | 获取Brand列表
//
// Parameters:
//  + name: page
//    require: true
//    in: query
//    type: integer
//  + name: pageSize
//    require: true
//    in: query
//    type: integer
//  + name: name
//    require: false
//    in: query
//    type: string
//  + name: status
//    require: false
//    in: query
//    type: integer
//  + name: createdAt
//    require: false
//    in: query
//    type: array
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
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
