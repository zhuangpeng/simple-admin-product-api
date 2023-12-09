package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/api/internal/logic/category"
	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"
)

// swagger:route post /category/create category CreateCategory
//
// Create category information | 创建Category
//
// Create category information | 创建Category
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CategoryCreateReq
//
// Responses:
//  200: BaseMsgResp

func CreateCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryCreateReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewCreateCategoryLogic(r.Context(), svcCtx)
		resp, err := l.CreateCategory(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
