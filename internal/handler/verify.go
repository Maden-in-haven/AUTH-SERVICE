package handler

import (
	"auth/internal/gen"
	"context"
	"log"

	"github.com/Maden-in-haven/crmlib/pkg/myjwt"
)

func (s *AuthService) AuthVerifyPost(ctx context.Context, req *gen.APIAuthVerifyPostReq) (gen.APIAuthVerifyPostRes, error) {
	// Логируем начало запроса валидации токена
	log.Printf("Запрос на валидацию токена: %s", req.AccessToken.Value)

	// Валидация токена
	_, err := myjwt.ValidateJWT(req.AccessToken.Value)
	if err != nil {
		log.Printf("Ошибка валидации токена: %v", err)
		// Токен недействителен или истек
		return &gen.APIAuthVerifyPostUnauthorized{
			Message: gen.OptString{
				Value: "Неверный или истекший токен",
				Set:   true,
			},
		}, nil
	}

	// Логируем успешную валидацию токена
	log.Println("Токен успешно валидирован")

	// Если токен валиден, возвращаем успешный ответ
	return &gen.APIAuthVerifyPostOK{
		Valid: gen.OptBool{Value: true, Set: true},
	}, nil
}
