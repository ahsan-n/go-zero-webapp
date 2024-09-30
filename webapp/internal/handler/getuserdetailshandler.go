package handler

import (
	"ahsan-n/go-zero-webapp/webapp/internal/types"
	"fmt"
	"net/http"

	"ahsan-n/go-zero-webapp/webapp/internal/logic"
	"ahsan-n/go-zero-webapp/webapp/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getUserDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetUserDetailsLogic(r.Context(), svcCtx)
		resp, err := l.GetUserDetails(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
