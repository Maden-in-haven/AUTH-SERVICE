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
	"net/http"
)


func main() {
	authService := &handler.AuthService{}
	srv, err := gen.NewServer(authService)
	if err != nil {
		log.Fatalf("Ошибка при создании сервера: %v", err)
	}

	// Запускаем сервер на порту 80.
	log.Println("Сервер запущен на 0.0.0.0:80")
	if err := http.ListenAndServe(":80", srv); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
