package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/internal/logic/category"
	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"
)

// swagger:route get /category/{id} category GetCategoryById
//
// Get category by ID | 通过ID获取Category
//
// Get category by ID | 通过ID获取Category
//
// Parameters:
//  + name: id
//    require: true
//    in: path
//    type: integer
//
// Responses:
//  200: CategoryResp

func GetCategoryByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDAtPathReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewGetCategoryByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetCategoryById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
