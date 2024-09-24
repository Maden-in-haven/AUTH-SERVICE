package handler

import (
	"auth/internal/gen"
	"auth/internal/services"
	"context"
)

func (s *AuthService) AuthRefreshPost(ctx context.Context, req *gen.AuthRefreshPostReq) (gen.AuthRefreshPostRes, error) {
    // Валидация рефреш токена
    claims, err := services.ValidateJWT(req.RefreshToken.Value)
    if err != nil {
        // Рефреш токен недействителен или истек
        return &gen.AuthRefreshPostUnauthorized{
            Message: gen.OptString{
                Value: "Неверный или истекший рефреш токен",
                Set:   true,
            },
        }, nil
    }

    // Извлекаем данные пользователя из claims
    userID, ok := claims["sub"].(string)
    if !ok {
        return &gen.AuthRefreshPostUnauthorized{
            Message: gen.OptString{
                Value: "Неверные данные в рефреш токене",
                Set:   true,
            },
        }, nil
    }

    // Генерация нового токена доступа (JWT)
    newAccessToken, err := services.GenerateJWT(userID)
    if err != nil {
        return &gen.AuthRefreshPostInternalServerError{
            Message: gen.OptString{
                Value: "Ошибка при генерации нового токена",
                Set:   true,
            },
        }, nil
    }

    refreshToken, err := services.GenerateRefreshToken(userID)
    if err != nil {
        return &gen.AuthRefreshPostInternalServerError{
            Message: gen.OptString{
                Value: "Ошибка при генерации рефреш токена",
                Set:   true,
            },
        }, nil
    }

    // Возвращаем новый токен доступа
    return &gen.AuthRefreshPostOK{
        AccessToken: gen.OptString{
            Value: newAccessToken,
            Set:   true,
        },
        RefreshToken: gen.OptString{
            Value: refreshToken, // Опционально: можете вернуть новый рефреш токен, если он был сгенерирован
            Set:   true,
        },
    }, nil
}
