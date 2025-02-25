package main

import (
	"app/internal/app"
	"app/internal/config"
)

// Точка входа в приложение
func main() {
	// чтение конфигурационных параметров
    cfg, err := config.New()
	if err != nil {
		panic("Can't initialize configuration")
	}
	// запуск приложения (то, что запускается описано в app)
	app.Run(cfg)
}