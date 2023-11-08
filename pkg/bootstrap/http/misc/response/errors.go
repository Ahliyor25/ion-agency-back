package response

import "errors"

// Errors used in the project. Sorted alphabetically
var (
	// Default errors
	ErrAccessDenied       = errors.New("Отказано в доступе")
	ErrBadRequest         = errors.New("Неверный запрос")
	ErrDataNotFound       = errors.New("Не найдено")
	ErrInternalServer     = errors.New("Внутренняя ошибка сервера")
	ErrLimitExceeded      = errors.New("Превышен лимит попыток. Повторите позже")
	ErrNoContent          = errors.New("no content")
	ErrNotImplementation  = errors.New("Не реализованно")
	ErrSomethingWentWrong = errors.New("Что-то пошло не так")
	ErrSuccess            = errors.New("Успешно")
	ErrTokenIsExpired     = errors.New("Время жизни токена просрочено")
	ErrUnauthorized       = errors.New("Не авторизованный пользователь")

	ErrAlreadyCreated               = errors.New("Заявка с таким идентификатором уже создана")
	ErrAlreadyHaveApplication       = errors.New("На данный момент у вас имеется заявка")
	ErrAlreadyRegistered            = errors.New("Этот номер телефона уже зарегистрирован")
	ErrApplicationHasBeenConcluded  = errors.New("Заявка приняла итоговое состояние")
	ErrApplicationIsBeingProcessed  = errors.New("Заявка находится в обработке")
	ErrBrokenFile                   = errors.New("Не удалось переобразовать файл")
	ErrEmptyFileds                  = errors.New("Одно или несколько полей пустые")
	ErrIncorrectApplicationID       = errors.New("Заявки с таким идентификатором не существует")
	ErrIncorrectFillFields          = errors.New("Одно или несколько полей заполненны не корректно")
	ErrIncorrectLoginOrPassword     = errors.New("Неправильный логин или пароль")
	ErrIncorrectOtpCode             = errors.New("Не верный OTP код")
	ErrNoFreeApplications           = errors.New("Нет свободных заявок")
	ErrNonExistentState             = errors.New("Указан несуществующее напрвление для текущей заявки")
	ErrUnknownChannel               = errors.New("Unknown channel")
	ErrUnknownProcess               = errors.New("Unknown process")
	ErrWrongBodyExternalApplication = errors.New("Нет свободных заявок")
)

// Error statuses
var errorCode = map[string]int{
	// Default errors
	ErrAccessDenied.Error():       403,
	ErrBadRequest.Error():         400,
	ErrDataNotFound.Error():       404,
	ErrInternalServer.Error():     500,
	ErrLimitExceeded.Error():      429,
	ErrNoContent.Error():          201,
	ErrNotImplementation.Error():  501,
	ErrSomethingWentWrong.Error(): 500,
	ErrSuccess.Error():            200,
	ErrTokenIsExpired.Error():     401,
	ErrUnauthorized.Error():       401,

	ErrAlreadyCreated.Error():               409,
	ErrAlreadyHaveApplication.Error():       409,
	ErrAlreadyRegistered.Error():            409,
	ErrApplicationHasBeenConcluded.Error():  400,
	ErrApplicationIsBeingProcessed.Error():  400,
	ErrBrokenFile.Error():                   400,
	ErrEmptyFileds.Error():                  400,
	ErrIncorrectApplicationID.Error():       400,
	ErrIncorrectFillFields.Error():          400,
	ErrIncorrectLoginOrPassword.Error():     400,
	ErrIncorrectOtpCode.Error():             400,
	ErrNoFreeApplications.Error():           404,
	ErrNonExistentState.Error():             404,
	ErrUnknownChannel.Error():               400,
	ErrUnknownChannel.Error():               400,
	ErrWrongBodyExternalApplication.Error(): 400,
}
