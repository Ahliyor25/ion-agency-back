package databases

import (
	"fmt"
	"github.com/Ahliyor25/ion-agency-back/internal/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Postgres ...
func Postgres(params Dependencies) (pdb *gorm.DB) {
	connStr := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",

		params.Config.Postgres.Host,
		fmt.Sprint(params.Config.Postgres.Port),
		params.Config.Postgres.DatabaseName,
		params.Config.Postgres.Username,
		params.Config.Postgres.Password,
		params.Config.Postgres.SSLMode,
	)

	pdb, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  connStr,
				PreferSimpleProtocol: true,
			},
		),
		&gorm.Config{
			DisableAutomaticPing:                     true,
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}

	conn, err := pdb.DB()
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(); err != nil {
		panic(err)
	}

	params.Logger.Info("Postgres pong!")

	pdb.AutoMigrate(&entities.User{})
	pdb.AutoMigrate(&entities.Role{})
	pdb.AutoMigrate(&entities.File{})

	return
}
