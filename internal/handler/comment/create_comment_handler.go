package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/agui-coder/simple-admin-product-api/internal/logic/comment"
	"github.com/agui-coder/simple-admin-product-api/internal/svc"
	"github.com/agui-coder/simple-admin-product-api/internal/types"
)

// swagger:route post /comment/create comment CreateComment
//
// Create comment information | 创建Comment
//
// Create comment information | 创建Comment
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CommentInfo
//
// Responses:
//  200: BaseMsgResp

func CreateCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewCreateCommentLogic(r.Context(), svcCtx)
		resp, err := l.CreateComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
