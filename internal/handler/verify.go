package handler

import (
	"auth/internal/gen"
	"context"
	"log"

	services "github.com/Maden-in-haven/crmlib/pkg/jwt"
)

func (s *AuthService) AuthVerifyPost(ctx context.Context, req *gen.AuthVerifyPostReq) (gen.AuthVerifyPostRes, error) {
	// Логируем начало запроса валидации токена
	log.Printf("Запрос на валидацию токена: %s", req.AccessToken.Value)

	// Валидация токена
	_, err := services.ValidateJWT(req.AccessToken.Value)
	if err != nil {
		log.Printf("Ошибка валидации токена: %v", err)
		// Токен недействителен или истек
		return &gen.AuthVerifyPostUnauthorized{
			Message: gen.OptString{
				Value: "Неверный или истекший токен",
				Set:   true,
			},
		}, nil
	}

	// Логируем успешную валидацию токена
	log.Println("Токен успешно валидирован")

	// Если токен валиден, возвращаем успешный ответ
	return &gen.AuthVerifyPostOK{
		Valid: gen.OptBool{Value: true, Set: true},
	}, nil
}
