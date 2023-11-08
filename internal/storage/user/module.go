package user

import (
	"github.com/Ahliyor25/ion-agency-back/internal/entities"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Options(

	fx.Provide(NewSLogin))

type SUser interface {
	Create(user entities.User) (err error)
	Get(target entities.User) (data entities.User, err error)
	Update(data entities.User) (err error)
}

type Dependencies struct {
	fx.In
	Logger   *logrus.Logger
	Postgres *gorm.DB
}

type provider struct {
	logger   *logrus.Logger
	postgres *gorm.DB
}

func NewSLogin(params Dependencies) SUser {
	return &provider{
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
