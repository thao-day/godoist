package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"godoist/internal/logic"
	"godoist/internal/svc"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
