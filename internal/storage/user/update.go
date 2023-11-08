package user

import (
	"github.com/Ahliyor25/ion-agency-back/internal/entities"
	"github.com/sirupsen/logrus"
)

func (p provider) Update(data entities.User) (err error) {

	err = p.postgres.Model(&entities.User{}).
		Where(data.ID).
		Updates(data).Error

	if err != nil {
		p.logger.WithFields(logrus.Fields{
			"err":  err,
			"data": data,
		}).Error("Error while updating user")

	}
	return
}
