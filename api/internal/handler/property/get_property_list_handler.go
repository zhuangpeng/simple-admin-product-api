package property

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/api/internal/logic/property"
	"github.com/agui-coder/simple-admin-product-api/api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/api/internal/types"
)

// swagger:route post /property/list property GetPropertyList
//
// Get property list | 获取Property列表
//
// Get property list | 获取Property列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PropertyListReq
//
// Responses:
//  200: PropertyListResp

func GetPropertyListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PropertyListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := property.NewGetPropertyListLogic(r.Context(), svcCtx)
		resp, err := l.GetPropertyList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
