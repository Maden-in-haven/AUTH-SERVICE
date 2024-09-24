// package main

// import (
// 	"log"
// 	"auth/internal/database"
// )

// func main() {
// 	// Инициализируем подключение к базе данных
// 	err := database.InitDatabase()
// 	if err != nil {
// 		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
// 	}
// 	defer database.DbPool.Close() // Используем напрямую глобальную переменную dbPool

// 	// Выполняем запросы к базе данных через dbPool
// }

package main

import (
	"auth/internal/gen"
	"auth/internal/handler"
	"log"
	"auth/internal/database"
	"net/http"
)

func main() {
	// Инициализируем подключение к базе данных
	err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer database.DbPool.Close() // Используем напрямую глобальную переменную dbPool

	// Создаем инстанс сервиса аутентификации.
	authService := &handler.AuthService{}

	// Создаем сгенерированный сервер от ogen.
	srv, err := gen.NewServer(authService)
	if err != nil {
		log.Fatalf("Ошибка при создании сервера: %v", err)
	}

	// Запускаем сервер на порту 8080.
	log.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
