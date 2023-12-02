package propertyvalue

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/internal/logic/propertyvalue"
	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"
)

// swagger:route post /property/value/page propertyvalue GetPropertyValuePage
//
// Get property value list | 获取PropertyValue列表
//
// Get property value list | 获取PropertyValue列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PropertyValuePageReq
//
// Responses:
//  200: PropertyValueListResp

func GetPropertyValuePageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PropertyValuePageReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := propertyvalue.NewGetPropertyValuePageLogic(r.Context(), svcCtx)
		resp, err := l.GetPropertyValuePage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
