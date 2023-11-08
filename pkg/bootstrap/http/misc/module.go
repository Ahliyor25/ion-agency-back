package misc

import (
	"github.com/Ahliyor25/ion-agency-back/pkg/bootstrap/http/misc/response"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	response.Module,
)
