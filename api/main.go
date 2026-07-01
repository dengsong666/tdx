package main

import "log"

// @title TDX API
// @version 1.0
// @description 通达信行情 HTTP API。
// @BasePath /api
func main() {
	cfg := LoadConfig()

	app, err := NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	defer app.Close()
	app.StartGbbqInit()

	router := NewRouter(app)
	if err := router.Run(cfg.Addr); err != nil {
		log.Fatalf("failed to run api server: %v", err)
	}
}
