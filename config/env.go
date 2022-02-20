package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Env(name string) string {
	godotenv.Load(".env")
	switch name {
	case "MONGO_DB_URL":
		return os.Getenv(name)
	case "SERVER_PORT":
		return os.Getenv(name)
	case "JWT_SECRET_KEY":
		return os.Getenv(name)
	case "JWT_ADMIN_SECRET_KEY":
		return os.Getenv(name)
	default:
		panic(name + " doesnt exist in file .env")
	}
}
