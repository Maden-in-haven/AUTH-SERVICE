package handler

import (
	"auth/internal/gen"
	"auth/internal/services"
	"context"
)

func (s *AuthService) AuthVerifyPost(ctx context.Context, req *gen.AuthVerifyPostReq) (gen.AuthVerifyPostRes, error) {
	_, err := services.ValidateJWT(req.AccessToken.Value)
	if err != nil {
		// Токен недействителен или произошла ошибка
		return &gen.AuthVerifyPostUnauthorized{
			Message: gen.OptString{
				Value: "Неверный или истекший токен",
				Set:   true,
			},
		}, nil
	}

	// Если токен валиден, возвращаем успешный ответ
	return &gen.AuthVerifyPostOK{
		Valid: gen.OptBool{Value: true, Set: true},
	}, nil
}
