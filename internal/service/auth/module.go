package auth

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Provide(NewBAuth)

// Service ...

type BAuth interface {
	Registration(phone string, password string, fullName string) (err error)
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

func NewBAuth(params Dependencies) BAuth {
	return &provider{
		logger:   params.Logger,
		postgres: params.Postgres,
	}
}
