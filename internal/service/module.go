package service

import (
	"github.com/Ahliyor25/ion-agency-back/internal/service/auth"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	auth.Module,
)
