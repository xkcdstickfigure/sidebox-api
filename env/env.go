package env

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var DatabaseUrl = os.Getenv("DATABASE_URL")
