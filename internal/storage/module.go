package storage

import (
	"github.com/Ahliyor25/ion-agency-back/internal/storage/user"
	"go.uber.org/fx"
)

// Module ...
var Module = fx.Options(
	user.Module,
)
