package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	log.Println("Reading configuration")

	viper.SetConfigName("config9")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("XTERM")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error in config file: %s \n", err)
	}
	log.Println("Done")

	log.Printf("Xterm locale: %s\n", viper.GetString("locale"))
	log.Printf("Xterm shell: %s\n", viper.GetString("shell"))
}
