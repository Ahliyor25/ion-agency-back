package main

import (
	"github.com/Ahliyor25/ion-agency-back/internal/service"
	"github.com/Ahliyor25/ion-agency-back/internal/storage"
	"github.com/Ahliyor25/ion-agency-back/internal/transport/http/handlers"
	"github.com/Ahliyor25/ion-agency-back/internal/transport/http/router"

	"github.com/Ahliyor25/ion-agency-back/pkg/bootstrap/http"
	"github.com/Ahliyor25/ion-agency-back/pkg/config"
	"github.com/Ahliyor25/ion-agency-back/pkg/databases"
	"github.com/Ahliyor25/ion-agency-back/pkg/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		http.Module,
		config.Module,
		logger.Module,
		databases.PostgresModule,
		service.Module,
		storage.Module,
		handlers.Module,
		router.Module,
	)

	app.Run()
}
