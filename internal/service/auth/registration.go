package auth

import (
	"github.com/Ahliyor25/ion-agency-back/internal/entities"
	"github.com/sirupsen/logrus"
)

//

func (p provider) Registration(phone string, password string, fullName string) (err error) {

	// 1. Check if user exists

	err = p.postgres.Where(&entities.User{Phone: phone}).First(&entities.User{}).Error

	if err != nil {
		if err.Error() == "record not found" {
			// 2. Create user
			err = p.postgres.Create(&entities.User{
				Phone:    phone,
				Password: password,
				FullName: fullName,
			}).Error
			if err != nil {
				p.logger.WithFields(logrus.Fields{
					"err": err,
				}).Error("Error while creating user")
				return
			}
		} else {
			p.logger.WithFields(logrus.Fields{
				"err": err,
			}).Error("Error while getting user")
			return
		}
	}

	// 3. Create token

	// 4. Send sms

	// 5. Return token

	return

}
