package handler

import (
	"net/http"

	"bookstore/api/internal/logic"
	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	"bookstore/api/utils"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func createUserHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			message := utils.ResponseError{
				Message: err.Error(),
			}
			httpx.WriteJson(w, 400, message)
			return
		}

		l := logic.NewCreateUserLogic(r.Context(), ctx)
		resp, err := l.CreateUser(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
