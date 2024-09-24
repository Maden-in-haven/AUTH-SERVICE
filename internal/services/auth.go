package services

import (
	"auth/internal/database"
	"golang.org/x/crypto/bcrypt"
	"auth/internal/model"
)

// Функция для проверки пароля
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Функция для аутентификации пользователя
func AuthenticateUser(username, password string) (*model.User, error) {
	// Находим пользователя по username
	user, err := database.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// Проверяем пароль
	err = CheckPassword(user.PasswordHash, password)
	if err != nil {
		return nil, err
	}

	// Возвращаем пользователя, если пароль верен
	return user, nil
}