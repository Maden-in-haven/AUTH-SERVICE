package handler

import (
	"auth/internal/gen"
	"context"
	"log"

	"github.com/Maden-in-haven/crmlib/pkg/myjwt"
	"github.com/Maden-in-haven/crmlib/pkg/user"
)

func (s *AuthService) AuthLoginPost(ctx context.Context, req *gen.AuthLoginPostReq) (gen.AuthLoginPostRes, error) {
	// Логируем начало запроса
	log.Printf("Начало запроса на авторизацию для пользователя: %s", req.Username.Value)

	// Проверка корректности логина и пароля
	user, err := user.AuthenticateUser(req.Username.Value, req.Password.Value)
	if err != nil {
		log.Printf("Ошибка авторизации пользователя %s: %v", req.Username.Value, err)
		return &gen.AuthLoginPostUnauthorized{
			Message: gen.OptString{
				Value: "Неверный указан пользователь или пароль",
				Set:   true,
			},
		}, nil
	}
	log.Printf("Пользователь %s успешно аутентифицирован, генерируем токен", req.Username.Value)

	// Генерация JWT токена
	token, err := myjwt.GenerateJWT(user.ID)
	if err != nil {
		log.Printf("Ошибка генерации JWT токена для пользователя %s: %v", user.ID, err)
		// Возвращаем 500 Internal Server Error в случае ошибки при генерации JWT
		return &gen.AuthLoginPostInternalServerError{
			Message: gen.OptString{
				Value: "Ошибка при генерации JWT токена",
				Set:   true,
			},
		}, nil
	}
	log.Printf("JWT токен успешно сгенерирован для пользователя %s", user.ID)

	// Генерация Refresh токена
	refreshToken, err := myjwt.GenerateRefreshToken(user.ID)
	if err != nil {
		log.Printf("Ошибка генерации Refresh токена для пользователя %s: %v", user.ID, err)
		// Возвращаем 500 Internal Server Error в случае ошибки при генерации Refresh токена
		return &gen.AuthLoginPostInternalServerError{
			Message: gen.OptString{
				Value: "Ошибка при генерации рефреш токена",
				Set:   true,
			},
		}, nil
	}
	log.Printf("Refresh токен успешно сгенерирован для пользователя %s", user.ID)

	// Логируем успешную авторизацию
	log.Printf("Авторизация пользователя %s успешно завершена", req.Username.Value)

	// Возвращаем ответ с токеном и рефреш токеном
	return &gen.AuthLoginPostOK{
		AccessToken:  gen.OptString{Value: token, Set: true},
		RefreshToken: gen.OptString{Value: refreshToken, Set: true},
	}, nil
}
