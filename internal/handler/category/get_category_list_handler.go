package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/internal/logic/category"
	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"
)

// swagger:route get /category/list category GetCategoryList
//
// Get category list | 获取Category列表
//
// Get category list | 获取Category列表
//
// Parameters:
//  + name: parentId
//    require: false
//    in: path
//    type: integer
//  + name: name
//    require: false
//    in: path
//    type: string
//  + name: status
//    require: false
//    in: path
//    type: integer
//
// Responses:
//  200: CategoryListResp

func GetCategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewGetCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.GetCategoryList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
