package user

import (
	"errors"
	"fmt"
	"github.com/Ahliyor25/ion-agency-back/internal/entities"
	"github.com/Ahliyor25/ion-agency-back/pkg/bootstrap/http/misc/response"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

func (p provider) Create(user entities.User) (err error) {
	duplicateEntryError := &pgconn.PgError{Code: "23505"}

	tx := p.postgres.Begin()

	err = p.postgres.Create(&user).Error

	if err != nil {
		tx.Rollback()

		if errors.As(err, &duplicateEntryError) {
			return response.ErrAlreadyRegistered
		}

		p.logger.WithFields(logrus.Fields{
			"err":  err,
			"user": fmt.Sprintf("%+v", user),
		}).Error("Error while creating user")

		err = response.ErrInternalServer
		return
	}

	tx.Commit()

	return
}
