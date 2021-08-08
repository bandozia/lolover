package global

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port    = 0
	Addr    = ""
	RootDir = "."
)

func LoadConfigs() error {
	var e error

	if e = godotenv.Load(); e != nil {
		return e
	}

	Port, e = strconv.Atoi(os.Getenv("HTTP_PORT"))
	Addr = os.Getenv("BIND_ADDR")
	RootDir = os.Getenv("ROOT_DIR")

	return e
}
