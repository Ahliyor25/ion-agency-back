package handlers

import (
	"encoding/json"
	"github.com/Ahliyor25/ion-agency-back/internal/entities"
	"github.com/Ahliyor25/ion-agency-back/pkg/bootstrap/http/misc/response"
	"github.com/Ahliyor25/ion-agency-back/pkg/utils"
	"net/http"
)

func (h *Handler) HRegistration(rw http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriterJSON(rw)

	var userData entities.User

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Проверка целостности полученных данных
	err := decoder.Decode(&userData)
	if err != nil {
		resp.Message = response.ErrBadRequest.Error()
		return
	}

	hashPass, err := utils.HashPassword(userData.Password)

	if err != nil {

		resp.Message = response.ErrInternalServer.Error()
		return
	}

	h.auth.Registration(userData.Phone, hashPass, userData.FullName)
	resp.Message = response.ErrSuccess.Error()

	return

}
