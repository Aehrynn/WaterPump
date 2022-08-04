package main

import (
	"erinmaguire/WaterPump/controllers"
	"erinmaguire/WaterPump/structs"
	"fmt"
	"log"
	"net/http"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/spf13/viper"
	"github.com/stianeikeland/go-rpio/v4"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	var configuration structs.Configuration

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	myDBDir := configuration.DatabaseLocation
	tiedotDBClient, dbErr := db.OpenDB(myDBDir)
	if dbErr != nil {
		panic(dbErr)
	}

	tiedotDBClient.Create("logs")
	rpio.Open()
	defer rpio.Close()
	pin := rpio.Pin(configuration.WaterPumpRelayGpio)
	//4 22 6 26
	pin.Output()

	log.Println("Water Pump Local Server Started")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Get("/Logs", func(w http.ResponseWriter, r *http.Request) {
			controllers.GetWaterPumpLogs(w, r, configuration.ApiKey, tiedotDBClient)
		})

		r.Route("/ToggleWaterPump", func(r chi.Router) {
			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				controllers.PostToggleWaterPump(w, r, configuration.ApiKey, pin)
			})

		})
	})

	httpErr := http.ListenAndServe(":3000", r)
	if httpErr != nil {
		fmt.Println(httpErr)
	}
}
