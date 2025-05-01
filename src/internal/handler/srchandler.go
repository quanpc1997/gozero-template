package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"template/gozero/src/internal/logic"
	"template/gozero/src/internal/svc"
	"template/gozero/src/internal/types"
)

func SrcHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSrcLogic(r.Context(), svcCtx)
		resp, err := l.Src(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
