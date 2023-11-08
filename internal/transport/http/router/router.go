package router

import (
	"github.com/Ahliyor25/ion-agency-back/internal/transport/http/handlers"
	"github.com/Ahliyor25/ion-agency-back/pkg/bootstrap/http/middlewares"
	transportHTTP "github.com/Ahliyor25/ion-agency-back/pkg/bootstrap/http/router"
)

// NewRouter ..
func NewRouter(h *handlers.Handler, mw middlewares.Middleware) (router *transportHTTP.HTTPRouter) {
	router = transportHTTP.NewRouter()
	router.ConnectSwagger(h.ServeSwaggerFiles)

	router.GET("/ping", h.HPingPong, mw.CORS)

	return
}
