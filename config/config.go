package config

import "os"

var JwtSecret []byte

func InitConfig() {
    JwtSecret = []byte(os.Getenv("JWT_SECRET"))
}
