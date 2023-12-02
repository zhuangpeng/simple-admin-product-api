package propertyvalue

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/internal/logic/propertyvalue"
	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"
)

// swagger:route post /property/value/create propertyvalue CreatePropertyValue
//
// Create property value information | 创建PropertyValue
//
// Create property value information | 创建PropertyValue
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PropertyValueBase
//
// Responses:
//  200: BaseMsgResp

func CreatePropertyValueHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PropertyValueBase
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := propertyvalue.NewCreatePropertyValueLogic(r.Context(), svcCtx)
		resp, err := l.CreatePropertyValue(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
