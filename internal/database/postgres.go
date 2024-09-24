package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"auth/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Database структура для работы с базой данных
type Database struct {
	Pool *pgxpool.Pool
}

// Глобальная переменная для хранения пула соединений
var DbPool *pgxpool.Pool

// InitDatabase инициализирует подключение к базе данных и сохраняет его в глобальной переменной
func InitDatabase() error {
	// Загружаем конфигурацию базы данных из переменных окружения
	// dbConfig := config.LoadDBConfig()

	// Формируем строку подключения на основе полученной конфигурации
	// connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	// 	dbConfig.User,
	// 	dbConfig.Password,
	// 	dbConfig.Host,
	// 	dbConfig.Port,
	// 	dbConfig.DBName,
	// )

	// Настраиваем контекст с тайм-аутом для подключения (5 секунд)
	// ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	// defer cancel() // Освобождаем ресурсы контекста по завершении
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		log.Fatalf("Unable to parse configuration: %v\n", err)
	}

	config.ConnConfig.Host = "db.crm.evil-chan.ru"
	config.ConnConfig.Port = 5432
	config.ConnConfig.User = "gen_user"
	config.ConnConfig.Password = "m:0oC.h?3L_WKl"
	config.ConnConfig.Database = "default_db"
	// Подключаемся к базе данных
	dbpool, err := pgxpool.New(context.Background(), config.ConnString())
	if err != nil {
		log.Printf("Ошибка подключения к базе данных: %v", err)
		return err
	}
	// postgres
	// Проверяем, что подключение успешно
	// err = dbpool.Ping(context.Background())
	// if err != nil {
	// 	log.Printf("Ошибка проверки подключения к базе данных (ping): %v", err)
	// 	return err
	// }

	// Сохраняем пул соединений в глобальную переменную
	DbPool = dbpool

	log.Println("Успешно подключено к базе данных")
	return nil
}

// NewDatabase создает новое подключение к базе данных
func NewDatabase() (*Database, error) {
	// Загружаем конфигурацию базы данных из переменных окружения
	dbConfig := config.LoadDBConfig()

	// Формируем строку подключения на основе полученной конфигурации
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	// Настраиваем контекст с тайм-аутом для подключения (5 секунд)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Освобождаем ресурсы контекста по завершении

	// Подключаемся к базе данных
	dbpool, err := pgxpool.New(ctx, connString)
	if err != nil {
		log.Printf("Ошибка подключения к базе данных: %v", err)
		return nil, err
	}

	// Проверяем, что подключение успешно
	err = dbpool.Ping(ctx)
	if err != nil {
		log.Printf("Ошибка проверки подключения к базе данных (ping): %v", err)
		return nil, err
	}

	log.Println("Успешно подключено к базе данных")

	// Возвращаем структуру базы данных с пулом соединений
	return &Database{
		Pool: dbpool,
	}, nil
}

// Close закрывает соединение с базой данных
func (db *Database) Close() {
	db.Pool.Close() // Закрываем пул соединений
	log.Println("Соединение с базой данных закрыто")
}
