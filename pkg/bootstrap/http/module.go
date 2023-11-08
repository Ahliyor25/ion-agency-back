package http

import (
	"github.com/Ahliyor25/ion-agency-back/pkg/bootstrap/http/middlewares"
	"github.com/Ahliyor25/ion-agency-back/pkg/bootstrap/http/misc"
	"github.com/Ahliyor25/ion-agency-back/pkg/bootstrap/http/server"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	middlewares.Module,
	misc.Module,

	server.ModuleLifecycleHooks,
	server.ServerModule,
)
