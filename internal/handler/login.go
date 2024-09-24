package handler

import (
	"auth/internal/gen"
	"auth/internal/services"
	"context"
)

func (s *AuthService) AuthLoginPost(ctx context.Context, req *gen.AuthLoginPostReq) (gen.AuthLoginPostRes, error) {
	// Проверка корректности логина и пароля
	user, err := services.AuthenticateUser(req.Username.Value, req.Password.Value)
	if err != nil {
		return &gen.AuthLoginPostUnauthorized{
			Message: gen.OptString{
				Value: "Неверный указан пользователь или пароль",
				Set:   true,
			},
		}, nil
	}

	// Генерация JWT токена
	token, err := services.GenerateJWT(user.ID)
	if err != nil {
		// Возвращаем 500 Internal Server Error в случае ошибки при генерации JWT
		return &gen.AuthLoginPostInternalServerError{
			Message: gen.OptString{
				Value: "Ошибка при генерации JWT токена",
				Set:   true,
			},
		}, nil
	}

	refreshToken, err := services.GenerateRefreshToken(user.ID)
	if err != nil {
		// Возвращаем 500 Internal Server Error в случае ошибки при генерации Refresh токена
		return &gen.AuthLoginPostInternalServerError{
			Message: gen.OptString{
				Value: "Ошибка при генерации рефреш токена",
				Set:   true,
			},
		}, nil
	}

	// Возвращаем ответ с токеном и рефреш токеном
	return &gen.AuthLoginPostOK{
		AccessToken:  gen.OptString{Value: token, Set: true},
		RefreshToken: gen.OptString{Value: refreshToken, Set: true},
	}, nil
}
