package handler

import (
	"auth/internal/gen"
	"context"
	"log"

	"github.com/Maden-in-haven/crmlib/pkg/myjwt"
)

func (s *AuthService) APIAuthRefreshPost(ctx context.Context, req *gen.APIAuthRefreshPostReq) (gen.APIAuthRefreshPostRes, error) {
	// Логируем начало запроса на обновление токенов
	log.Printf("Запрос на обновление токенов для рефреш токена: %s", req.RefreshToken.Value)

	// Валидация рефреш токена
	claims, err := myjwt.ValidateJWT(req.RefreshToken.Value)
	if err != nil {
		log.Printf("Ошибка валидации рефреш токена: %v", err)
		// Рефреш токен недействителен или истек
		return &gen.APIAuthRefreshPostUnauthorized{
			Message: gen.OptString{
				Value: "Неверный или истекший рефреш токен",
				Set:   true,
			},
		}, nil
	}

	// Извлекаем данные пользователя из claims
	userID, ok := claims["sub"].(string)
	if !ok {
		log.Println("Неверные данные в рефреш токене: отсутствует 'sub'")
		return &gen.APIAuthRefreshPostUnauthorized{
			Message: gen.OptString{
				Value: "Неверные данные в рефреш токене",
				Set:   true,
			},
		}, nil
	}
	log.Printf("Рефреш токен успешно валидирован для пользователя %s", userID)

	// Генерация нового токена доступа (JWT)
	newAccessToken, err := myjwt.GenerateJWT(userID)
	if err != nil {
		log.Printf("Ошибка генерации нового токена доступа для пользователя %s: %v", userID, err)
		return &gen.APIAuthRefreshPostInternalServerError{
			Message: gen.OptString{
				Value: "Ошибка при генерации нового токена",
				Set:   true,
			},
		}, nil
	}
	log.Printf("Новый токен доступа успешно сгенерирован для пользователя %s", userID)

	// Генерация нового рефреш токена
	refreshToken, err := myjwt.GenerateRefreshToken(userID)
	if err != nil {
		log.Printf("Ошибка генерации рефреш токена для пользователя %s: %v", userID, err)
		return &gen.APIAuthRefreshPostInternalServerError{
			Message: gen.OptString{
				Value: "Ошибка при генерации рефреш токена",
				Set:   true,
			},
		}, nil
	}
	log.Printf("Новый рефреш токен успешно сгенерирован для пользователя %s", userID)

	// Логируем успешное завершение запроса на обновление токенов
	log.Printf("Обновление токенов для пользователя %s успешно завершено", userID)

	// Возвращаем новый токен доступа и новый рефреш токен
	return &gen.APIAuthRefreshPostOK{
		AccessToken: gen.OptString{
			Value: newAccessToken,
			Set:   true,
		},
		RefreshToken: gen.OptString{
			Value: refreshToken,
			Set:   true,
		},
	}, nil
}
