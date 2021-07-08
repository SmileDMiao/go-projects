package main

import (
	"go-learn/global"
	"go-learn/initialize"
	"log"
	"time"

	"github.com/fvbock/endless"
)

func main() {
	global.VP = initialize.Viper()

	initialize.Viper()
	initialize.Redis()

	global.DB = initialize.Gorm()

	db, _ := global.DB.DB()
	defer db.Close()

	router := initialize.Routers()

	port := "127.0.0.1:8000"
	server := endless.NewServer(port, router)
	server.ReadHeaderTimeout = 10 * time.Millisecond
	server.WriteTimeout = 10 * time.Second
	server.MaxHeaderBytes = 1 << 20

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
