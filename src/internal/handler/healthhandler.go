package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"template/gozero/src/internal/logic"
	"template/gozero/src/internal/svc"
)

func HealthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.Request
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//	return
		//}

		l := logic.NewHealthLogic(r.Context(), svcCtx)
		resp, err := l.Health()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
