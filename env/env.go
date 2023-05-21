package env

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var DatabaseUrl = os.Getenv("DATABASE_URL")

var Origin = os.Getenv("ORIGIN")

var GoogleClientId = os.Getenv("GOOGLE_CLIENT_ID")
var GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

var ReceiveSecret = os.Getenv("RECEIVE_SECRET")
