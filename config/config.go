package config

import (
	"os"
)

var (
	PORT string
)

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "3001"
	}

}
