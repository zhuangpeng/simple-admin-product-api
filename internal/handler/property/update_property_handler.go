package property

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/internal/logic/property"
	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"
)

// swagger:route post /property/update property UpdateProperty
//
// Update property information | 更新Property
//
// Update property information | 更新Property
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: PropertyUpdateReq
//
// Responses:
//  200: BaseMsgResp

func UpdatePropertyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PropertyUpdateReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := property.NewUpdatePropertyLogic(r.Context(), svcCtx)
		resp, err := l.UpdateProperty(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
