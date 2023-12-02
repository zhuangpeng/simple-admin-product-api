package propertyvalue

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/internal/logic/propertyvalue"
	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"
)

// swagger:route post /property/value/update propertyvalue UpdatePropertyValue
//
// Update property value information | 更新PropertyValue
//
// Update property value information | 更新PropertyValue
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PropertyValueUpdateReq
//
// Responses:
//  200: BaseMsgResp

func UpdatePropertyValueHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PropertyValueUpdateReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := propertyvalue.NewUpdatePropertyValueLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePropertyValue(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
