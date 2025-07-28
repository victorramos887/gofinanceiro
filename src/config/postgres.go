package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/victorramos887/gofinanceiro/src/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializerPostgres() (*gorm.DB, error) {
	logger := GetLogger("Postgres")

	// Carrega variáveis do .env
	err := godotenv.Load()
	if err != nil {
		logger.Error("Erro ao carregar .env:", err)
		return nil, err
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, dbname, port)
	
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Erro ao conectar ao Postgres:", err)
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Ganhos{},
		&models.Gastos{},
	)
	if err != nil {
		logger.Error("Erro ao migrar o banco de dados:", err)
		return nil, err
	}

	logger.Info("Conexão com Postgres estabelecida com sucesso!")
	return db, nil
}
