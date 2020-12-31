package main

import (
	"YubbeServer/yubbe-server/web/router"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	err := loadEnvVars()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	} else {
		log.Println("Variáveis de ambiente carregadas")
	}

	router := router.Router{}
	router.ApplyRoutes()

	fmt.Println("Server running")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func loadEnvVars() error {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	return viper.ReadInConfig()
}
