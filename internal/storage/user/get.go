package user

import (
	"errors"
	"github.com/Ahliyor25/ion-agency-back/internal/entities"
	"github.com/Ahliyor25/ion-agency-back/pkg/bootstrap/http/misc/response"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (p provider) Get(target entities.User) (data entities.User, err error) {
	err = p.postgres.Where(&target).
		Preload("Role").
		Preload("File").
		First(&data).Error

	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			err = response.ErrDataNotFound
			return
		}

		p.logger.WithFields(logrus.Fields{
			"err":    err,
			"target": target,
		}).Error("Error while getting user")

		return
	}
	return
}
