package handler

import (
	"go_one/helper/server/http/response"
	"go_one/internal/registry"
	"net/http"
)

type ClientHandler interface {
	List() http.HandlerFunc
}

type clientHandler struct {
	reg *registry.ServiceContext
}

func NewClientHandler(reg *registry.ServiceContext) ClientHandler {
	return &clientHandler{
		reg: reg,
	}
}

func (p *clientHandler) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{
			"hello": "hello",
		}
		response.OkJson(r.Context(), w, resp, nil)
	}
}
