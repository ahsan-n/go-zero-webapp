package handler

import (
	"net/http"

	"bazaartechnologies/rating-service/go-zero-webapp/webapp/internal/logic"
	"bazaartechnologies/rating-service/go-zero-webapp/webapp/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getUserDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetUserDetailsLogic(r.Context(), svcCtx)
		resp, err := l.GetUserDetails()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
